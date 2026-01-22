package database

import "gorm.io/gorm"

func createSchemas(db *gorm.DB) error {
	schemas := []string{
		"core",
		"employee",
		"hospitality",
		"accounting_engine",
		"organization_property",
		"setup_wizard",
		"legal_entity",
	}

	for _, s := range schemas {
		if err := db.Exec(`CREATE SCHEMA IF NOT EXISTS ` + s).Error; err != nil {
			return err
		}
	}
	return nil
}
