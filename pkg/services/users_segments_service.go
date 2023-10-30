package services

import (
	"backend-trainee-assignment-2023/pkg/models"
	"backend-trainee-assignment-2023/pkg/repository"
)

type usersSegmentsService struct {
	repo repository.UsersSegment
}

func newUsersSegmentService(repo repository.UsersSegment) *usersSegmentsService {
	return &usersSegmentsService{repo: repo}
}

func (service *usersSegmentsService) ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error) {
	return service.repo.ManageUserToSegments(slugsToAdd, slugsToRemove, userId)
}
func (service *usersSegmentsService) GetUserSegments(userId uint) ([]string, error) {
	return service.repo.GetUserSegments(userId)
}
