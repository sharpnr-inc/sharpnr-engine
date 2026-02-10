package sharp_iam_models

import (
	"time"

	"github.com/google/uuid"
)

type SharpPolicy struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OrgID  uuid.UUID `gorm:"index;not null"`
	Domain string    `gorm:"size:50;index;not null"` // Hospitality, Medical, etc.

	PrincipalType string `gorm:"size:20;index;default:'User'"`

	PrincipalID  string `gorm:"size:100;index;not null"`
	ResourceType string `gorm:"size:50;index;not null"`
	ResourceID   string `gorm:"size:100;index;not null"`
	ScopeID      string `gorm:"size:100;index"` // For Property or LegalEntity level scoping
	CedarPolicy  string `gorm:"type:text;not null"`
	IsEnabled    bool   `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
