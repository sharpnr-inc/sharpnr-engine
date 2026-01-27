package kitchenservices

import (
	"context"

	"github.com/google/uuid"
)

type KitchenServiceInterface interface {
	// Kitchen setup
	CreateDefaultKitchen(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, module string) error
}
