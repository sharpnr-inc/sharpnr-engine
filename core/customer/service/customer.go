package org_customer_service

import (
	"context"
	"errors"
	org_customer_contracts "sharpnr-engine/core/customer/contracts"
	org_customer_models "sharpnr-engine/core/customer/models"
	dbtx "sharpnr-engine/core/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgCustomerService) Search(ctx context.Context, orgID uuid.UUID, query string) ([]org_customer_models.Customer, error) {
	// Basic validation: don't hit the DB for empty strings
	if len(query) < 2 {
		return []org_customer_models.Customer{}, nil
	}

	return s.customerRepo.SearchCustomers(ctx, orgID, query)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgCustomerService) AddCustomer(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	req org_customer_contracts.AddCustomerRequestDto,
) error {
	// 1. Availability Check (Before starting the transaction to save resources)
	isTaken, err := s.customerRepo.IsMobileTaken(ctx, orgID, req.Mobile)
	if err != nil {
		return err
	}
	if isTaken {
		return errors.New("a customer with this mobile number already exists in your organization")
	}

	// 2. Execute within Transaction
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		// Inject the current transaction into the context for the repo to extract
		txCtx := dbtx.InjectTx(ctx, tx)

		// 3. Prepare Customer Model
		customer := &org_customer_models.Customer{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           *req.Name,  // *string (can be nil)
			Email:          *req.Email, // *string (can be nil)
			Phone:          req.Mobile, // string (required)
		}

		// 4. Create Customer
		if err := s.customerRepo.CreateCustomer(txCtx, customer); err != nil {
			return err
		}

		// 5. Prepare Metadata Model
		meta := &org_customer_models.CustomerOrgMeta{
			ID:         uuid.New(),
			CustomerID: customer.ID,
			AddedBy:    emplID,
			// CreatedAt is left null/empty per your instructions
		}

		// 6. Create Metadata
		if err := s.customerRepo.CreateCustomerMeta(txCtx, meta); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgCustomerService) UpdateCustomer(
	ctx context.Context,
	orgID, emplID, customerID uuid.UUID,
	req org_customer_contracts.UpdateCustomerRequestDto,
) error {
	// 1. Build the update map
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Mobile != nil {
		// Optional: Check if the new mobile is already taken by another customer
		isTaken, err := s.customerRepo.IsMobileTaken(ctx, orgID, *req.Mobile)
		if err != nil {
			return err
		}
		if isTaken {
			return errors.New("this mobile number is already assigned to another customer")
		}
		updates["phone"] = *req.Mobile
	}

	// 2. If nothing was sent, just return
	if len(updates) == 0 {
		return nil
	}

	// 3. Execute Transaction
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 4. Update the Customer
		if err := s.customerRepo.UpdateCustomer(txCtx, orgID, customerID, updates); err != nil {
			return err
		}

		// 5. Update Metadata 'UpdatedAt' or 'ModifiedBy' if needed
		// Since you mentioned future modules (Bar/Kitchen),
		// tracking who last modified a customer is helpful.

		return nil
	})
}
