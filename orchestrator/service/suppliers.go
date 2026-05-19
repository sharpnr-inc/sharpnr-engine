package orchestrator_service

import (
	"context"
	"errors"
	supplier_contracts "sharpnr-engine/core/org_supplier/contracts"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

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
					LegalEntityID:   leDetails.ID,
					LegalEntityName: leDetails.DisplayName,
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

/* ---------------------------------- <***> --------------------------------- */

// This Service method Connects a supplier to a legal entity by creating a link between them. It first checks for existing links to prevent duplicates,

func (s *orchestratorService) AddLegalEntityLinksInSuppliers(ctx context.Context, orgID, emplID, suppliersID uuid.UUID, request *supplier_contracts.AddLegalEntityLinkInSuppliersRequestDto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Use the 'suppliersID' from the URL/Params, not the request body
		existingLinks, err := s.supplierService.ListAllLegalEntityLinksOfSuppliers(txCtx, orgID, emplID, suppliersID)
		if err != nil {
			return err // Return the actual error so the transaction rolls back
		}

		// 2. Check for duplicates
		for _, link := range existingLinks {
			if link.LegalEntityID == request.LegalEntityID {
				// Return a specific error if the link already exists
				return errors.New("this legal entity is already linked to this supplier")
			}
		}

		// 3. Ensure the request object has the correct SupplierID before passing it down
		// This is why your DB was getting 0000-0000
		request.SuppliersID = suppliersID

		if err := s.supplierService.AddLegalEntityLinksToSupplier(txCtx, orgID, emplID, *request); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */
