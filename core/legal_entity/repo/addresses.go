package legal_entity_repo

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (o *orgLegalEntityRepo) CreateAddressOfLegalEntity(ctx context.Context, address *legal_entity_models.LegalEntityAddress) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	// .Create inserts the new address record.
	// Ensure address.LegalEntityID is set before calling this to link it correctly.
	err := db.WithContext(ctx).Create(address).Error

	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (o *orgLegalEntityRepo) GetAddressesOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityAddress, error) {
	// 1. Initialize a slice to hold multiple results
	var addresses []legal_entity_models.LegalEntityAddress

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}

	// 2. Use .Find() instead of .First()
	// .Find() populates the slice and does not return ErrRecordNotFound
	// if no addresses are found (it just returns an empty slice).
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&addresses).Error

	if err != nil {
		return nil, err
	}

	// 3. Return the populated slice
	return addresses, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (o *orgLegalEntityRepo) UpdateLegalEntityAddress(
	ctx context.Context,
	legalEntityID uuid.UUID,
	addressID uuid.UUID,
	req legal_entity_contracts.UpdateLegalEntityAddressRequestDto,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}

	// 1. We use .Model() to specify the target table
	// 2. We filter by both addressID and legalEntityID for security (ownership check)
	// 3. .Updates(req) only updates the fields where the pointer is not nil
	result := db.WithContext(ctx).
		Model(&legal_entity_models.LegalEntityAddress{}).
		Where("id = ? AND legal_entity_id = ?", addressID, legalEntityID).
		Updates(req)

	// Handle DB execution errors
	if result.Error != nil {
		return result.Error
	}

	// Check if the record was actually found and modified
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
