package database

import (
	bar_models "sharpnr-engine/modules/hospitality/bar/models"

	"gorm.io/gorm"
)

func migrateBar(db *gorm.DB) error {

	if err := db.Exec(
		`CREATE SCHEMA IF NOT EXISTS hospitality`,
	).Error; err != nil {
		return err
	}
	return db.AutoMigrate(

		&bar_models.Bar{},
		&bar_models.BarFloor{},
		&bar_models.BarTable{},
		&bar_models.BarCharge{},
		&bar_models.BarTableSession{},
		&bar_models.BarTicket{},
		&bar_models.BarTicketItem{},
		&bar_models.BarOrderTicketLink{},
		&bar_models.BarPayment{},
	)
}
