package models

import "time"

type CommissionRule struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	OrganizationID string       `gorm:"size:16;not null;index"`                            // Tenant scope
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgID"` // Relation to Organization

	ServiceID uint `gorm:"not null;index"` // Which service (e.g. Hair Color)

	EmployeeID *uint   `gorm:"index;default:null"`   // Nullable: if null applies to all employees
	Role       *string `gorm:"size:50;default:null"` // e.g. "Stylist", "Manager"

	CommissionType  string  `gorm:"type:varchar(20);not null"` // "percentage" or "fixed"
	CommissionValue float64 `gorm:"not null"`                  // e.g. 10.5 (% or fixed amount)

	CreatedAt time.Time
	UpdatedAt time.Time
}
