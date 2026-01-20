package database

import (
	"gorm.io/gorm"
)

// -------------------------------
// 🧱 Run Migration to Create Tables
// -------------------------------
func MigrateDB(db *gorm.DB) error {

	if err := migrateCore(db); err != nil {
		return err
	}

	if err := migrateEmployee(db); err != nil {
		return err
	}

	if err := migrateHotel(db); err != nil {
		return err
	}

	if err := migrateBar(db); err != nil {
		return err
	}

	if err := migrateRestaurant(db); err != nil {
		return err
	}

	if err := migrateKitchen(db); err != nil {
		return err
	}

	if err := migrateHospitalityMenu(db); err != nil {
		return err
	}

	if err := migrateAccounting(db); err != nil {
		return err
	}

	if err := migrateSetupWizard(db); err != nil {
		return err
	}

	if err := migrateOrgProperty(db); err != nil {
		return err
	}

	return nil
}
