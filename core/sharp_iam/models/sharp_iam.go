package sharp_iam_models

import (
	"time"

	"github.com/google/uuid"
)

type SharpIamPolicies struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrganizationID uuid.UUID `gorm:"index;not null"`
	Domain         string    `gorm:"size:50"`                 // Hospitality, Medical, etc.
	PrincipalID    string    `gorm:"size:100;index;not null"` // EmployeeID, RoleID, etc. depending on PrincipalType
	ResourceType   string    `gorm:"size:50;index;not null"`
	ResourceID     string    `gorm:"size:100;index;not null"`
	ScopeID        string    `gorm:"size:100;index"` // For Property or LegalEntity level scoping
	CedarPolicy    string    `gorm:"type:text;not null"`
	UIMetadata     []byte    `gorm:"type:jsonb"` // To remember checkbox states
	IsEnabled      bool      `gorm:"default:true"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (SharpIamPolicies) TableName() string {
	return "policy_engine.sharp_iam_policies"
}

type SharpIamRoles struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrganizationID uuid.UUID `gorm:"index;not null"`
	Domain         string    `gorm:"size:50"`

	RoleName    string `gorm:"size:50;index;not null"`
	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SharpIamRoles) TableName() string {
	return "policy_engine.sharp_iam_roles"
}

type SharpIamEmployeeRoles struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	OrganizationID uuid.UUID `gorm:"index;not null"`

	EmployeeID      string    `gorm:"size:100;index;not null"`
	SharpIamRolesID uuid.UUID `gorm:"type:uuid;index;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SharpIamEmployeeRoles) TableName() string {
	return "policy_engine.sharp_iam_employee_roles"
}
