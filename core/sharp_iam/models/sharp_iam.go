package sharp_iam_models

import (
	"time"

	"github.com/google/uuid"
)

type SharpIamPolicies struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrgID        uuid.UUID `gorm:"index;not null"`
	Domain       string    `gorm:"size:50;index;not null"`  // Hospitality, Medical, etc.
	PrincipalID  string    `gorm:"size:100;index;not null"` // EmployeeID, RoleID, etc. depending on PrincipalType
	ResourceType string    `gorm:"size:50;index;not null"`
	ResourceID   string    `gorm:"size:100;index;not null"`
	ScopeID      string    `gorm:"size:100;index"` // For Property or LegalEntity level scoping
	CedarPolicy  string    `gorm:"type:text;not null"`
	UIMetadata   []byte    `gorm:"type:jsonb"` // To remember checkbox states
	IsEnabled    bool      `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SharpIamRoles struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrgID  uuid.UUID `gorm:"index;not null"`
	Domain string    `gorm:"size:50"`

	RoleName    string `gorm:"size:50;index;not null"`
	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SharpIamEmployeeRoles struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrgID uuid.UUID `gorm:"index;not null"`

	EmployeeID      string    `gorm:"size:100;index;not null"`
	SharpIamRolesID uuid.UUID `gorm:"type:uuid;index;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
