package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (r *menuRepo) CreateEntryInKitchenMenuRegister(ctx context.Context, entry *hospitalityMenu_models.MenuItemKitchenRegister) error {
	// clause.OnConflict allows us to handle the unique constraint on MenuItemID
	err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_item_id"}}, // The unique index column
		DoUpdates: clause.AssignmentColumns([]string{"kitchen_id", "owner_type", "updated_at"}),
	}).Create(entry).Error

	return err
}

func (r *menuRepo) GetKitchenMenuRegisterByMenuItemID(ctx context.Context, menuItemID uuid.UUID, kitchenID uuid.UUID) (*hospitalityMenu_models.MenuItemKitchenRegister, error) {
	var register hospitalityMenu_models.MenuItemKitchenRegister
	err := r.db.WithContext(ctx).Where("menu_item_id = ? AND kitchen_id = ?", menuItemID, kitchenID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}

func (r *menuRepo) GetMenuItemKitchen(ctx context.Context, menuItemID uuid.UUID) (*hospitalityMenu_models.MenuItemKitchenRegister, error) {
	var register hospitalityMenu_models.MenuItemKitchenRegister
	err := r.db.WithContext(ctx).Where("menu_item_id = ?", menuItemID).First(&register).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}
