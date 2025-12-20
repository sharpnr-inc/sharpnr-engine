package models

import "github.com/google/uuid"

type OrganizationBrandingDetails struct {
	Id             string
	LogoUrl        string
	MarketingLine  string
	OrganizationId uuid.UUID
}
