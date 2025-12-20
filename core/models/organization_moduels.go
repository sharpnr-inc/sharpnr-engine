package models

import (
	"github.com/google/uuid"
)

type OrganizationBusinessLogicSubModule struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	Organization   Organization
	SubModuleID    uuid.UUID `gorm:"type:uuid;not null"`
	SubModuleName  string    `gorm:"size:100"`
	IsActive       bool      `gorm:"default:true"`
}
