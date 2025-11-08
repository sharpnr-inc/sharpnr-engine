package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type AccountGroup struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookID    uuid.UUID `gorm:"type:uuid;index"`
	Name      string    `gorm:"not null"`
	ParentID  *uuid.UUID
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
