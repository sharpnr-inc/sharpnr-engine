package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *hotelService) GetFrontDeskTerminal(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.HotelBookingTerminal_Response, error) {

	hotelFrontDesk, err := s.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return nil, err
	}

	res := hospitalitydtos.HotelBookingTerminal_Response{
		Hotel_Id: hotelFrontDesk.ID,
		Org_Id:   hotelFrontDesk.OrganizationID,
	}

	for _, f := range hotelFrontDesk.Floors {
		floorRes := hospitalitydtos.FloorResponse{
			FloorID:     f.ID,
			FloorNumber: f.FloorNumber,
			FloorName:   f.FloorName,
		}

		for _, room := range f.Rooms {
			roomRes := hospitalitydtos.RoomResponse{
				RoomID:      room.ID,
				RoomName:    room.RoomName,
				Description: room.Amenities,
				Status:      string(room.Status),
			}

			// 🔹 Fetch room price
			roomPrice, err := s.hotelRepo.GetRoomPriceByRoomId(ctx, room.ID)
			if err != nil {
				return nil, err // or log & continue based on your policy
			}

			if roomPrice != nil {
				roomRes.RoomPrice = int(roomPrice.Price)
			}

			// ✅ Proper pointer + uuid.Nil check
			if room.CurrentBookingID != nil && *room.CurrentBookingID != uuid.Nil {
				roomRes.CurrentBookingId = room.CurrentBookingID
			}

			floorRes.Rooms = append(floorRes.Rooms, roomRes)
		}

		res.Floors = append(res.Floors, floorRes)

	}
	return &res, nil
}
