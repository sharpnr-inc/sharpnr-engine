package legal_entity_repo

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

func (r *orgLegalEntityRepo) AddBankAccountInLegalEntity(
	ctx context.Context,
	request legal_entity_models.LegalEntityBankAccount,
) (bankAccountID uuid.UUID, err error) {
	panic("Unimplemented")
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) UpdateBankAccountDetails(
	ctx context.Context,
	bankAccountID uuid.UUID,
	updateReq legal_entity_contracts.UpdateLegalEntityBankAccountDetailsRequest,
) error {
	panic("Unimplemented")
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) FindBankAccountByID(
	ctx context.Context,
	entryID uuid.UUID,
) (*legal_entity_models.LegalEntityBankAccount, error) {
	panic("Unimplemented")
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) FindBankAccountsByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) ([]legal_entity_models.LegalEntityBankAccount, error) {
	panic("Unimplemented")
}
