package kitchenservices

import (
	"context"
	kitchen_contracts "sharpnr-engine/property_infra/hospitality_infra/kitchen/contracts"

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

	GetKitchenDetails(ctx context.Context, kitchenID uuid.UUID) (*kitchen_contracts.Kitchen_Reader, error)
	KitchenConfigServiceInterface
}

type KitchenConfigServiceInterface interface {
	UpdateGeneralDetails(ctx context.Context, kitchenID uuid.UUID, req_dto kitchen_contracts.UpdateKitchenGeneralDetails_req_dto) error
	GetKitchenConfigs(ctx context.Context, kitchenID uuid.UUID) (*kitchen_contracts.KitchenConfigResponseDto, error)
	ActivateDeactivate(ctx context.Context, orgID, emplID, kitchenID uuid.UUID, kitchenSwitch bool) error
}
