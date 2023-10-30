package repository

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Segment interface {
	Create(slug string) (uint, error)
	Delete(slug string) error
}

type UsersSegment interface {
	ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error)
	GetUserSegments(userId uint) ([]string, error)
}

type Authorization interface {
	CreateUser(username, password string) error
	GetUser(username, password string) (uint, error)
}

type Repository struct {
	Segment
	UsersSegment
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment:       newSegmentRepository(db),
		UsersSegment:  newUsersSegmentsRepository(db),
		Authorization: newAuthorizationRepository(db),
	}
}
