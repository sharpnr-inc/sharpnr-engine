package migrations

import (
	"gorm.io/gorm"
)

func RunDataMigrations(db *gorm.DB) error {

	if err := TransferRestaurantConfiguration(db); err != nil {
		return err
	}

	if err := RemoveRestaurantLegacyColumns(db); err != nil {
		return err
	}

	return nil
}
