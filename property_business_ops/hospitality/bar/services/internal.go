package barservices

import (
	"context"
	bar_contracts "sharpnr-engine/property_business_ops/hospitality/bar/contracts"

	"github.com/google/uuid"
)

func (r *barServices) GetBarDetailsOfProperty(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
) ([]*bar_contracts.BarDetailsReader, error) {

	_, err := r.barRepo.GetBarByPropertyId(ctx, propertyID)
	return nil, err

}
