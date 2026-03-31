package org_supplier_repository

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"
)

func (r *supplierRepository) CreateSupplierLegalEntityLink(ctx context.Context, link *org_suppliers_models.SupplierLegalEntityLink) error {
	// 1. Extract transaction from context if it exists
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Use GORM to persist the link record
	// .WithContext(ctx) ensures the trace and timeout are respected
	return db.WithContext(ctx).Create(link).Error
}
