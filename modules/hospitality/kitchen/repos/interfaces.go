package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

type KitchenRepositoryInterface interface {
	CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error
	GetKitchenById(ctx context.Context, kitchenId uuid.UUID) ([]*kitchen_models.Kitchen, error)
	GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) ([]*kitchen_models.Kitchen, error)
	GetKitchenByPropertyId(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_models.Kitchen, error)

	//
	SwitchKitchenOnOff(ctx context.Context, KitchenID uuid.UUID, b bool) error
}
