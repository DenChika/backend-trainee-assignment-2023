package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
