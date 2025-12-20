package hotelservices

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	"time"

	"github.com/google/uuid"
)

func (s *hotelService) PreviewBill(ctx context.Context, orgID, bookingID uuid.UUID) (*hoteldtos.BillPreviewDTO, error) {

	// -------------------------
	// 1️⃣ Resolve hotel
	// -------------------------
	hotel, err := s.hotelRepo.GetHotelByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// -------------------------
	// 2️⃣ Booking
	// -------------------------
	booking, err := s.hotelRepo.GetBookingByBookingId(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	if booking.HotelID != hotel.ID {
		return nil, errors.New("booking does not belong to hotel")
	}

	// -------------------------
	// 3️⃣ Charges
	// -------------------------
	charges, err := s.hotelRepo.GetChargesByBookingID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	baseAmount := 0.0
	for _, c := range charges {
		baseAmount += c.Amount
	}

	// -------------------------
	// 4️⃣ Payments
	// -------------------------
	payments, err := s.hotelRepo.GetPaymentsByBookingId(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	paid := 0.0
	for _, p := range payments {
		paid += p.Amount
	}

	// -------------------------
	// 5️⃣ Tax rules
	// -------------------------
	taxRules, err := s.hotelRepo.GetActiveTaxRulesByHotelAndType(
		ctx,
		hotel.ID,
		hospitalityconstants.TaxTypeRoomGST,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	rule, err := evaluateTaxRule(taxRules, baseAmount)
	if err != nil {
		return nil, err
	}

	// -------------------------
	// 6️⃣ Tax components
	// -------------------------
	comps, err := s.hotelRepo.GetTaxComponentsByRuleID(ctx, rule.ID)
	if err != nil {
		return nil, err
	}

	taxAmount := 0.0
	breakup := []hoteldtos.TaxComponentPreviewDTO{}

	for _, c := range comps {
		amt := baseAmount * c.Rate / 100
		taxAmount += amt

		breakup = append(breakup, hoteldtos.TaxComponentPreviewDTO{
			ComponentType: string(c.ComponentType),
			Rate:          c.Rate,
			Amount:        amt,
		})
	}

	total := baseAmount + taxAmount

	// -------------------------
	// 7️⃣ Final preview
	// -------------------------
	return &hoteldtos.BillPreviewDTO{
		BookingID:   bookingID,
		BaseAmount:  baseAmount,
		TaxAmount:   taxAmount,
		TotalAmount: total,
		TaxBreakup:  breakup,
		PaidAmount:  paid,
		Balance:     total - paid,
	}, nil

}

func evaluateTaxRule(
	rules []hospitalitymodels.HotelTaxRule,
	baseAmount float64,
) (*hospitalitymodels.HotelTaxRule, error) {

	for _, rule := range rules {

		// PRICE_SLAB
		if rule.ConditionType == hospitalityconstants.TaxConditionPriceSlab {

			if rule.MinValue != nil && baseAmount < *rule.MinValue {
				continue
			}
			if rule.MaxValue != nil && baseAmount > *rule.MaxValue {
				continue
			}

			return &rule, nil
		}
	}

	return nil, errors.New("no applicable tax rule found")
}
