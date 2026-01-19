package org_signup_dtos

type OrganizationRegistrationDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required,len=10,numeric"`
	Password string `json:"password" binding:"required,min=8"`
}

type VerifyOtpRequestDTO struct {
	Otp string `json:"otp" binding:"required,len=6,numeric"`
}
