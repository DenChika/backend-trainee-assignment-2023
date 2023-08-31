package repository

import (
	"backend-trainee-assignment-2023/pkg/helpers"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Transaction struct {
	*sqlx.Tx
}

func NewTransaction(tx *sqlx.Tx) *Transaction {
	return &Transaction{tx}
}

func (tx *Transaction) filterIfSegmentsExist(slugs []string) ([]uint, error) {
	slugsString := helpers.PrintSliceByComma(slugs)
	querySegments := fmt.Sprintf(
		`SELECT s.id FROM %s AS s 
            WHERE s.slug IN (%s)`, segmentsTable, slugsString)
	var segments []uint
	err := tx.Select(&segments, querySegments)
	return segments, err
}

func (tx *Transaction) findSlugsByIds(segmentsIds []uint) ([]string, error) {
	segmentsString := helpers.PrintSliceByComma(segmentsIds)
	querySegments := fmt.Sprintf(
		`SELECT s.slug FROM %s AS s 
            WHERE s.id IN (%s)`, segmentsTable, segmentsString)
	var segments []string
	err := tx.Select(&segments, querySegments)
	return segments, err
}

func (tx *Transaction) filterIfUsersSegmentsNotInDb(ids []uint) ([]uint, error) {
	slugsString := helpers.PrintSliceByComma(ids)
	var slugsInUsersSegments []uint
	querySlugs := fmt.Sprintf(
		`SELECT s.id FROM %s AS us
         JOIN %s AS s ON s.id=us.segment_id
         WHERE us.segment_id IN (%s)`, usersSegmentsTable, segmentsTable, slugsString)
	err := tx.Select(&slugsInUsersSegments, querySlugs)
	if err != nil {
		return []uint{}, err
	}
	return helpers.SlicesDifference(ids, slugsInUsersSegments), nil
}

func (tx *Transaction) insertSegmentsIntoUser(segmentsToAdd []uint, userId uint) error {
	var usersSegmentsBuilder strings.Builder
	usersSegmentsBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_id) VALUES ",
		usersSegmentsTable))

	for i, id := range segmentsToAdd {
		if i != 0 {
			usersSegmentsBuilder.WriteString(", ")
		}
		usersSegmentsBuilder.WriteString(
			fmt.Sprintf("('%d', '%d')", userId, id))
	}

	query := usersSegmentsBuilder.String()
	_, err := tx.Exec(query)
	return err
}

func (tx *Transaction) deleteSegmentsFromUser(segmentsToRemove []uint, userId uint) ([]uint, error) {
	segmentsToRemoveString := helpers.PrintSliceByComma(segmentsToRemove)
	query := fmt.Sprintf(
		`DELETE FROM %s AS us 
       WHERE us.segment_id IN (%s) AND us.user_id=$1
       RETURNING us.segment_id`, usersSegmentsTable, segmentsToRemoveString)
	rows, err := tx.Query(query, userId)
	if err != nil {
		return []uint{}, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var deletedRows []uint
	for rows.Next() {
		var row uint
		if err = rows.Scan(&row); err != nil {
			return []uint{}, err
		}
		deletedRows = append(deletedRows, row)
	}
	return deletedRows, nil
}

func (tx *Transaction) saveInHistory(segments []uint, userId uint, op operation) error {
	var usersSegmentsHistoryBuilder strings.Builder
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_slug, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	for i, id := range segments {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%d', '%s', 'now()')", userId, id, op))
	}

	query := usersSegmentsHistoryBuilder.String()
	_, err := tx.Exec(query)
	return err
}
