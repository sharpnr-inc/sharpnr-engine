package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrgUIRendererTree struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null"`
	UI_Render_tree datatypes.JSON `gorm:"type:jsonb;default:'{}'"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}
