package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrganizationPolicy struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Policy         datatypes.JSON `gorm:"type:jsonb;default:'{}'"`
	CreatedAt      string
	UpdatedAt      string
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
