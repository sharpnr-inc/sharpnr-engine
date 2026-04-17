package suppliers_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	supplier_contracts "sharpnr-engine/core/org_supplier/contracts"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *supplierService) CreateNewSupplier(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID,
	request *suppliers_dtos.CreateSuppliersAPIRequestdto,
) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		supplier := &org_suppliers_models.Supplier{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           request.SuppliersName,
			Phone:          request.Phone,
			Email:          request.Email,
			AddressLine:    request.AddressLine,
			Pincode:        request.Pincode,
			Country:        request.Country,
			City:           request.City,
			State:          request.State,
		}

		if err := s.supplierRepo.CreateSupplier(txCtx, supplier); err != nil {
			return err
		}

		return nil

	})
}

func (s *supplierService) GetAllSupplierByOrgID(ctx context.Context, orgID, emplID uuid.UUID) ([]org_suppliers_models.Supplier, error) {
	// 1. (Optional) Authorization Check
	// You can use emplID here to verify if the user has permission
	// to view suppliers for this specific orgID.

	// 2. Call the repository to fetch the data
	suppliers, err := s.supplierRepo.GetSuppliersByOrgID(ctx, orgID)
	if err != nil {
		// Log error and return a wrapped error if necessary
		return nil, err
	}

	// 3. Return the results
	return suppliers, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *supplierService) ListAllSuppliersOfLegalEntity(
	ctx context.Context,
	orgID,
	emplID,
	legalEntityID uuid.UUID,
) ([]org_suppliers_models.Supplier, error) {

	var results []org_suppliers_models.Supplier

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		suppliers, err := s.supplierRepo.GetSuppliersByLegalEntity(
			ctx,
			orgID,
			legalEntityID,
		)
		if err != nil {
			return err
		}

		results = suppliers
		return nil
	})

	if err != nil {
		return results, err
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *supplierService) ListAllLegalEntityLinksOfSuppliers(ctx context.Context, orgID, emplID, suppliersID uuid.UUID) ([]org_suppliers_models.SupplierLegalEntityLink, error) {
	var links []org_suppliers_models.SupplierLegalEntityLink

	// Wrap in a transaction to ensure a consistent view of the mapping
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		// Inject the transaction into the context for the repository
		txCtx := dbtx.InjectTx(ctx, tx)

		// Fetch the links associated with this specific supplier
		// Note: Ensure your repository method supports filtering by SupplierID
		result, err := s.supplierRepo.GetSuppliersLinkOfSuppliers(txCtx, suppliersID)
		if err != nil {
			return err
		}

		links = result
		return nil
	})

	if err != nil {
		// Return nil on failure to signal that data is untrusted
		return nil, err
	}

	return links, nil
}

func (s *supplierService) AddLegalEntityLinksToSupplier(ctx context.Context, orgID, emplID uuid.UUID, req supplier_contracts.AddLegalEntityLinkInSuppliersRequestDto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		link := &org_suppliers_models.SupplierLegalEntityLink{
			ID:            uuid.New(),
			SupplierID:    req.SuppliersID,
			LegalEntityID: req.LegalEntityID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := s.supplierRepo.CreateSupplierLegalEntityLink(txCtx, link); err != nil {
			return err
		}

		return nil
	})
}
