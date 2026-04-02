package org_supplier_repository

import (
	"context"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

	"github.com/google/uuid"
)

type SupplierRepositoryInterface interface {

	// This Repo methods Create the Supplier in DB
	CreateSupplier(ctx context.Context, supplier *org_suppliers_models.Supplier) error
	GetSuppliersByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_suppliers_models.Supplier, error)

	CreateSupplierLegalEntityLink(ctx context.Context, supplierLegalEntityLink *org_suppliers_models.SupplierLegalEntityLink) error
	GetSuppliersLinkOfSuppliers(ctx context.Context, suppliersID uuid.UUID) ([]org_suppliers_models.SupplierLegalEntityLink, error)
}
