package kitchenservices

import (
	"context"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"

	"github.com/google/uuid"
)

type KitchenServiceInterface interface {
	// Kitchen setup
	CreateDefaultKitchen(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, module string) error
	GetKitchensByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_contracts.Kitchen_Reader, error)
	GetKitchensByOrgID(ctx context.Context, orgID uuid.UUID) ([]*kitchen_contracts.Kitchen_Reader, error)
	CreateKOT(ctx context.Context, input kitchen_contracts.CreateKOT_req_dto) error
	GetKots(ctx context.Context, kitchenID uuid.UUID) ([]kitchen_contracts.KOT_resp_dto, error)
	UpdateKotStatus(ctx context.Context, kitchenID uuid.UUID, kotID uuid.UUID, status string) error
}
