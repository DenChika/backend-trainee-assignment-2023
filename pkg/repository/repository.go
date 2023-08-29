package repository

import "github.com/jmoiron/sqlx"

type Segment interface {
	Create(slug string) (int, error)
	Delete(slug string) error
}

type User interface {
	AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId uint) error
	GetUserSegments(userId uint) ([]string, error)
}

type Repository struct {
	Segment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment: NewSegmentRepository(db),
		User:    NewUserRepository(db),
	}
}
