package models

import "gorm.io/datatypes"

type OrgUIShape struct {
	ID             string
	OrganizationId string
	UI_Shape       datatypes.JSON
}
