package hotelrepo

import (
	"context"
	"errors"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
)

func (h *hotelRepository) CreateTaxRuleOfHotel(ctx context.Context, taxRule *hotel_models.HotelTaxRule) error {

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
	taxComp *hotel_models.HotelTaxComponent,
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

func (h *hotelRepository) GetActiveTaxRulesByHotelAndType(
	ctx context.Context,
	hotelId uuid.UUID,
	taxType hospitalityconstants.TaxType,
	at time.Time,
) ([]hotel_models.HotelTaxRule, error) {

	var rules []hotel_models.HotelTaxRule

	err := h.db.
		WithContext(ctx).
		Where("hotel_id = ?", hotelId).
		Where("tax_type = ?", taxType).
		Where("is_active = ?", true).
		Where("effective_from <= ?", at).
		Where(
			h.db.
				Where("effective_to IS NULL").
				Or("effective_to >= ?", at),
		).
		Order("effective_from DESC").
		Find(&rules).
		Error

	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (h *hotelRepository) GetTaxComponentsByRuleID(
	ctx context.Context,
	taxRuleId uuid.UUID,
) ([]hotel_models.HotelTaxComponent, error) {

	var components []hotel_models.HotelTaxComponent

	err := h.db.
		WithContext(ctx).
		Where("hotel_tax_rule_id = ?", taxRuleId).
		Find(&components).
		Error

	if err != nil {
		return nil, err
	}

	return components, nil
}

func (h *hotelRepository) CreateHotelBillTax(ctx context.Context, hotelBillTax *hotel_models.HotelBillTax) error {
	if hotelBillTax == nil {
		return errors.New("hotel bill tax cannot be nil")
	}

	return h.db.WithContext(ctx).Create(hotelBillTax).Error
}

func (h *hotelRepository) GetBillTaxes(
	ctx context.Context,
	billID uuid.UUID,
) ([]hotel_models.HotelBillTax, error) {

	var taxes []hotel_models.HotelBillTax

	err := h.db.WithContext(ctx).
		Where("hotel_bill_id = ?", billID).
		Find(&taxes).Error

	if err != nil {
		return nil, err
	}

	return taxes, nil
}
