package database

import (
	payment_models "sharpnr-engine/payment_engine/models"

	"gorm.io/gorm"
)

func migratePaymentsEngine(db *gorm.DB) error {
	return db.AutoMigrate(
		&payment_models.OrgPaymentGetawayCredentials{},
		&payment_models.Payments{},
	)
}
