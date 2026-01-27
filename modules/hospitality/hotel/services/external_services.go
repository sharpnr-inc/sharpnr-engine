package hotelservices

import (
	"context"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"

	"github.com/google/uuid"
)

// TODO --> This is used inside the function where we fetch the hotel for the Property please
// ?? Please Fix the this
func (r *hotelService) GetHotelDetails(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*hotel_contracts.HotelsReader, error) {
	hotel, err := r.hotelRepo.GetHotelByOrgID(ctx, orgID)
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
