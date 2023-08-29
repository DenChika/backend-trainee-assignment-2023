package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId int) error {
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
	query := fmt.Sprintf(
		`SELECT s.id FROM %s AS us
            JOIN %s AS s ON s.id=us.segment_id
            WHERE s.slug IN ($1)`, usersSegmentsTable, segmentsTable)
	var segmentsToAddIds []int
	if err := repo.db.Select(&segmentsToAddIds, query, slugsToAddString); err != nil {
		tx.Rollback()
		return err
	}
	var segmentsToRemoveIds []int
	if err := repo.db.Select(&segmentsToRemoveIds, query, slugsToRemoveString); err != nil {
		tx.Rollback()
		return err
	}

	var usersSegmentsHistoryBuilder, usersSegmentsBuilder strings.Builder
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_id, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	usersSegmentsBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_id) VALUES ",
		usersSegmentsTable))
	for i, id := range segmentsToAddIds {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
			usersSegmentsBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%d', '%s', '%s')", userId, id, "add", time.Now().String()))
		usersSegmentsBuilder.WriteString(
			fmt.Sprintf("('%d', '%d')", userId, id))
	}

	query = usersSegmentsHistoryBuilder.String()
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

	usersSegmentsHistoryBuilder.Reset()
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_id, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	for i, id := range segmentsToRemoveIds {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%d', '%s', '%s')", userId, id, "delete", time.Now().String()))
	}
	query = usersSegmentsHistoryBuilder.String()
	_, err = repo.db.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	segmentsToRemoveString := printSliceByComma(segmentsToRemoveIds)
	query = fmt.Sprintf("DELETE FROM %s AS us WHERE us.segment_id IN ($1)", usersSegmentsTable)
	_, err = repo.db.Exec(query, segmentsToRemoveString)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (repo *UserRepository) GetUserSegments(userId int) ([]string, error) {
	var slugs []string
	query := fmt.Sprintf(
		`SELECT s.slug FROM %s AS us 
                JOIN %s AS s ON s.id=us.segment_id 
            	WHERE us.user_id=$1`, usersSegmentsTable, segmentsTable)
	err := repo.db.Select(&slugs, query, userId)
	return slugs, err
}

func printSliceByComma[T int | string](slice []T) string {
	var buf strings.Builder
	if len(slice) == 0 {
		return "''"
	}
	buf.WriteString(fmt.Sprintf("'%s'", slice[0]))
	for _, el := range slice[1:] {
		buf.WriteString(fmt.Sprintf(", '%s'", el))
	}
	return buf.String()
}
