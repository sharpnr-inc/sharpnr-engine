package employee_models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ POLICY DEFINITION (The Permission Group)
// -------------------------------
type EmployeePolicy struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Code: e.g., "KITCHEN_MANAGER_ACCESS"
	Code string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name string `gorm:"type:varchar(100)"`

	Description string `gorm:"type:text"`

	// Type: SYSTEM (read-only defaults), CUSTOM (created by org)
	Type string `gorm:"type:varchar(30);default:'CUSTOM'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeePolicy) TableName() string {
	return "employee.employee_policy"
}

// -------------------------------
// 2️⃣ POLICY RULES (The Granular Logic)
// -------------------------------
type EmployeePolicyRule struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeePolicyID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Scope: BAR, RESTAURANT, HOTEL, HR
	Module string `gorm:"type:varchar(50);not null"`
	// Resource: ORDER, MENU, BOOKING, BILLING
	Resource string `gorm:"type:varchar(50);not null"`
	// Action: VIEW, CREATE, UPDATE, DELETE, VOID, REFUND
	Action string `gorm:"type:varchar(50);not null"`

	// Effect: ALLOW or DENY
	Effect string `gorm:"type:varchar(10);default:'ALLOW'"`

	CreatedAt time.Time
}

func (EmployeePolicyRule) TableName() string {
	return "employee.employee_policy_rule"
}

// -------------------------------
// 3️⃣ POLICY ASSIGNMENT (Linking Policy to Staff)
// -------------------------------
type EmployeePrivilegePolicy struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID       uuid.UUID `gorm:"type:uuid;not null;index"`
	EmployeePolicyID uuid.UUID `gorm:"type:uuid;not null;index"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Domain         string    `gorm:"type:varchar(50)"` // HOSPITALITY, HR

	AssignedAt time.Time
	AssignedBy uuid.UUID `gorm:"type:uuid"` // Employee ID of the assigner

	CreatedAt time.Time
}

func (EmployeePrivilegePolicy) TableName() string {
	return "employee.employee_privilege_policy"
}
