package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type BarRepository interface {
	CreateBar(ctx context.Context, bar *hospitalitymodels.Bar) error
}
