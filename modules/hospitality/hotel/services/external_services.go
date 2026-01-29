package hotelservices

import (
	"context"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
)

// TODO --> This is used inside the function where we fetch the hotel for the Property please
// ?? Please Fix the this
func (r *hotelService) GetHotelDetails(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*hotel_contracts.HotelsReader, error) {
	hotel, err := r.hotelRepo.GetHotelsByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	items := make([]hotel_contracts.HotelReader, len(hotel))

	for i, h := range hotel {
		items[i] = hotel_contracts.HotelReader{
			ID:     h.ID,
			Name:   h.Name,
			Status: h.IsActive,
		}
	}

	return &hotel_contracts.HotelsReader{
		Hotels: items,
	}, nil
}

func (h *hotelService) SetupDefaultHotel(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error) {
	defHotel := &hotel_models.Hotel{
		ID:             uuid.New(),
		OrganizationID: orgID,
		PropertyID:     propertyID,
		Name:           "Default Hotel",
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := h.hotelRepo.CreateHotel(ctx, defHotel); err != nil {
		return uuid.Nil, err
	}

	return defHotel.ID, nil
}
