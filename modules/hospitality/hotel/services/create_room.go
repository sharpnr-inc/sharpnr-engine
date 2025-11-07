package hotelservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"

	"github.com/google/uuid"
)

// CreateRoomInHotel implements HotelService.
func (h *hotelService) CreateRoomInHotel(ctx context.Context, orgID uuid.UUID, req hospitalitydtos.CreateRoomInFloor_REQ) (*hospitalitydtos.CreateRoomInFloor_RES, error) {
	// ✅ Verify floor exists
	floor, err := h.hotelRepo.GetFloorByID(ctx, req.FloorId)
	if err != nil {
		return nil, err
	}
	if floor == nil {
		return nil, errors.New("floor not found")
	}

	// ✅ Create room model
	room := hospitalitymodels.Room{
		ID:             uuid.New().String(),
		FloorID:        floor.ID,
		HotelID:        floor.HotelID,
		OrganizationID: orgID.String(),
		RoomName:       req.RoomName,
		Status:         hospitalityconstants.Disabled,
		Price:          req.RoomPrice,
	}

	// ✅ Save room
	if err := h.hotelRepo.CreateRoom(ctx, &room); err != nil {
		return nil, err
	}

	// ✅ Prepare response
	res := &hospitalitydtos.CreateRoomInFloor_RES{

		RoomName:  room.RoomName,
		RoomPrice: req.RoomPrice,
	}

	return res, nil
}
