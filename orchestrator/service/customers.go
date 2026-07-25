package orchestrator_service

import (
	"context"
	org_customer_contracts "sharpnr-engine/core/customer/contracts"
	org_customer_models "sharpnr-engine/core/customer/models"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// customerOrchestrator implements CustomerOrchestratorServiceInterface.
type customerOrchestrator struct {
	txManager      dbtx.TxManager
	customerEngine org_customer_service.OrgCustomerServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *customerOrchestrator) AddCustomer(ctx context.Context, orgID, emplID uuid.UUID, req org_customer_contracts.AddCustomerRequestDto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		return s.customerEngine.AddCustomer(txCtx, orgID, emplID, req)

	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *customerOrchestrator) SearchCustomer(ctx context.Context, orgID, emplID uuid.UUID, searchQuery string) ([]org_customer_models.Customer, error) {
	var results []org_customer_models.Customer

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Fix: Use a local variable to capture the engine output
		foundCustomers, err := s.customerEngine.Search(txCtx, orgID, searchQuery)
		if err != nil {
			return err
		}

		// Fix: Assign the found data to the outer scope variable
		results = foundCustomers
		return nil
	})

	return results, err
}

/* ---------------------------------- <***> --------------------------------- */
