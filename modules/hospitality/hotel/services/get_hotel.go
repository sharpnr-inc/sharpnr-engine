package hotelservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (r *hotelService) GetHotelDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetHotel_RESPONSE, error) {
	// Fetch hotel by organization ID
	hotel, err := r.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return nil, err
	}
	if hotel == nil {
		return nil, errors.New("hotel not found for this organization")
	}

	// Prepare DTO response
	response := &hospitalitydtos.GetHotel_RESPONSE{
		HotelId:   hotel.ID.String(),
		HotelName: hotel.Name,
	}

	for _, f := range hotel.Floors {
		response.Floors = append(response.Floors, hospitalitydtos.FloorResponse{
			FloorID:     f.ID,
			FloorNumber: f.FloorNumber,
			FloorName:   f.FloorName,
		})
	}

	return response, nil
}
