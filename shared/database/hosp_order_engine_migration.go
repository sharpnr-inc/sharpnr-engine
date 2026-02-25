package database

import (
	hospitality_order_engine_models "sharpnr-engine/modules/hospitality/order_engine/models"

	"gorm.io/gorm"
)

func migrateHospitalityOrderEngine(db *gorm.DB) error {
	return db.AutoMigrate(
		&hospitality_order_engine_models.HospitalityOrder{},
		&hospitality_order_engine_models.HospitalityOrderItem{},
		&hospitality_order_engine_models.HospitalityOrderItemAddon{},
		&hospitality_order_engine_models.HospitalityOrderItemVariant{},
		&hospitality_order_engine_models.OrderKitchenRouting{},
	)
}
