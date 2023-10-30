package services

import (
	"backend-trainee-assignment-2023/pkg/models"
	"backend-trainee-assignment-2023/pkg/repository"
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
	SignUp(username, password string) error
	SignIn(username, password string) (string, error)
}

type Service struct {
	Segment
	UsersSegment
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Segment:       newSegmentService(repo.Segment),
		UsersSegment:  newUsersSegmentService(repo.UsersSegment),
		Authorization: newAuthorizationService(repo.Authorization),
	}
}
