package hotelservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CreateRoomInHotel implements HotelService.
func (h *hotelService) CreateRoomInHotel(
	ctx context.Context,
	orgID uuid.UUID,
	req hospitalitydtos.CreateRoomInFloor_REQ) (*hospitalitydtos.CreateRoomInFloor_RES, error) {

	if strings.TrimSpace(req.Description) == "" ||
		strings.TrimSpace(req.RoomName) == "" {
		return nil, errors.New("description, room name and floor id are required")
	}

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
		ID:             uuid.New(),
		FloorID:        floor.ID,
		HotelID:        floor.HotelID,
		OrganizationID: orgID,
		RoomName:       req.RoomName,
		Status:         hospitalityconstants.Available,
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

/*************************************
 * Function for Updating Room status *
 *************************************/

func (r *hotelService) UpdateRoomStatus(
	ctx context.Context,
	orgId, hotelId, roomId uuid.UUID,
	status string,
) error {

	// 1️⃣ Fetch Room
	room, err := r.hotelRepo.GetRoomByID(ctx, roomId)
	if err != nil {
		return err
	}

	// 2️⃣ Ownership Validation
	if room.OrganizationID != orgId {
		return errors.New("room does not belong to this organization")
	}

	if room.HotelID != hotelId {
		return errors.New("room does not belong to this hotel")
	}

	// 3️⃣ Build fields to update
	fields := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	// 4️⃣ Update using dynamic fields
	if err := r.hotelRepo.UpdateRoom(ctx, fields, roomId); err != nil {
		return err
	}

	return nil
}
