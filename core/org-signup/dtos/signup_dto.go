package org_signup_dtos

type InitiateOrgSignUpRequest_Dto struct {
	OrgName     string `json:"org_name"`
	OrgEmail    string `json:"org_email" binding:"required,email"`
	OrgMobile   string `json:"org_mobile" binding:"required"`
	OrgPassword string `json:"org_password" binding:"required,min=8"`
}

type VerifyOtpRequest_Dto struct {
	Otp string `json:"otp" binding:"required, otp"`
}
