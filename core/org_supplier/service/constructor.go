package suppliers_service

import (
	org_supplier_repository "sharpnr-engine/core/org_supplier/repository"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type supplierService struct {
	supplierRepo org_supplier_repository.SupplierRepositoryInterface
	txManager    dbtx.TxManager
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	seroEngine   *sero.SEROEngine
}

func NewSupplierService(
	supplierRepo org_supplier_repository.SupplierRepositoryInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
) *supplierService {
	return &supplierService{
		supplierRepo: supplierRepo,
		txManager:    txManager,
		mailerEngine: mailerEngine,
		seroEngine:   seroEngine,
	}
}
