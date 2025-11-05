package authcontrollers

import "sharpnr-engine/core/interfaces"

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(service interfaces.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}
