package migrations

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TransferRestaurantConfiguration(db *gorm.DB) error {

	// If old columns don't exist, migration already done
	if !db.Migrator().HasColumn(&restaurant_models.Restaurant{}, "IsActive") {
		return nil
	}

	var count int64
	if err := db.Model(&restaurant_models.RestaurantConfiguration{}).Count(&count).Error; err != nil {
		return err
	}

	// If configuration records already exist → skip
	if count > 0 {
		return nil
	}

	var restaurants []restaurant_models.Restaurant
	batchSize := 100

	return db.FindInBatches(&restaurants, batchSize, func(tx *gorm.DB, batch int) error {

		var configs []restaurant_models.RestaurantConfiguration

		for _, r := range restaurants {

			status := restaurant_constants.RestaurantStatusCreated
			if r.IsConfigured {
				status = restaurant_constants.RestaurantStatusCompleted
			}

			configs = append(configs, restaurant_models.RestaurantConfiguration{
				ID:             uuid.New(),
				OrganizationID: r.OrganizationID,
				RestaurantID:   r.ID,
				IsActive:       r.IsActive,
				IsConfigured:   r.IsConfigured,
				CurrentStatus:  status,
			})
		}

		return tx.CreateInBatches(&configs, batchSize).Error
	}).Error
}

func RemoveRestaurantLegacyColumns(db *gorm.DB) error {

	// Drop IsActive if it exists
	if db.Migrator().HasColumn(&restaurant_models.Restaurant{}, "IsActive") {
		if err := db.Migrator().DropColumn(&restaurant_models.Restaurant{}, "IsActive"); err != nil {
			return err
		}
	}

	// Drop IsConfigured if it exists
	if db.Migrator().HasColumn(&restaurant_models.Restaurant{}, "IsConfigured") {
		if err := db.Migrator().DropColumn(&restaurant_models.Restaurant{}, "IsConfigured"); err != nil {
			return err
		}
	}

	return nil
}
