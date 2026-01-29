package kitchenservices

import (
	"context"
	"fmt"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
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

func (r *kitchenService) GetKitchensByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_contracts.Kitchen_Reader, error) {

	kitchen, err := r.kitchenRepo.GetKitchenByPropertyId(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	readers := make([]*kitchen_contracts.Kitchen_Reader, 0, len(kitchen))

	for _, k := range kitchen {
		reader := &kitchen_contracts.Kitchen_Reader{
			ID:     k.ID,
			Name:   k.Name,
			Status: k.IsActive,
		}

		readers = append(readers, reader)
	}

	return readers, nil
}
