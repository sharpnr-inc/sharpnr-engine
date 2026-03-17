// This is the organization 's Customer
// Customer belongs to the top Organization Which Can be used accross the
// Mutiple Business usints Shring the Detials inside the Organization
package org_customer_models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	email          string
	Phone          string
	GSTIN          *string
	gorm.Model
}

func (Customer) TableName() string {
	return "customers.customer"
}
