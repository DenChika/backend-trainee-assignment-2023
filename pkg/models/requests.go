package models

type SegmentRequest struct {
	Slug string `validate:"required"`
}

type ManageUserToSegmentsRequest struct {
	SlugsToAdd    []string `json:"slugs-to-add" validate:"required"`
	SlugsToRemove []string `json:"slugs-to-remove" validate:"required"`
}

type AuthRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
