package barservices

import (
	"context"
	bar_contracts "sharpnr-engine/modules/hospitality/bar/contracts"

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
