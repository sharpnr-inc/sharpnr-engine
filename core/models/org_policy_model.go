package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrganisationPolicy struct {
	ID             uuid.UUID
	OrganisationID uuid.UUID
	Policy         datatypes.JSON `gorm:"type:jsonb;default:'{}'"`
	CreatedAt      string
	UpdatedAt      string
	Organization   Organization `gorm:"foreignKey:OrganisationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
