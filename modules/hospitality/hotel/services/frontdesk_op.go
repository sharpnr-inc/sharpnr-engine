package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *hotelService) GetFrontDeskTerminal(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.HotelBookingTerminal_Response, error) {

	// Todo We need to find the hotel based on orgId and HotelID and response

	hotelFrontDesk, err := s.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return nil, err
	}

	// if hotelId != hotelFrontDesk.ID {
	// 	return nil, errors.New("hotel is not available in the Org")
	// }

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
			floorRes.Rooms = append(floorRes.Rooms, hospitalitydtos.RoomResponse{
				RoomID:           room.ID,
				RoomName:         room.RoomName,
				Description:      room.Description,
				Status:           string(room.Status),
				CurrentBookingId: room.CurrentBookingID,
			})
		}

		res.Floors = append(res.Floors, floorRes)

	}
	return &res, nil
}
