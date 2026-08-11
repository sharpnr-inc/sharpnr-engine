package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/property_business_ops/hospitality/menu/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

/* -------------------------------------------------------------------------- */
/*                         Menu Group Register Methods                        */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuGroupKitchenRegister(ctx context.Context, entry *hospitalityMenu_models.MenuGroupKitchenRegister) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// clause.OnConflict allows us to handle the unique constraint on MenuItemID
	err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_group_id"}}, // The unique index column
		DoUpdates: clause.AssignmentColumns([]string{"kitchen_id", "updated_at"}),
	}).Create(entry).Error

	return err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuGroupPreferredKitchen(ctx context.Context, menuGroupID uuid.UUID) (*hospitalityMenu_models.MenuGroupKitchenRegister, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var register hospitalityMenu_models.MenuGroupKitchenRegister
	err := db.WithContext(ctx).Where("menu_group_id = ?", menuGroupID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}

/* -------------------------------------------------------------------------- */
/*                       Menu Category Registers Methods                      */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuCategoryKitchenRegister(ctx context.Context, entry *hospitalityMenu_models.MenuCategoryKitchenRegister) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// clause.OnConflict allows us to handle the unique constraint on MenuItemID
	err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_category_id"}}, // The unique index column
		DoUpdates: clause.AssignmentColumns([]string{"kitchen_id", "updated_at"}),
	}).Create(entry).Error

	return err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetCategoryPreferredKitchen(ctx context.Context, menuCategoryID uuid.UUID) (*hospitalityMenu_models.MenuCategoryKitchenRegister, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var register hospitalityMenu_models.MenuCategoryKitchenRegister
	err := db.WithContext(ctx).Where("menu_category_id = ?", menuCategoryID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}

/* -------------------------------------------------------------------------- */
/*                     Menu Sub Category Register Methods                     */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateMenuSubCategoryKitchenRegister(ctx context.Context, entry *hospitalityMenu_models.MenuSubCategoryKitchenRegister) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// clause.OnConflict allows us to handle the unique constraint on MenuItemID
	err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_sub_category_id"}}, // The unique index column
		DoUpdates: clause.AssignmentColumns([]string{"kitchen_id", "updated_at"}),
	}).Create(entry).Error

	return err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetSubCategoryPreferredKitchen(ctx context.Context, menuSubCategoryID uuid.UUID) (*hospitalityMenu_models.MenuSubCategoryKitchenRegister, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var register hospitalityMenu_models.MenuSubCategoryKitchenRegister
	err := db.WithContext(ctx).Where("menu_sub_category_id = ?", menuSubCategoryID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}

/* -------------------------------------------------------------------------- */
/*                         Menu Item Register Methods                         */
/* -------------------------------------------------------------------------- */

func (r *menuRepo) CreateEntryInKitchenMenuRegister(ctx context.Context, entry *hospitalityMenu_models.MenuItemKitchenRegister) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// clause.OnConflict allows us to handle the unique constraint on MenuItemID
	err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_item_id"}}, // The unique index column
		DoUpdates: clause.AssignmentColumns([]string{"kitchen_id", "updated_at"}),
	}).Create(entry).Error

	return err
}

/* ---------------------------------- <***> --------------------------------- */

func (r *menuRepo) GetMenuItemKitchen(ctx context.Context, menuItemID uuid.UUID) (*hospitalityMenu_models.MenuItemKitchenRegister, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var register hospitalityMenu_models.MenuItemKitchenRegister
	err := db.WithContext(ctx).Where("menu_item_id = ?", menuItemID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}

/* ----------------------------------- XXX ---------------------------------- */
