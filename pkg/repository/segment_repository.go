package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type segmentRepository struct {
	db *sqlx.DB
}

type SegmentEntity struct {
	Id   uint
	Slug string
}

func newSegmentRepository(db *sqlx.DB) *segmentRepository {
	return &segmentRepository{db: db}
}

func (repo *segmentRepository) Create(slug string) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (slug) VALUES ($1) RETURNING id", segmentsTable)
	var id uint
	row := repo.db.QueryRow(query, slug)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *segmentRepository) Delete(slug string) error {
	tx := NewTransaction(repo.db.MustBegin())
	queryDelete := fmt.Sprintf("DELETE FROM %s AS s WHERE s.slug=$1", segmentsTable)
	_, err := tx.Exec(queryDelete, slug)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	queryLastOperation := fmt.Sprintf(
		`SELECT ush.user_id, ush.segment_slug FROM %s AS ush
                JOIN %s AS s ON ush.segment_slug=s.slug
                WHERE s.slug=$1 AND ush.operation::TEXT LIKE 'add'`,
		usersSegmentsHistoryTable, segmentsTable)
	type saveRequest struct {
		UserId uint   `json:"user_id" db:"user_id"`
		Slug   string `json:"segment_slug" db:"segment_slug"`
	}
	var requests []saveRequest
	err = tx.Select(&requests, queryLastOperation, slug)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	log.Printf("ab: %v", requests)
	if len(requests) == 0 {
		return tx.Commit()
	}
	var usersSegmentsHistoryBuilder strings.Builder
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_slug, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	for i, req := range requests {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%s', '%s', 'now()')", req.UserId, req.Slug, "delete"))
	}

	query := usersSegmentsHistoryBuilder.String()
	log.Printf("abobus: %s", query)
	_, err = tx.Exec(query)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
