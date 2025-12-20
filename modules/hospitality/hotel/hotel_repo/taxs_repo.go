package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (h *hotelRepository) CreateTaxRuleOfHotel(ctx context.Context, taxRule *hospitalitymodels.HotelTaxRule) error {

	// Ensure ID exists
	if taxRule.ID == uuid.Nil {
		taxRule.ID = uuid.New()
	}

	// Default active flag (optional safety)
	if !taxRule.IsActive {
		taxRule.IsActive = true
	}

	// Create with associations (HotelTaxComponent)
	err := h.db.
		WithContext(ctx).
		Create(taxRule).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (h *hotelRepository) AddTaxComponentInTaxRule(
	ctx context.Context,
	taxComp *hospitalitymodels.HotelTaxComponent,
) error {

	// Validate required FK
	if taxComp.HotelTaxRuleID == uuid.Nil {
		return errors.New("hotel_tax_rule_id is required")
	}

	// Ensure ID exists
	if taxComp.ID == uuid.Nil {
		taxComp.ID = uuid.New()
	}

	// Insert tax component
	if err := h.db.
		WithContext(ctx).
		Create(taxComp).
		Error; err != nil {
		return err
	}

	return nil
}
