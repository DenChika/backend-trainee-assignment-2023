package services

import "backend-trainee-assignment-2023/pkg/repository"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) AddUserToSegment(slugsToAdd []string, slugsToRemove []string, userId int) error {
	return service.repo.AddUserToSegment(slugsToAdd, slugsToRemove, userId)
}
func (service *UserService) GetUserSegments(userId int) ([]string, error) {
	return service.repo.GetUserSegments(userId)
}
