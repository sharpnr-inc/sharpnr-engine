package hotelservices

import (
	"context"
	"errors"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
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
	rules []hotel_models.HotelTaxRule,
	baseAmount float64,
) (*hotel_models.HotelTaxRule, error) {

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

func (s *hotelService) FinalizeBill(
	ctx context.Context,
	orgID uuid.UUID,
	bookingID uuid.UUID,
) (*hotel_models.HotelBill, error) {

	var finalizedBill *hotel_models.HotelBill

	err := s.hotelRepo.WithTx(ctx, func(txRepo hotelrepo.HotelRepositoryInterface) error {
		// 1️⃣ Resolve hotel
		hotel, err := txRepo.GetHotelByOrgID(ctx, orgID)
		if err != nil {
			return err
		}

		// 2️⃣ Booking
		booking, err := txRepo.GetBookingByBookingId(ctx, bookingID)
		if err != nil {
			return err
		}

		if booking.HotelID != hotel.ID {
			return errors.New("booking does not belong to hotel")
		}

		// 3️⃣ Prevent double finalize
		existingBill, _ := txRepo.GetBillByBookingId(ctx, bookingID)
		if existingBill != nil {
			return errors.New("bill already finalized")
		}

		// 4️⃣ Charges must exist
		charges, err := txRepo.GetChargesByBookingID(ctx, bookingID)
		if err != nil {
			return err
		}

		if len(charges) == 0 {
			return errors.New("cannot finalize bill without charges")
		}

		// 5️⃣ Base amount
		baseAmount := 0.0
		for _, c := range charges {
			baseAmount += c.Amount
		}

		// 6️⃣ Tax evaluation (your existing logic)
		taxResult, err := s.CalculateTaxForBooking(
			ctx,
			hotel.ID,
			baseAmount,
		)
		if err != nil {
			return err
		}

		now := time.Now()
		// 7️⃣ Create bill
		bill := &hotel_models.HotelBill{
			ID:                 uuid.New(),
			HotelRoomBookingID: bookingID,
			HotelID:            hotel.ID,
			BaseAmount:         baseAmount,
			TaxAmount:          taxResult.TotalTax,
			TotalAmount:        baseAmount + taxResult.TotalTax,
			Status:             string(hospitalityconstants.BillFinalized),
			CreatedAt:          time.Now(),
			FinalizedAt:        &now,
		}

		if err := txRepo.CreateBookingBill(ctx, bill); err != nil {
			return err
		}

		// 8️⃣ Freeze tax snapshot
		for _, t := range taxResult.Components {
			bt := &hotel_models.HotelBillTax{
				ID:            uuid.New(),
				HotelBillID:   bill.ID,
				ComponentType: string(t.ComponentType),
				Rate:          t.Rate,
				Amount:        t.Amount,
			}
			if err := txRepo.CreateHotelBillTax(ctx, bt); err != nil {
				return err
			}
		}

		// 9️⃣ Update booking status → CHECKED_OUT
		if err := txRepo.UpdateBookingStatus(
			ctx,
			bookingID,
			hospitalityconstants.HotelRoomCheckedOut,
		); err != nil {
			return err
		}

		if err := txRepo.UpdateRoomStatus(ctx, hotel.ID, booking.RoomID, string(hospitalityconstants.Available)); err != nil {
			return err
		}
		finalizedBill = bill
		return nil

	})
	if err != nil {
		return nil, err
	}

	return finalizedBill, nil
}

func (s *hotelService) FinalizeBillAndGetInvoice(
	ctx context.Context,
	orgID uuid.UUID,
	bookingID uuid.UUID,
) (*hoteldtos.HotelInvoiceDTO, error) {

	// 1️⃣ Finalize bill (DB write)
	bill, err := s.FinalizeBill(ctx, orgID, bookingID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Fetch booking (needed for invoice)
	booking, err := s.hotelRepo.GetBookingByBookingId(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Build invoice DTO (pure read)
	return s.buildInvoiceDTO(ctx, booking, bill)
}
