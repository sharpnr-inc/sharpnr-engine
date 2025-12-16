package hotelservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *hotelService) CreateFloor(ctx context.Context, orgID uuid.UUID, req hospitalitydtos.CreateFloorInHotel_REQ) (*hospitalitydtos.CreateFloorInHotel_RESPONSE, error) {
	// ✅ Validate hotel exists
	hotel, err := r.hotelRepo.GetHotelByID(ctx, uuid.MustParse(req.HotelId))
	if err != nil {
		return nil, err
	}
	if hotel == nil {
		return nil, errors.New("hotel not found")
	}

	hotelUUID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return nil, errors.New("invalid hotel id")
	}

	// ✅ Create new floor
	floor := hospitalitymodels.HotelFloor{
		ID:             uuid.New(),
		HotelID:        hotelUUID,
		OrganizationID: orgID,
		FloorNumber:    hotel.NumberOfFloors + 1, // auto increment floor number
	}

	// ✅ Save floor in repo
	if err := r.hotelRepo.CreateFloorInHotel(ctx, &floor); err != nil {
		return nil, err
	}

	// // ✅ Update hotel floor count (optional)
	// hotel.NumberOfFloors += 1
	// if err := r.hotelRepo.UpdateHotel(ctx, hotel); err != nil {
	// 	fmt.Println("Warning: could not update floor count:", err)
	// }

	// ✅ Prepare response
	resp := &hospitalitydtos.CreateFloorInHotel_RESPONSE{
		HotelId:   req.HotelId,
		FloorName: req.FloorName,
		FloorID:   floor.ID,
	}

	return resp, nil
}
