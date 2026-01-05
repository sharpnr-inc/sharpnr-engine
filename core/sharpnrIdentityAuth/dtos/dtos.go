package sharpnr_identity_auth_dtos

type LoginRequest_dto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type WebLoginResponse_dto struct {
	Email      string
	employeeID string
	Name       string
}

type MobileLoginRequest_dto struct {
}
