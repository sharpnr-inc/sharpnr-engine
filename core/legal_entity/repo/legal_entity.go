package legal_entity_repo

import (
	"context"
	"fmt"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (o *orgLegalEntityRepo) CreateLegalEntity(ctx context.Context, legalEntity *legal_entity_models.LegalEntity) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	if err := db.WithContext(ctx).Create(legalEntity).Error; err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (o *orgLegalEntityRepo) GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (*legal_entity_models.LegalEntity, error) {
	var legalEntity legal_entity_models.LegalEntity

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}

	// .First() searches for the record matching the primary key.
	// It returns gorm.ErrRecordNotFound if no entry matches the ID.
	err := db.WithContext(ctx).
		Where("id = ?", legalEntityID).
		First(&legalEntity).Error

	if err != nil {
		return nil, err
	}

	return &legalEntity, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (o *orgLegalEntityRepo) UpdateLegalEntity(
	ctx context.Context,
	legalEntityID uuid.UUID,
	req legal_entity_contracts.UpdateLegalEntityNamesRequestDto,
) error {
	// 1. Extract transaction or use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}

	// 2. Perform optional business validation before hitting DB
	// Even though it's a pointer, if the user provided it, we should validate it.
	if req.FYStartMonth != nil {
		month := *req.FYStartMonth
		if month < 1 || month > 12 {
			return fmt.Errorf("invalid FYStartMonth: %d (must be between 1 and 12)", month)
		}
	}

	// 3. Execute the Update
	// .Model() ensures GORM knows we are targeting the LegalEntity table
	// .Where() ensures we only touch the specific record
	// .Updates() maps the non-nil fields in the DTO to the DB columns
	result := db.WithContext(ctx).
		Model(&legal_entity_models.LegalEntity{}).
		Where("id = ?", legalEntityID).
		Updates(req)

	// 4. Handle Errors
	if result.Error != nil {
		return result.Error
	}

	// 5. Check if the record actually existed
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (o *orgLegalEntityRepo) GetLegalEntitiesByOrgID(ctx context.Context, orgID uuid.UUID) ([]legal_entity_models.LegalEntity, error) {
	var legalEntities []legal_entity_models.LegalEntity

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}

	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&legalEntities).Error

	if err != nil {
		return nil, err
	}

	// Manually handle the "Not Found" case for slices if desired
	if len(legalEntities) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return legalEntities, nil
}

/* ---------------------------------- <***> --------------------------------- */

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

func (o *orgLegalEntityRepo) UpdateLegalEntityStatus(ctx context.Context, legalEntityID uuid.UUID, isActive bool) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	// We use .Model() to target the table and .Update() for the specific column.
	// This is more efficient than fetching the whole record first.
	result := db.WithContext(ctx).
		Model(&legal_entity_models.LegalEntity{}).
		Where("id = ?", legalEntityID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	// Safety check to ensure the ID actually existed
	if result.RowsAffected == 0 {
		return fmt.Errorf("legal entity not found: %s", legalEntityID)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
