package org_supplier_repository

import (
	"context"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"
	dbtx "sharpnr-engine/pkg/dbTx"

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

/* ---------------------------------- <***> --------------------------------- */

func (r *supplierRepository) GetSuppliersLinkOfSuppliers(ctx context.Context, suppliersID uuid.UUID) ([]org_suppliers_models.SupplierLegalEntityLink, error) {
	var links []org_suppliers_models.SupplierLegalEntityLink

	// 1. Extract transaction from context if it exists
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Query for all links associated with the specific Legal Entity
	// We use .Find() to populate the slice of links
	err := db.WithContext(ctx).
		Where("supplier_id = ?", suppliersID).
		Find(&links).
		Error

	if err != nil {
		return nil, err
	}

	return links, nil
}

func (r *supplierRepository) GetSuppliersByLegalEntity(
	ctx context.Context,
	orgID uuid.UUID,
	legalEntityID uuid.UUID,
) ([]org_suppliers_models.Supplier, error) {

	var suppliers []org_suppliers_models.Supplier

	// Extract transaction if exists
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Table("core.org_suppliers AS s").
		Joins(`
			INNER JOIN core.org_supplier_legal_entity_links AS sel
			ON sel.supplier_id = s.id
		`).
		Where(`
			s.organization_id = ?
			AND sel.legal_entity_id = ?
		`, orgID, legalEntityID).
		Find(&suppliers).
		Error

	if err != nil {
		return nil, err
	}

	return suppliers, nil
}
