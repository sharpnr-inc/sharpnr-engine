package org_supplier_repository

import (
	"context"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

func (r *supplierRepository) CreateSupplier(ctx context.Context, supplier *org_suppliers_models.Supplier) error {
	// Extract transaction from context if it exists; otherwise, use the default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Use GORM's Create method to persist the entity
	return db.WithContext(ctx).Create(supplier).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *supplierRepository) GetSuppliersByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_suppliers_models.Supplier, error) {
	var suppliers []org_suppliers_models.Supplier

	// Extract transaction from context if it exists; otherwise, use the default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Querying with a Where clause and Find to populate the slice
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&suppliers).
		Error

	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

/* ---------------------------------- <***> --------------------------------- */
