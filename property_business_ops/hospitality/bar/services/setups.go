package barservices

import (
	"context"
	bar_models "sharpnr-engine/business_ops/hospitality/bar/models"
	"time"

	"github.com/google/uuid"
)

func (c *barServices) SetupDefaultBar(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error) {
	defBar := &bar_models.Bar{
		ID:             uuid.New(),
		Name:           "Default Bar",
		OrganizationID: orgID,
		PropertyID:     propertyID,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := c.barRepo.CreateBar(ctx, defBar); err != nil {
		return uuid.Nil, err
	}
	return defBar.ID, nil
}
