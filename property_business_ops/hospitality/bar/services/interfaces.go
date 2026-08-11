package barservices

import (
	"context"
	bar_contracts "sharpnr-engine/property_business_ops/hospitality/bar/contracts"

	"github.com/google/uuid"
)

type BarServicesInterface interface {
	SetupDefaultBar(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error)
	GetBarDetailsOfProperty(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) ([]*bar_contracts.BarDetailsReader, error)
	// AddLegalEntity()
	// GetLegalEntity()
}
