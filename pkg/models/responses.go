package models

type CreateSegmentResponse struct {
	Id uint
}

type GetUserSegmentsResponse struct {
	Slugs []string
}

type ManageUserToSegmentsResponse struct {
	SlugsHaveBeenAdded   []string `json:"slugs-that-have-been-added"`
	SlugsHaveBeenRemoved []string `json:"slugs-that-have-been-removed"`
}

type SignInResponse struct {
	Token string
}
