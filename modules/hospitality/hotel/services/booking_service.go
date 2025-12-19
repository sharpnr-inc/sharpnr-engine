package hotelservices

import (
	"context"
	"errors"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
)

func (r *hotelService) CreateBookingForRoom(ctx context.Context, orgId uuid.UUID, req *hospitalitydtos.Hotel_Room_booking_Req) error {

	room, err := r.hotelRepo.GetRoomByID(ctx, req.RoomId)
	if err != nil {
		return errors.New("cannot find the room")
	}

	// Here it Will Be Guest i mean get Guest By Id

	guest := &hospitalitymodels.HotelGuest{
		ID:        uuid.New(),
		Name:      req.GuestName,
		Mobile:    req.GuestMobile,
		Email:     req.GuestEmailAddr,
		Gender:    "MALE",
		AadhaarId: "9828823314",
	}

	if err := r.hotelRepo.CreatePrimaryGuest(ctx, guest); err != nil {
		return err
	}

	booking := &hospitalitymodels.HotelRoomBooking{
		ID:             uuid.New(),
		HotelID:        req.Hotel_Id,
		RoomID:         req.RoomId,
		Status:         "RESERVED",
		Room:           room,
		PrimaryGuestID: guest.ID,
		PrimaryGuest:   guest,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := r.hotelRepo.CreateRoomBooking(ctx, booking); err != nil {
		return err
	}

	// STEP 4: Add room rent charge
	nights := int(req.CheckoutDate.Sub(req.CheckInnDate).Hours() / 24)
	if nights <= 0 {
		nights = 1
	}

	charge := &hospitalitymodels.HotelBookingCharge{
		ID:                 uuid.New(),
		HotelRoomBookingID: booking.ID,
		ChargeType:         "room_rent",
		Description:        fmt.Sprintf("Room rent for %d nights", nights),
		Amount:             float64(room.Price) * float64(nights),
	}

	if err := r.hotelRepo.CreateCharges(ctx, charge); err != nil {
		return err
	}

	if req.AdvancedPayment > 0 {
		bookingPayment := &hospitalitymodels.HotelRoomBookingPayment{
			ID:                 uuid.New(),
			HotelRoomBookingID: booking.ID,
			Amount:             req.AdvancedPayment,
			Mode:               "CASH",
			Status:             "SUCCESS",
			Reference:          "CASH",
			CreatedAt:          time.Now(),
		}

		if err := r.hotelRepo.CreateOrAddBookingPayment(ctx, bookingPayment); err != nil {
			return err
		}
	}

	return nil
}

func (r *hotelService) GetRoomBookingDetails(ctx context.Context, roomId, orgID uuid.UUID) (*hospitalitymodels.HotelRoomBooking, error) {

	// 1. Validate room belongs to this org
	room, err := r.hotelRepo.GetRoomByID(ctx, roomId)
	if err != nil {
		return nil, errors.New("room not found")
	}

	if room.OrganizationID != orgID {
		return nil, errors.New("room does not belong to this organization")
	}

	// 2. Fetch booking details
	booking, err := r.hotelRepo.GetBookingDetailsByRoomId(ctx, roomId)
	if err != nil {
		return nil, errors.New("no active booking found for this room")
	}

	return booking, nil
}
