package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

/*
"employee.employee_provisioning" or "EmployeeProvisioning"

Backs the invite link an employee gets emailed after an admin creates them:
auth.sharpnr.com/sa/{token}/{employee_email}. The frontend posts the token
back to verify it, and on success the employee sets their password — that's
when EmployeeAuth gets created.

The raw token is only ever emailed, never stored: TokenHash is sha256(token),
mirroring how EmployeeAuth never stores the raw password. ConsumedAt makes
the token single-use — once the password is set, it's marked consumed and
any later attempt with the same token is rejected even if not yet expired.
*/
type EmployeeProvisioning struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`
	Email      string    `gorm:"type:varchar(150);not null"`

	TokenHash string `gorm:"type:text;not null;uniqueIndex"`

	ExpiresAt  time.Time `gorm:"not null"`
	ConsumedAt *time.Time

	CreatedAt time.Time
}

func (EmployeeProvisioning) TableName() string {
	return "employee.employee_provisioning"
}
