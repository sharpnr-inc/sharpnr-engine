package accounting_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *accountingServices) GetLegalEntitySuppliers(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]accounting_engine_contracts.LegalEntitySuppliersResponseDto, error) {

	var results []accounting_engine_contracts.LegalEntitySuppliersResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		suppliersOfLegalEntity, err := s.supplierService.
			ListAllSuppliersOfLegalEntity(txCtx, orgID, emplID, legalEntityID)
		if err != nil {
			return err
		}

		results = make(
			[]accounting_engine_contracts.LegalEntitySuppliersResponseDto,
			0,
			len(suppliersOfLegalEntity),
		)

		for _, supplier := range suppliersOfLegalEntity {
			results = append(results,
				accounting_engine_contracts.LegalEntitySuppliersResponseDto{
					SuppliersID:   supplier.ID,
					SuppliersName: supplier.Name,
					GSTNumber:     supplier.GSTNumber,
					PANNumber:     supplier.PANNumber,
					Phone:         supplier.Phone,
					Email:         supplier.Email,
					AddressLine:   supplier.AddressLine,
					City:          supplier.City,
					State:         supplier.State,
					Pincode:       supplier.Pincode,
					Country:       supplier.Country,
					CreatedAt:     supplier.CreatedAt,
					UpdatedAt:     supplier.UpdatedAt,
				},
			)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingServices) AddNewSuppliers(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	request *accounting_engine_contracts.AddNewSuppliersForLegalEntityRequestDto,
) error {
	panic("")
}
