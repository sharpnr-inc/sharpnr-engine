package legal_entity_repo

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgLegalEntityRepo) AddBankAccountInLegalEntity(
	ctx context.Context,
	request legal_entity_models.LegalEntityBankAccount,
) (bankAccountID uuid.UUID, err error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Use WithContext to ensure telemetry, tracing, and cancellations are respected
	if err := db.WithContext(ctx).Create(&request).Error; err != nil {
		return uuid.Nil, err
	}

	// Assuming your model struct uses 'ID' or specifies the UUID field
	// GORM automatically populates the struct field after a successful insert
	return request.ID, nil

}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) UpdateBankAccountDetails(
	ctx context.Context,
	bankAccountID uuid.UUID,
	updateReq legal_entity_contracts.UpdateLegalEntityBankAccountDetailsRequest,
) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Explicitly target the model type and narrow down the scope with a Where clause
	result := db.WithContext(ctx).
		Model(&legal_entity_models.LegalEntityBankAccount{}).
		Where("id = ?", bankAccountID).
		Updates(updateReq)

	if result.Error != nil {
		return result.Error
	}

	// Optional: If you want to return an error when the record doesn't exist
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) FindBankAccountByID(
	ctx context.Context,
	ID uuid.UUID,
) (*legal_entity_models.LegalEntityBankAccount, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var bankAccount legal_entity_models.LegalEntityBankAccount

	// First will return gorm.ErrRecordNotFound if the ID doesn't exist
	err := db.WithContext(ctx).
		First(&bankAccount, "id = ?", ID).
		Error

	if err != nil {
		return nil, err
	}

	return &bankAccount, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) FindBankAccountsByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) ([]legal_entity_models.LegalEntityBankAccount, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var bankAccounts []legal_entity_models.LegalEntityBankAccount

	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&bankAccounts).
		Error

	if err != nil {
		return nil, err
	}

	// If the database query succeeded but found 0 rows, explicitly return the GORM error
	if len(bankAccounts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return bankAccounts, nil
}
