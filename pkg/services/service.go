package services

import (
	"backend-trainee-assignment-2023/pkg/repository"
)

type Segment interface {
	Create(slug string) (int, error)
	Delete(slug string) error
}

type User interface {
	AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId int) error
	GetUserSegments(userId int) ([]string, error)
}

type Service struct {
	Segment
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Segment: NewSegmentService(repo.Segment),
		User:    NewUserService(repo.User),
	}
}
