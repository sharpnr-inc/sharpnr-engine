package sharpnr_identity_auth_dtos

type ProfileResponse_dto struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
}
