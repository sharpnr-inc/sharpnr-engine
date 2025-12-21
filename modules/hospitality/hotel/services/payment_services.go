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

func (c *hotelService) AddPaymentForBooking(ctx context.Context, orgID uuid.UUID, req hoteldtos.AddPaymentsForBookingRequest_dto) error {

	hotel, err := c.hotelRepo.GetHotelByOrgID(ctx, orgID)
	if err != nil {
		return err
	}

	booking, err := c.hotelRepo.GetBookingByBookingId(ctx, req.BookingID)
	if err != nil {
		return err
	}

	if booking.HotelID != hotel.ID {
		return errors.New("booking does not belong to hotel")
	}

	payment := &hospitalitymodels.HotelRoomBookingPayment{
		ID:                 uuid.New(),
		HotelRoomBookingID: booking.ID,
		Amount:             req.Amount,
		Mode:               req.Method,
		Status:             hospitalityconstants.PaymentSuccess,
		Reference:          *req.Reference,
		CreatedAt:          time.Now(),
	}

	return c.hotelRepo.AddPaymentsInBooking(ctx, payment)

}
