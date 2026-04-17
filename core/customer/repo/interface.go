package org_customer_repo

import (
	"context"
	org_customer_models "sharpnr-engine/core/customer/models"

	"github.com/google/uuid"
)

type OrgCustomerRepoInterface interface {
	CreateCustomer(ctx context.Context, customer *org_customer_models.Customer) error
	CreateCustomerMeta(ctx context.Context, customerMeta *org_customer_models.CustomerOrgMeta) error
	GetCustomersByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_customer_models.Customer, error)
	GetCustomerByEmail(ctx context.Context, customerEmail string) (*org_customer_models.Customer, error)
	SearchCustomers(
		ctx context.Context,
		orgID uuid.UUID,
		query string,
	) ([]org_customer_models.Customer, error)

	UpdateCustomer(
		ctx context.Context,
		orgID, customerID uuid.UUID,
		updates interface{},
	) error

	IsMobileTaken(ctx context.Context, orgID uuid.UUID, mobile string) (bool, error)
}
