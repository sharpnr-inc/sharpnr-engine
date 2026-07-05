package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

/*
"employee.employee_auth" or "EmployeeAuth"
This Models is responsible for storing the authentication details of the employee

Relationships:
- EmployeeAuth has a one-to-one relationship with EmployeeIdentity (EmployeeID is a foreign key to EmployeeIdentity.ID)
- EmployeeAuth is used for authentication and authorization purposes, it stores the password hash and other related information for the employee. It also has fields to track the last login time and whether the employee is active or not.

In this table we have the username which is in the organization context not in the global context, this is because the employee can have different usernames in different organizations, but the email is unique across the system and can be used for authentication.
*/
type EmployeeAuth struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	Username     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email        string `gorm:"type:varchar(150);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:text;not null"`

	IsActive      bool `gorm:"default:true"`
	IsProvisioned bool `gorm:"default:false"` // True once IT setup is complete

	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (EmployeeAuth) TableName() string {
	return "employee.employee_auth"
}
