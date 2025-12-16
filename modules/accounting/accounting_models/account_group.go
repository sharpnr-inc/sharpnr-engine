package accounting_models

import (
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
)

type AccountGroup struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Organization models.Organization

	BookID    uuid.UUID `gorm:"type:uuid;index"`
	Name      string    `gorm:"not null"`
	ParentID  *uuid.UUID
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
