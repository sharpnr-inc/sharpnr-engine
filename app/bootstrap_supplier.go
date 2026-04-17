package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_supplier_repository "sharpnr-engine/core/org_supplier/repository"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	"sharpnr-engine/sero"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapSupplier(
	db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
) suppliers_service.SupplierServiceInterface {
	suppliersRepo := org_supplier_repository.NewSupplierRepository(db)
	return suppliers_service.NewSupplierService(
		suppliersRepo,
		txManager,
		mailerEngine,
		seroEngine,
	)
}
