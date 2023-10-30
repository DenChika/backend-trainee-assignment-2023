package services

import "backend-trainee-assignment-2023/pkg/repository"

type segmentService struct {
	repo repository.Segment
}

func newSegmentService(repo repository.Segment) *segmentService {
	return &segmentService{repo: repo}
}

func (service *segmentService) Create(slug string) (uint, error) {
	return service.repo.Create(slug)
}

func (service *segmentService) Delete(slug string) error {
	return service.repo.Delete(slug)
}
