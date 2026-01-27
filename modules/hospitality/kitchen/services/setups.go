package kitchenservices

import (
	"context"
	"fmt"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	"time"

	"github.com/google/uuid"
)

func (r *kitchenService) CreateDefaultKitchen(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, module string) error {

	kitchen := &kitchen_models.Kitchen{
		ID:                     uuid.New(),
		OrganizationID:         orgID,
		OrganizationPropertyID: propertyID,
		Name:                   fmt.Sprintf("%s Kitchen", module),
		Description:            fmt.Sprintf("Default Kitchen for the %s", module),
		Type:                   module,
		IsActive:               false,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
	return r.kitchenRepo.CreateKitchen(ctx, kitchen)
}
