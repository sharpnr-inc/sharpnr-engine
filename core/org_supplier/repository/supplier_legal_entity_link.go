package org_supplier_repository

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

	"github.com/google/uuid"
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

func (r *supplierRepository) GetSuppliersLinkOfSuppliers(ctx context.Context, legalEntityID uuid.UUID) ([]org_suppliers_models.SupplierLegalEntityLink, error) {
	var links []org_suppliers_models.SupplierLegalEntityLink

	// 1. Extract transaction from context if it exists
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Query for all links associated with the specific Legal Entity
	// We use .Find() to populate the slice of links
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&links).
		Error

	if err != nil {
		return nil, err
	}

	return links, nil
}
