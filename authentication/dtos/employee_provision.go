package authenticationEngineDtos

type VerifyEmployeeProvisionRequestDto struct {
	Token     string `json:"token" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
}
