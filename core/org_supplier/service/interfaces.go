package suppliers_service

import (
	"context"
	supplier_contracts "sharpnr-engine/core/org_supplier/contracts"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

	"github.com/google/uuid"
)

type SupplierServiceInterface interface {
	CreateNewSupplier(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, request *suppliers_dtos.CreateSuppliersAPIRequestdto) error
	GetAllSupplierByOrgID(ctx context.Context, orgID, emplID uuid.UUID) ([]org_suppliers_models.Supplier, error)
	ListAllSuppliersOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]org_suppliers_models.Supplier, error)
	ListAllLegalEntityLinksOfSuppliers(ctx context.Context, orgID, emplID, suppliersID uuid.UUID) ([]org_suppliers_models.SupplierLegalEntityLink, error)
	AddLegalEntityLinksToSupplier(ctx context.Context, orgID, emplID uuid.UUID, req supplier_contracts.AddLegalEntityLinkInSuppliersRequestDto) error
}
