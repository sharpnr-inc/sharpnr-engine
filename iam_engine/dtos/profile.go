package iamEngineDtos

type ProfileResponse_dto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
}
