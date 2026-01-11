package org_signup_dtos

import "github.com/google/uuid"

type OrgOnBoardingRequest_dto struct {
	OrgSector        string   `json:"org_sector" binding:"required"`
	OrgSectorModules []string `json:"org_sector_modules" binding:"required,min=1"`
	OrgAddressLine   string   `json:"org_address_line" binding:"required"`
	OrgCity          string   `json:"org_city" binding:"required"`
	OrgState         string   `json:"org_state" binding:"required"`
	OrgCountry       string   `json:"org_country" binding:"required"`
	OrgPincode       string   `json:"org_pincode" binding:"required"`
}

type OrgOnBoardingSuccessResponse_dto struct {
	OrgId        uuid.UUID
	EmployeeID   uuid.UUID
	EmployeeRole string
}
