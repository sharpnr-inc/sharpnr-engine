package accounting_service

import (
	"context"
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
		// txCtx := dbtx.InjectTx(ctx, tx)

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
