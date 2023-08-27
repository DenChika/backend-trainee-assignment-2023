package repository

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId int) error {
	return nil
}
func (repo *UserRepository) GetUserSegments(userId int) ([]string, error) {
	return nil, nil
}
