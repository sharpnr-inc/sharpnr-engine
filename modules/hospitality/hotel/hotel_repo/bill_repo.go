package hotelrepo

import (
	"context"
	"errors"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) GetBillByBookingId(
	ctx context.Context,
	bookingId uuid.UUID,
) (*hotel_models.HotelBill, error) {

	var bill hotel_models.HotelBill

	err := d.db.WithContext(ctx).
		Where("hotel_room_booking_id = ?", bookingId).
		First(&bill).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &bill, nil
}

func (d *hotelRepository) CreateBookingBill(ctx context.Context, bill *hotel_models.HotelBill) error {
	if bill == nil {
		return errors.New("bill cannot be nil")
	}

	if bill.Status == "" {
		bill.Status = "DRAFT"
	}

	return d.db.WithContext(ctx).Create(bill).Error
}
