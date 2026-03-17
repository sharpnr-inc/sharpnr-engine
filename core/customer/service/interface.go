package org_customer_service

import (
	"context"
	org_customer_contracts "sharpnr-engine/core/customer/contracts"
	org_customer_models "sharpnr-engine/core/customer/models"

	"github.com/google/uuid"
)

type OrgCustomerServiceInterface interface {
	Search(ctx context.Context, orgID uuid.UUID, query string) ([]org_customer_models.Customer, error)
	AddCustomer(ctx context.Context, orgID, emplID uuid.UUID, req org_customer_contracts.AddCustomerRequestDto) error
	UpdateCustomer(ctx context.Context, orgID, emplID, customerID uuid.UUID, updates org_customer_contracts.UpdateCustomerRequestDto) error
}
