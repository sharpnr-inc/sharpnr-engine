package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orchestratorService) AddNewSupplier(ctx context.Context, orgID, emplID uuid.UUID, request *suppliers_dtos.CreateSuppliersAPIRequestdto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.supplierService.CreateNewSupplier(txCtx, orgID, emplID, request)
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) ListAllOrganizationSuppliers(ctx context.Context, orgID, emplID uuid.UUID) ([]suppliers_dtos.SuppliersResponseDetails, error) {
	var results []suppliers_dtos.SuppliersResponseDetails

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Get all base supplier records
		suppliers, err := s.supplierService.GetAllSupplierByOrgID(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		for _, sp := range suppliers {
			// Initialize the detail DTO for this specific supplier
			supplierDetail := suppliers_dtos.SuppliersResponseDetails{
				ID:            sp.ID,
				SuppliersName: sp.Name,
				Email:         sp.Email,
				Phone:         sp.Phone,
				AddressLine:   sp.AddressLine,
				City:          sp.City,
				State:         sp.State,
				Pincode:       sp.Pincode,
				Country:       sp.Country,
			}

			// 2. Fetch all legal entity links for this supplier
			legalEntityLinks, err := s.supplierService.ListAllLegalEntityLinksOfSuppliers(txCtx, orgID, emplID, sp.ID)
			if err != nil {
				return err
			}

			for _, leLink := range legalEntityLinks {
				// 3. Fetch details for the specific legal entity
				leDetails, err := s.legalEntityService.GetLegalEntityDetails(txCtx, orgID, emplID, leLink.LegalEntityID)
				if err != nil {
					return err
				}

				// Append the legal entity info to the supplier's link list
				supplierDetail.LegalEntityLinks = append(supplierDetail.LegalEntityLinks, suppliers_dtos.SuppliersLegalEntity{
					LegalEntityID:   leDetails.LegalEntityID,
					LegalEntityName: leDetails.Name,
				})
			}

			// Add the fully populated supplier detail to our final results
			results = append(results, supplierDetail)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
