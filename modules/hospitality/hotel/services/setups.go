package hotelservices

import (
	"context"
	"fmt"
	hotel_constants "sharpnr-engine/modules/hospitality/hotel/constants"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
)

func (c *hotelService) SetupFloors(ctx context.Context, req *hotel_contracts.FloorInput) (floorID uuid.UUID, err error) {

	input := &hotel_models.HotelFloor{
		ID:          uuid.New(),
		HotelID:     req.HotelID,
		Name:        req.FloorName,
		FloorNumber: req.FloorNumber,
		CreatedAt:   time.Now(),
	}

	if err := c.hotelRepo.CreateFloor(ctx, input); err != nil {
		return uuid.Nil, err
	}

	return input.ID, nil
}

func (c *hotelService) SetupRoom(ctx context.Context, req *hotel_contracts.RoomInput) (hotelRoomID uuid.UUID, err error) {
	roomInput := &hotel_models.HotelRoom{
		ID:           uuid.New(),
		HotelID:      req.HotelID,
		HotelFloorID: req.FloorID,
		RoomCode:     req.Code,
		Name:         req.Name,
		Status:       req.Status,
		BedType:      req.BedType,
		MaxGuests:    req.MaxGuests,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := c.hotelRepo.CreateRoom(ctx, roomInput); err != nil {
		return uuid.Nil, err
	}

	return roomInput.ID, nil
}



func (c *hotelService) SetUpLayout(ctx context.Context, layout []hotel_contracts.FloorInput) error {
	for _, floorReq := range layout {
		// 1. Create the Floor
		floorID, err := c.SetupFloors(ctx, &floorReq)
		if err != nil {
			return fmt.Errorf("failed to setup floor %s: %w", floorReq.FloorName, err)
		}

		// 2. Generate Rooms for this floor
		for i := 1; i <= floorReq.NumberOfRooms; i++ {
			roomReq := &hotel_contracts.RoomInput{
				HotelID:   floorReq.HotelID,
				FloorID:   floorID,
				Name:      fmt.Sprintf("Room %d%02d", floorReq.FloorNumber, i),
				Code:      fmt.Sprintf("%d%02d", floorReq.FloorNumber, i),
				Status:    hotel_constants.Available, // Assuming a default constant
				BedType:   "Standard",                // Default or adjust as needed
				MaxGuests: 2,                         // Default or adjust as needed
			}

			if _, err := c.SetupRoom(ctx, roomReq); err != nil {
				return fmt.Errorf("failed to setup room %d on floor %s: %w", i, floorReq.FloorName, err)
			}
		}
	}

	return nil
}

func (c *hotelService) AddHotelGenDetails(ctx context.Context, input *hotel_contracts.AddHotelGeneralInfoInput) error {

	opTiming := &hotel_models.HotelOperationsTiming{
		ID:           uuid.New(),
		HotelID:      input.HotelID,
		Is24X7:       input.Is24X7,
		CheckInTime:  input.HotelOpeningTime,
		CheckOutTime: input.HotelClosingTime,
	}

	if err := c.hotelRepo.CreateHotelOperationalTiming(ctx, opTiming); err != nil {
		return err
	}

	updateGenDetails := hotel_contracts.HotelMainGenDetailsInput{
		HotelID:      input.HotelID,
		Name:         input.HotelName,
		IsConfigured: true,
	}

	if err := c.hotelRepo.UpdateHotelGenDetails(ctx, &updateGenDetails); err != nil {
		return err
	}

	return nil

}
