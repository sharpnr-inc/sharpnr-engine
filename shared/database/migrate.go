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
		migrateLegalEntity,
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
		migrateSharpIam,
		mingrateCustomerEngine,
		migrateInvoicingEngine,
		migratePaymentsEngine,
		migrateInventoryEngine,
		migrateOrgSuppliers,
	}

	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			return err
		}
	}

	return nil
}
