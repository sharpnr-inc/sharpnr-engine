package menurepo

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                         Menu Group Tax Category OP                         */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuGroupTaxCategory(
	ctx context.Context,
	menuGroupTaxCategory *hospitalityMenu_models.MenuGroupTaxCategory,
) error {
	// 1. Extract transaction from context (if one exists)
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		// Fallback to the default repository DB connection
		db = r.db
	}

	// 2. Execute the create operation with the provided context
	// This handles timeouts and cancellations automatically
	if err := db.WithContext(ctx).Create(menuGroupTaxCategory).Error; err != nil {
		// Return the raw GORM error to be handled by the service layer
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuGroupTaxCategoryByMenuGroupID(
	ctx context.Context,
	menuGroupID uuid.UUID,
) (*hospitalityMenu_models.MenuGroupTaxCategory, error) {
	// 1. Extract transaction or use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuGroupTaxCategory

	// 2. Perform the query with context and filter by MenuGroupID
	err := db.WithContext(ctx).
		Where("menu_group_id = ?", menuGroupID).
		First(&result).
		Error

	// 3. Handle Record Not Found specifically
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Returning nil, nil is common if "not found" isn't an "error"
			// but a valid state. Otherwise, return a custom "not found" error.
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuGroupTaxCategoryByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) (*hospitalityMenu_models.MenuGroupTaxCategory, error) {
	// 1. Extract transaction from context or use default connection
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuGroupTaxCategory

	// 2. Query filtered by LegalEntityID
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		First(&result).
		Error

	// 3. Specific error handling for Record Not Found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return nil, nil to indicate no record exists for this Entity
			return nil, nil
		}
		// Return the actual database error (connection issues, etc.)
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) UpdateMenuGroupTaxCategory(
	ctx context.Context,
	id uuid.UUID,
	updates map[string]interface{}, // Pass a map of fields to update
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Add the current timestamp to the updates map
	updates["updated_at"] = time.Now()

	result := db.WithContext(ctx).
		Model(&hospitalityMenu_models.MenuGroupTaxCategory{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) DeleteMenuGroupTaxCategory(
	ctx context.Context,
	id uuid.UUID,
) error {
	// 1. Extract transaction or use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Perform the delete operation
	// We pass the model type so GORM knows which table to target
	result := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&hospitalityMenu_models.MenuGroupTaxCategory{})

	// 3. Handle database errors
	if result.Error != nil {
		return result.Error
	}

	// 4. Handle "Not Found" case
	// If RowsAffected is 0, the ID didn't exist in the database
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* -------------------------------------------------------------------------- */
/*                        Menu Category Tax-Category OP                       */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuCategoryTaxCategory(
	ctx context.Context,
	menuCategoryTaxCategoty *hospitalityMenu_models.MenuCategoryTaxCategory,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Always use the context-aware DB instance
	if err := db.WithContext(ctx).Create(menuCategoryTaxCategoty).Error; err != nil {
		// Consider logging or wrapping the error here
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) UpdateMenuCategoryTaxCategory(
	ctx context.Context,
	menuCategoryID uuid.UUID,
	updates map[string]interface{},
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Ensure we update the timestamp
	updates["updated_at"] = time.Now()

	result := db.WithContext(ctx).
		Model(&hospitalityMenu_models.MenuCategoryTaxCategory{}).
		Where("menu_category_id = ?", menuCategoryID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuCategoryTaxCategoryByMenuCategoryID(
	ctx context.Context,
	menuCategoryID uuid.UUID,
) (*hospitalityMenu_models.MenuCategoryTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuCategoryTaxCategory
	err := db.WithContext(ctx).
		Where("menu_category_id = ?", menuCategoryID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuCategoryTaxCategoryByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) (*hospitalityMenu_models.MenuCategoryTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuCategoryTaxCategory
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) DeleteMenuCategoryTaxCategory(
	ctx context.Context,
	id uuid.UUID,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&hospitalityMenu_models.MenuCategoryTaxCategory{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                      Menu Sub-Category Tax-Category OP                     */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuSubCategoryTaxCategory(
	ctx context.Context,
	data *hospitalityMenu_models.MenuSubCategoryTaxCategory,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(data).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) UpdateMenuSubCategoryTaxCategory(
	ctx context.Context,
	menuSubCategoryID uuid.UUID,
	updates map[string]interface{},
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	updates["updated_at"] = time.Now()

	result := db.WithContext(ctx).
		Model(&hospitalityMenu_models.MenuSubCategoryTaxCategory{}).
		Where("menu_sub_category_id = ?", menuSubCategoryID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuSubCategoryTaxCategoryBySubCategoryID(
	ctx context.Context,
	subCategoryID uuid.UUID,
) (*hospitalityMenu_models.MenuSubCategoryTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuSubCategoryTaxCategory
	err := db.WithContext(ctx).
		Where("menu_sub_category_id = ?", subCategoryID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuSubCategoryTaxCategoryByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) (*hospitalityMenu_models.MenuSubCategoryTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuSubCategoryTaxCategory
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) DeleteMenuSubCategoryTaxCategory(
	ctx context.Context,
	id uuid.UUID,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&hospitalityMenu_models.MenuSubCategoryTaxCategory{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                          MenuItem Tax Category OP                          */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuItemTaxCategory(
	ctx context.Context,
	data *hospitalityMenu_models.MenuItemTaxCategory,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(data).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) UpdateMenuItemTaxCategory(
	ctx context.Context,
	menuItemID uuid.UUID,
	updates map[string]interface{},
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	updates["updated_at"] = time.Now()

	result := db.WithContext(ctx).
		Model(&hospitalityMenu_models.MenuItemTaxCategory{}).
		Where("menu_item_id = ?", menuItemID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuItemTaxCategoryByMenuItemID(
	ctx context.Context,
	menuItemID uuid.UUID,
) (*hospitalityMenu_models.MenuItemTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuItemTaxCategory
	err := db.WithContext(ctx).
		Where("menu_item_id = ?", menuItemID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuItemTaxCategoryByLegalEntityID(
	ctx context.Context,
	legalEntityID uuid.UUID,
) (*hospitalityMenu_models.MenuItemTaxCategory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result hospitalityMenu_models.MenuItemTaxCategory
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		First(&result).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) DeleteMenuItemTaxCategory(
	ctx context.Context,
	id uuid.UUID,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&hospitalityMenu_models.MenuItemTaxCategory{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
