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

	// if err := migrateCore(db); err != nil {
	// 	return err
	// }

	// if err := migrateEmployee(db); err != nil {
	// 	return err
	// }

	// if err := migrateHotel(db); err != nil {
	// 	return err
	// }

	// if err := migrateBar(db); err != nil {
	// 	return err
	// }

	// if err := migrateRestaurant(db); err != nil {
	// 	return err
	// }

	// if err := migrateKitchen(db); err != nil {
	// 	return err
	// }

	// if err := migrateHospitalityMenu(db); err != nil {
	// 	return err
	// }

	// if err := migrateAccounting(db); err != nil {
	// 	return err
	// }

	// if err := migrateSetupWizard(db); err != nil {
	// 	return err
	// }

	// if err := migrateOrgProperty(db); err != nil {
	// 	return err
	// }

	// return nil

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
	}

	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			return err
		}
	}

	return nil
}
