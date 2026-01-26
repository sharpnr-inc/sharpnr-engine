package barservices

import (
	"context"
	bar_contracts "sharpnr-engine/modules/hospitality/bar/contracts"

	"github.com/google/uuid"
)

func (r *barServices) GetBarDetails(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
) (*bar_contracts.BarDetailsReader, error) {

	return nil, nil
}
