package hotelservices

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/hotel_events"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"time"

	"github.com/google/uuid"
)

// CreateBookingForRoom creates a room booking along with optional guest entries.
// Steps:
// 1. Start a DB transaction
// 2. Validate room, hotel, and primary guest
// 3. Create room booking record
// 4. Create additional booking guests (if any)
// 5. Commit transaction (auto-rollback on error)
func (r *hotelService) CreateBookingForRoom(
	ctx context.Context,
	orgId uuid.UUID,
	req *hoteldtos.RoomBookingRequest_dto,
) error {

	var booking *hospitalitymodels.HotelRoomBooking

	err := r.hotelRepo.WithTx(ctx, func(txRepo hotelrepo.HotelRepositoryInterface) error {

		// Step 1: Validate room
		room, err := txRepo.GetRoomByID(ctx, req.RoomId)
		if err != nil {
			return errors.New("cannot find the room")
		}

		// Step 2: Validate hotel
		hotel, err := txRepo.GetHotelByOrgID(ctx, orgId)
		if err != nil {
			return err
		}

		// Step 3: Validate primary guest
		guest, err := txRepo.GetHotelGuestById(ctx, req.PrimaryGuestID)
		if err != nil {
			return err
		}

		// Step 4: Create booking
		booking = &hospitalitymodels.HotelRoomBooking{
			ID:             uuid.New(),
			HotelID:        hotel.ID,
			RoomID:         room.ID,
			Status:         "RESERVED",
			PrimaryGuestID: guest.ID,
			CheckInDate:    req.CheckinDate,
			CheckOutDate:   req.CheckoutDate,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := txRepo.CreateRoomBooking(ctx, booking); err != nil {
			return err
		}

		// Step 5: Create booking guests (optional)
		for _, bg := range req.BookingGuests {
			bookingGuest := &hospitalitymodels.BookingGuest{
				ID:                 uuid.New(),
				HotelID:            hotel.ID,
				HotelGuestID:       guest.ID,
				HotelRoomBookingID: booking.ID,
				Name:               bg.Name,
				Mobile:             bg.Mobile,
				Email:              *bg.Email,
			}

			if err := txRepo.CreateBookingGuest(ctx, bookingGuest); err != nil {
				return err
			}
		}

		if err := txRepo.UpdateRoomStatus(ctx, hotel.ID, room.ID, string(hospitalityconstants.Booked)); err != nil {
			return err
		}

		if err := txRepo.AddCurrentBookingIdToRoom(ctx, room.ID, booking.ID); err != nil {
			return err
		}

		roomPrice, err := txRepo.GetRoomPriceByRoomId(ctx, req.RoomId)
		if err != nil {
			return nil
		}

		charges := &hospitalitymodels.HotelBookingCharge{
			ID:                 uuid.New(),
			HotelRoomBookingID: booking.ID,
			ChargeType:         "Room Tariff",
			ReferenceID:        nil,
			Description:        "Room tariff for One night",
			Amount:             float64(roomPrice.Price),
			CreatedAt:          time.Now(),
			PaymentStatus:      "PENDING",
		}

		if err := txRepo.AddChargesInBooking(ctx, charges); err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return err
	}

	// ✅ Step 6: Publish event AFTER successful commit
	r.sere.Publish(ctx, "BookingCreated", hotelevents.RoomBookingCreated{
		HotelRoomBookingId: booking.ID,
		HotelId:            booking.HotelID,
		RoomId:             booking.RoomID,
		PrimaryGuestId:     booking.PrimaryGuestID,
	})

	return nil
}

func (r *hotelService) GetRoomBookingDetails(ctx context.Context, roomId, orgID uuid.UUID) (*hoteldtos.BookingDetailsDTO, error) {

	hotel, err := r.hotelRepo.GetHotelByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	room, err := r.hotelRepo.GetRoomByID(ctx, roomId)
	if err != nil {
		return nil, errors.New("room not found")
	}

	booking, err := r.hotelRepo.GetBookingByBookingId(ctx, *room.CurrentBookingID)
	if err != nil {
		return nil, err
	}

	if booking.HotelID != hotel.ID {
		return nil, errors.New("booking does not belong to hotel")
	}

	guest, err := r.hotelRepo.GetHotelGuestById(ctx, booking.PrimaryGuestID)
	if err != nil {
		return nil, err
	}

	// 5️⃣ Charges
	charges, err := r.hotelRepo.GetChargesByBookingID(ctx, *room.CurrentBookingID)
	if err != nil {
		return nil, err
	}

	// 6️⃣ Payments
	payments, err := r.hotelRepo.GetPaymentsByBookingId(ctx, *room.CurrentBookingID)
	if err != nil {
		return nil, err
	}

	bill, _ := r.hotelRepo.GetBillByBookingId(ctx, *room.CurrentBookingID)

	// 8️⃣ Calculate balance
	totalCharges := 0.0
	for _, c := range charges {
		totalCharges += c.Amount
	}

	totalPaid := 0.0
	for _, p := range payments {
		totalPaid += p.Amount
	}

	balance := totalCharges - totalPaid
	if bill != nil {
		balance = bill.TotalAmount - totalPaid
	}

	resp := &hoteldtos.BookingDetailsDTO{
		Booking: hoteldtos.BookingInfoDTO{
			ID:           booking.ID,
			CheckInDate:  booking.CheckInDate,
			CheckOutDate: booking.CheckOutDate,
			Status:       booking.Status,
			CreatedAt:    booking.CreatedAt,
		},
		Room: hoteldtos.RoomInfoDTO{
			ID:       room.ID,
			RoomName: room.RoomName,
			RoomType: room.BedSize,
		},
		Guest: hoteldtos.GuestInfoDTO{
			ID:     guest.ID,
			Name:   guest.Name,
			Mobile: guest.Mobile,
		},
		Charges:  mapCharges(charges),
		Payments: mapPayments(payments),
		Bill:     mapBill(bill),
		Balance:  balance,
	}

	return resp, nil
}

func mapCharges(charges []hospitalitymodels.HotelBookingCharge) []hoteldtos.ChargeDTO {
	out := make([]hoteldtos.ChargeDTO, 0, len(charges))
	for _, c := range charges {
		out = append(out, hoteldtos.ChargeDTO{
			ID:          c.ID,
			Type:        c.ChargeType,
			Description: c.Description,
			Amount:      c.Amount,
			CreatedAt:   c.CreatedAt,
		})
	}
	return out
}

func mapPayments(payments []hospitalitymodels.HotelRoomBookingPayment) []hoteldtos.PaymentDTO {
	out := make([]hoteldtos.PaymentDTO, 0, len(payments))
	for _, p := range payments {
		out = append(out, hoteldtos.PaymentDTO{
			ID:        p.ID,
			Amount:    p.Amount,
			Method:    p.Mode,
			Status:    p.Status,
			CreatedAt: p.CreatedAt,
		})
	}
	return out
}

func mapBill(bill *hospitalitymodels.HotelBill) *hoteldtos.BillDTO {
	if bill == nil {
		return nil
	}
	return &hoteldtos.BillDTO{
		ID:          bill.ID,
		BaseAmount:  bill.BaseAmount,
		TaxAmount:   bill.TaxAmount,
		TotalAmount: bill.TotalAmount,
		Status:      bill.Status,
	}
}
