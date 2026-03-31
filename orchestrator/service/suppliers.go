package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orchestratorService) AddNewSupplier(ctx context.Context, orgID, emplID uuid.UUID, request *suppliers_dtos.CreateSuppliersAPIRequestdto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.supplierService.CreateNewSupplier(txCtx, orgID, emplID, request)
	})
}

func (s *orchestratorService) ListAllOrganizationSuppliers(ctx context.Context, orgID, emplID uuid.UUID) ([]org_suppliers_models.Supplier, error) {
	var results []org_suppliers_models.Supplier

	// Using a transaction for a simple READ is generally optional.
	// It is only strictly necessary if you require "Repeatable Read" isolation
	// for complex reports or consistency across multiple queries.
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		suppliers, err := s.supplierService.GetAllSupplierByOrgID(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		results = suppliers
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
