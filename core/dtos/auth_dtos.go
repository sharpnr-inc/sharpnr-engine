package dto

type LoginRequestDTO struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	OrgId   string
	OrgType string
	Tokens  string
	Message string
}
