package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type authorizationRepository struct {
	db *sqlx.DB
}

func newAuthorizationRepository(db *sqlx.DB) *authorizationRepository {
	return &authorizationRepository{db: db}
}

func (repo *authorizationRepository) CreateUser(username, password string) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2)", usersTable)
	if _, err := repo.db.Exec(query, username, password); err != nil {
		logrus.Fatalf("failed creating user: %s\n", err.Error())
		return err
	}
	return nil
}

func (repo *authorizationRepository) GetUser(username, password string) (uint, error) {
	var id uint
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	if err := repo.db.Get(&id, query, username, password); err != nil {
		logrus.Fatalf("failed getting user from db: %s\n", err.Error())
		return 0, err
	}
	return id, nil
}
