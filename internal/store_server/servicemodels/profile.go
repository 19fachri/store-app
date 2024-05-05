package servicemodels

type ProfileBase struct {
	Name string `json:"name"`
}

type ProfileCreateRequest struct {
	ProfileBase
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ProfileUpdateRequest = ProfileBase
