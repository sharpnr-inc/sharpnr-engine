package hotelservices

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"sharpnr-engine/sere/events"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CreateRoomInHotel implements HotelService.
func (h *hotelService) CreateRoomInHotel(
	ctx context.Context,
	orgID uuid.UUID,
	req hoteldtos.CreateRoomInFloor_REQ) error {

	var room *hospitalitymodels.Room

	err := h.hotelRepo.WithTx(ctx, func(txRepo hotelrepo.HotelRepositoryInterface) error {
		// -------------------------
		// Validation
		// -------------------------
		if strings.TrimSpace(req.RoomName) == "" {
			return errors.New("room name is required")
		}
		if req.FloorId == "" {
			return errors.New("floor id is required")
		}
		if req.InitialPrice.Price <= 0 {
			return errors.New("initial room price must be greater than zero")
		}

		// -------------------------
		// Load hotel
		// -------------------------
		hotel, err := txRepo.GetHotelByOrgID(ctx, orgID)
		if err != nil {
			return err
		}

		// -------------------------
		// Load & validate floor
		// -------------------------
		floor, err := txRepo.GetFloorByID(ctx, req.FloorId)
		if err != nil {
			return err
		}

		if floor.HotelID != hotel.ID {
			return errors.New("floor does not belong to hotel")
		}

		// -------------------------
		// Create Room (DISABLED first)
		// -------------------------
		room = &hospitalitymodels.Room{
			ID:             uuid.New(),
			FloorID:        floor.ID,
			HotelID:        hotel.ID,
			OrganizationID: orgID,
			RoomName:       req.RoomName,
			Description:    req.Description,
			BedSize:        req.BedSize,
			Amenities:      req.Amenities,
			Status:         hospitalityconstants.Disabled,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := txRepo.CreateRoom(ctx, room); err != nil {
			return err
		}
		// -------------------------
		// Create Initial Room Price
		// -------------------------
		roomPrice := &hospitalitymodels.HotelRoomPrice{
			ID:            uuid.New(),
			RoomID:        room.ID,
			HotelID:       hotel.ID,
			Price:         int(req.InitialPrice.Price),
			EffectiveFrom: req.InitialPrice.EffectiveFrom,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := txRepo.CreatePriceForRoom(ctx, roomPrice); err != nil {
			return err
		}

		// -------------------------
		// Enable Room
		// -------------------------
		if err := txRepo.UpdateRoomStatus(
			ctx,
			hotel.ID,
			room.ID,
			string(hospitalityconstants.Available),
		); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	h.sere.Publish(ctx, "RoomCreated", events.RoomCreated{
		RoomID:  room.ID,
		HotelID: room.HotelID,
		Name:    room.RoomName,
	})

	return nil
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
