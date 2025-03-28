package repository

import (
	"backend-trainee-assignment-2023/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type usersSegmentsRepository struct {
	db *sqlx.DB
}

type operation string

const (
	add    operation = "add"
	delete operation = "delete"
)

func newUsersSegmentsRepository(db *sqlx.DB) *usersSegmentsRepository {
	return &usersSegmentsRepository{db: db}
}

func (repo *usersSegmentsRepository) ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error) {
	tx := NewTransaction(repo.db.MustBegin())

	_, err := tx.Exec(fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", "REPEATABLE READ"))
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed setting transaction isolation level: %s\n", err.Error())
		return nil, err
	}

	segmentsToAddIds, err := tx.filterIfSegmentsExist(slugsToAdd)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed filtering existing segments for adding: %s\n", err.Error())
		return nil, err
	}
	segmentsToRemoveIds, err := tx.filterIfSegmentsExist(slugsToRemove)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed filtering existing segments for removing: %s\n", err.Error())
		return nil, err
	}
	segmentsToAddIds, err = tx.filterIfUsersSegmentsNotInDb(segmentsToAddIds)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed filtering existing users segments: %s\n", err.Error())
		return nil, err
	}
	if len(segmentsToAddIds) != 0 {
		if err = tx.insertSegmentsIntoUser(segmentsToAddIds, userId); err != nil {
			_ = tx.Rollback()
			logrus.Fatalf("failed inserting segments into user: %s\n", err.Error())
			return nil, err
		}
		if err = tx.saveInHistory(segmentsToAddIds, userId, add); err != nil {
			_ = tx.Rollback()
			logrus.Fatalf("failed saving segments for adding to history: %s\n", err.Error())
			return nil, err
		}
	}

	segmentsToRemoveIds, err = tx.deleteSegmentsFromUser(segmentsToRemoveIds, userId)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed deleting segments from user: %s\n", err.Error())
		return nil, err
	}

	if len(segmentsToRemoveIds) != 0 {
		if err = tx.saveInHistory(segmentsToRemoveIds, userId, delete); err != nil {
			_ = tx.Rollback()
			logrus.Fatalf("failed saving segments for removing to history: %s\n", err.Error())
			return nil, err
		}
	}
	slugsHaveBeenAdded, err := tx.findSlugsByIds(segmentsToAddIds)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed finding slugs for adding by ids: %s\n", err.Error())
		return nil, err
	}
	slugsHaveBeenRemoved, err := tx.findSlugsByIds(segmentsToRemoveIds)
	if err != nil {
		_ = tx.Rollback()
		logrus.Fatalf("failed finding slugs for removing by ids: %s\n", err.Error())
		return nil, err
	}
	return &models.ManageUserToSegmentsResponse{
		SlugsHaveBeenAdded:   slugsHaveBeenAdded,
		SlugsHaveBeenRemoved: slugsHaveBeenRemoved,
	}, tx.Commit()
}

func (repo *usersSegmentsRepository) GetUserSegments(userId uint) ([]string, error) {
	var slugs []string
	query := fmt.Sprintf(
		`SELECT s.slug FROM %s AS us 
                JOIN %s AS s ON s.id=us.segment_id 
            	WHERE us.user_id=$1`, usersSegmentsTable, segmentsTable)
	if err := repo.db.Select(&slugs, query, userId); err != nil {
		logrus.Fatalf("failed finding user segments: %s\n", err.Error())
		return nil, err
	}
	return slugs, nil
}
