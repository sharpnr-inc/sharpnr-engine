package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrgUIRendererTree struct {
	ID             uuid.UUID      `gorm:"primaryKey"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null"`
	UIRendererTree datatypes.JSON `gorm:"type:jsonb;default:'{}'"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}
