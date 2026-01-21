package hotelrepo

import (
	"context"
	"errors"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) CreateRoomBooking(ctx context.Context, booking *hotel_models.HotelRoomBooking) error {
	if err := d.db.WithContext(ctx).Create(booking).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hotel_models.HotelRoomBookingPayment) error {
	if err := d.db.WithContext(ctx).Create(bookingPayment).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateCharges(ctx context.Context, charges *hotel_models.HotelBookingCharge) error {
	if err := d.db.WithContext(ctx).Create(charges).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) GetBookingDetailsByRoomId(
	ctx context.Context,
	roomId uuid.UUID,
) (*hotel_models.HotelRoomBooking, error) {

	var booking hotel_models.HotelRoomBooking

	err := d.db.WithContext(ctx).
		Preload("Room").
		Preload("PrimaryGuest").
		// Preload("Guests").
		Preload("Charges").
		Preload("Payments").
		Preload("Orders").
		Preload("Orders.Items").
		Preload("Services").
		Preload("Services.Items").
		Where("room_id = ? AND status IN ('RESERVED', 'CHECKED_IN', 'BOOKED')", roomId). // only active bookings
		First(&booking).Error

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (d *hotelRepository) GetBookingByBookingId(ctx context.Context, bookingId uuid.UUID) (*hotel_models.HotelRoomBooking, error) {
	var booking hotel_models.HotelRoomBooking

	err := d.db.WithContext(ctx).
		Where("id = ?", bookingId).
		First(&booking).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &booking, nil
}

func (h *hotelRepository) UpdateBookingStatus(
	ctx context.Context,
	bookingID uuid.UUID,
	status hospitalityconstants.HotelBookingStatusType,
) error {

	return h.db.WithContext(ctx).
		Model(&hotel_models.HotelRoomBooking{}).
		Where("id = ?", bookingID).
		Update("status", status).
		Error
}
