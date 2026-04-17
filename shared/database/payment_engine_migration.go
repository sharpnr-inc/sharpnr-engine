package database

import (
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"gorm.io/gorm"
)

func migratePaymentsEngine(db *gorm.DB) error {
	return db.AutoMigrate(
		// Tenant registry
		&payment_engine_models.PaymentAccount{},

		// Gateway
		&payment_engine_models.PaymentGatewayConfig{},

		// Payment lifecycle
		&payment_engine_models.PaymentOrder{},
		&payment_engine_models.Payment{},
		&payment_engine_models.PaymentAllocation{},
		&payment_engine_models.Refund{},

		// Webhook
		&payment_engine_models.WebhookEndpoint{},
		&payment_engine_models.WebhookDeliveryLog{},
	)
}
