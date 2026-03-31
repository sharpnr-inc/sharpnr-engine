package suppliers_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

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
