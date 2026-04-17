package database

import (
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

	"gorm.io/gorm"
)

func migrateOrgSuppliers(db *gorm.DB) error {
	return db.AutoMigrate(
		&org_suppliers_models.Supplier{},
		&org_suppliers_models.SupplierLegalEntityLink{},
	)
}
