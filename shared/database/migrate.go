package database

import (
	"gorm.io/gorm"
)

// -------------------------------
// 🧱 Run Migration to Create Tables
// -------------------------------
func MigrateDB(db *gorm.DB) error {

	if err := createSchemas(db); err != nil {
		return err
	}

	migrations := []func(*gorm.DB) error{

		migrateCore,
		migrateOrgProperty,
		migrateOrgLegalEntity,
		migrateEmployee,
		migrateHotel,
		migrateRestaurant,
		migrateBar,
		migrateKitchen,
		migrateHospitalityMenu,
		migrateAccounting,
		migrateSetupWizard,
		migrateInvoicingEngine,
		migrateHospitalityOrderEngine,
		migrateSeroEngine,
	}

	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			return err
		}
	}

	return nil
}
