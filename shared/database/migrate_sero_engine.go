package database

import (
	"sharpnr-engine/sero"

	"gorm.io/gorm"
)

func migrateSeroEngine(db *gorm.DB) error {

	return db.AutoMigrate(
		&sero.SeroOutbox{},
	)
}
