// This is the organization 's Customer
// Customer belongs to the top Organization Which Can be used accross the
// Mutiple Business usints Shring the Detials inside the Organization
package org_customer_models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Email          string
	Phone          string
	GSTIN          *string
}

func (Customer) TableName() string {
	return "customers.customer"
}

type CustomerDetails struct {
	ID             uuid.UUID
	CustomerID     uuid.UUID // Relation to the Customer Table
	SecondaryEmail string
	SecondaryPhone string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (CustomerDetails) TableName() string {
	return "customers.customer_details"
}

type CustomerOrgMeta struct {
	ID         uuid.UUID
	CustomerID uuid.UUID // Relation to the Customer Table
	AddedBy    uuid.UUID // this is the uuid of employee who Created this Customer
	CreatedAt  string    // Name of Business unit where is Was Created Restaurant, Hotel or Bar
}

func (CustomerOrgMeta) TableName() string {
	return "customers.customer_org_meta"
}
