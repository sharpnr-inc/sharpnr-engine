package database

import (
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"gorm.io/gorm"
)

func migrateHospitalityMenu(db *gorm.DB) error {

	return db.AutoMigrate(

		&hospitalityMenu_models.MenuGroup{},
		&hospitalityMenu_models.MenuCategory{},
		&hospitalityMenu_models.MenuSubCategory{},
		&hospitalityMenu_models.MenuItem{},
		&hospitalityMenu_models.MenuVariant{},
		&hospitalityMenu_models.MenuVariantGroup{},
		&hospitalityMenu_models.MenuItemAddon{},

		/* -------------------------------------------------------------------------- */
		/*                           Menu Kitchen Resgisters                          */
		/* -------------------------------------------------------------------------- */
		&hospitalityMenu_models.MenuGroupKitchenRegister{},
		&hospitalityMenu_models.MenuCategoryKitchenRegister{},
		&hospitalityMenu_models.MenuSubCategoryKitchenRegister{},
		&hospitalityMenu_models.MenuItemKitchenRegister{},

		/* -------------------------------------------------------------------------- */
		/*                         Menu Tax Category Register                         */
		/* -------------------------------------------------------------------------- */

		&hospitalityMenu_models.MenuGroupTaxCategory{},
		&hospitalityMenu_models.MenuCategoryTaxCategory{},
		&hospitalityMenu_models.MenuSubCategoryTaxCategory{},
		&hospitalityMenu_models.MenuItemTaxCategory{},
	)
}
