package signupcontrollers

import "sharpnr-engine/core/interfaces"

type SignUpController struct {
	service interfaces.OrganizationSignUpService
}

// NewSignUpController creates and returns a new instance of SignUpController
func NewSignUpController(service interfaces.OrganizationSignUpService) *SignUpController {
	return &SignUpController{
		service: service,
	}
}
