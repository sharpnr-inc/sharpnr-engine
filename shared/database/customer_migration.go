package database

import (
	org_customer_models "sharpnr-engine/core/customer/models"

	"gorm.io/gorm"
)

func mingrateCustomerEngine(db *gorm.DB) error {
	return db.AutoMigrate(
		org_customer_models.Customer{},
	)
}
