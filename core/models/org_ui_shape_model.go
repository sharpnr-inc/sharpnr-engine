package models

import "gorm.io/datatypes"

type OrgUIShape struct {
	ID             string
	OrganisationId string
	UI_Shape       datatypes.JSON
}
