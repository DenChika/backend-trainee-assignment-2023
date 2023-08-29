package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId uint) error {
	slugsToAddString := printSliceByComma(slugsToAdd)
	slugsToRemoveString := printSliceByComma(slugsToRemove)

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", "REPEATABLE READ"))
	if err != nil {
		tx.Rollback()
		return err
	}
	type segment struct {
		Id   uint
		Slug string
	}

	querySegmentsToAdd := fmt.Sprintf(
		`SELECT * FROM %s AS s 
            WHERE s.slug IN (%s)`, segmentsTable, slugsToAddString)
	querySegmentsToRemove := fmt.Sprintf(
		`SELECT * FROM %s AS s 
            WHERE s.slug IN (%s)`, segmentsTable, slugsToRemoveString)
	var segmentsToAdd []segment
	if err := repo.db.Select(&segmentsToAdd, querySegmentsToAdd); err != nil {
		tx.Rollback()
		return err
	}
	var segmentsToRemove []segment
	if err := repo.db.Select(&segmentsToRemove, querySegmentsToRemove); err != nil {
		tx.Rollback()
		return err
	}

	var usersSegmentsHistoryBuilder, usersSegmentsBuilder strings.Builder

	if len(segmentsToAdd) != 0 {
		usersSegmentsBuilder.WriteString(fmt.Sprintf(
			"INSERT INTO %s (user_id, segment_id) VALUES ",
			usersSegmentsTable))
		usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
			"INSERT INTO %s (user_id, segment_slug, operation, updated_at) VALUES ",
			usersSegmentsHistoryTable))
		for i, seg := range segmentsToAdd {
			if i != 0 {
				usersSegmentsHistoryBuilder.WriteString(", ")
				usersSegmentsBuilder.WriteString(", ")
			}
			usersSegmentsHistoryBuilder.WriteString(
				fmt.Sprintf("('%d', '%s', '%s', 'now()')", userId, seg.Slug, "add"))
			usersSegmentsBuilder.WriteString(
				fmt.Sprintf("('%d', '%d')", userId, seg.Id))
		}

		query := usersSegmentsHistoryBuilder.String()
		_, err = repo.db.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}

		query = usersSegmentsBuilder.String()
		_, err = repo.db.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(segmentsToRemove) == 0 {
		return tx.Commit()
	}
	usersSegmentsHistoryBuilder.Reset()
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_slug, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	for i, seg := range segmentsToRemove {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%s', '%s', 'now()')", userId, seg.Slug, "delete"))
	}
	query := usersSegmentsHistoryBuilder.String()
	_, err = repo.db.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	var segmentsToRemoveIds []uint
	for _, el := range segmentsToRemove {
		segmentsToRemoveIds = append(segmentsToRemoveIds, el.Id)
	}
	segmentsToRemoveString := printSliceByComma(segmentsToRemoveIds)
	query = fmt.Sprintf("DELETE FROM %s AS us WHERE us.segment_id IN (%s)", usersSegmentsTable, segmentsToRemoveString)
	_, err = repo.db.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repo *UserRepository) GetUserSegments(userId uint) ([]string, error) {
	var slugs []string
	query := fmt.Sprintf(
		`SELECT s.slug FROM %s AS us 
                JOIN %s AS s ON s.id=us.segment_id 
            	WHERE us.user_id=$1`, usersSegmentsTable, segmentsTable)
	err := repo.db.Select(&slugs, query, userId)
	return slugs, err
}

func printSliceByComma[T uint | string](slice []T) string {
	var buf strings.Builder
	if len(slice) == 0 {
		return "''"
	}
	buf.WriteString(fmt.Sprintf("'%v'", slice[0]))
	for _, el := range slice[1:] {
		buf.WriteString(fmt.Sprintf(", '%v'", el))
	}
	return buf.String()
}
