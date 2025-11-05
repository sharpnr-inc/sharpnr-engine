package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type KitchenRepository interface {
	CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error
}
