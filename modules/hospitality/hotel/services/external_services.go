package hotelservices

import (
	"context"
	"fmt"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/dtos"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
)

func (r *hotelService) GetHotelDetails(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*hotel_contracts.HotelsReader, error) {
	hotel, err := r.hotelRepo.GetHotelsByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	items := make([]hotel_contracts.HotelReader, len(hotel))

	for i, h := range hotel {
		items[i] = hotel_contracts.HotelReader{
			ID:           h.ID,
			Name:         h.Name,
			Status:       h.IsActive,
			IsConfigured: h.IsConfigured,
		}
	}

	return &hotel_contracts.HotelsReader{
		Hotels: items,
	}, nil
}

func (c *hotelService) GetHotelLayout(ctx context.Context, hotelID uuid.UUID) (res *hoteldtos.HotelLayoutResponseDTO, err error) {
	// 1. Fetch all floors for the hotel
	floors, err := c.hotelRepo.GetFloorsByHotelID(ctx, hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch floors: %w", err)
	}

	var layout []hoteldtos.HotelFloorResponseDTO

	// 2. Loop through floors to gather rooms
	for _, f := range floors {
		rooms, err := c.hotelRepo.GetHotelRoomsByFloorID(ctx, f.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch rooms for floor %s: %w", f.Name, err)
		}

		var roomDTOs []hoteldtos.HotelRoomResponseDTO
		for _, r := range rooms {
			roomDTOs = append(roomDTOs, hoteldtos.HotelRoomResponseDTO{
				RoomID:   r.ID,
				RoomName: r.Name,
				Status:   string(r.Status),
				Code:     r.RoomCode,
				BedType:  r.BedType,
				// Price: r.Price, // Note: Ensure Price exists in your hotel_models.HotelRoom
			})
		}

		layout = append(layout, hoteldtos.HotelFloorResponseDTO{
			FloorID:     f.ID,
			FloorName:   f.Name,
			FloorNumber: f.FloorNumber,
			Rooms:       roomDTOs,
		})
	}

	return &hoteldtos.HotelLayoutResponseDTO{
		Layout: layout,
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
