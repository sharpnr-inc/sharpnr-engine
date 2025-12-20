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
		HotelId:          hotel.ID.String(),
		SharpId:          hotel.Organization.SharpnrID,
		HotelName:        hotel.Name,
		HotelDisplayName: hotel.DisplayName,
		HotelLegalName:   hotel.LegalName,
		KitchenId:        *hotel.KitchenID,
		IsKitchen:        hotel.IsKitchen,
		Address:          hotel.Address,
		City:             hotel.City,
		State:            hotel.State,
		HotelGstNumber:   hotel.GstinNumber,
		HotelFSSAINumber: hotel.FssaiLicenseNumber,
	}

	for _, f := range hotel.Floors {
		floorRes := hospitalitydtos.FloorResponse{
			FloorID:     f.ID,
			FloorNumber: f.FloorNumber,
			FloorName:   f.FloorName,
		}

		// ✅ Add rooms for this floor
		for _, room := range f.Rooms {
			floorRes.Rooms = append(floorRes.Rooms, hospitalitydtos.RoomResponse{
				RoomID:           room.ID,
				RoomName:         room.RoomName,
				RoomPrice:        room.Price, // Add if you have room price field in model
				Description:      room.Amenities,
				Status:           string(room.Status),
				CurrentBookingId: room.CurrentBookingID,
			})
		}

		response.Floors = append(response.Floors, floorRes)
	}

	return response, nil
}

func (r *hotelService) UpdateHotelDetails(
	ctx context.Context,
	orgId uuid.UUID,
	req hospitalitydtos.Hotel_UpdateOrConfigure_REQUEST,
) error {

	// 1️⃣ Fetch hotel (org safety)
	hotel, err := r.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return err
	}

	// 2️⃣ Build update map (PATCH behavior)
	updates := map[string]interface{}{}

	if req.DisplayName != nil {
		updates["display_name"] = *req.DisplayName
	}
	if req.LegalName != nil {
		updates["legal_name"] = *req.LegalName
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.City != nil {
		updates["city"] = *req.City
	}
	if req.State != nil {
		updates["state"] = *req.State
	}
	if req.Pincode != nil {
		updates["pincode"] = *req.Pincode
	}
	if req.HotelGstNumber != nil {
		updates["gstin_number"] = *req.HotelGstNumber
	}
	if req.HotelFSSAINumber != nil {
		updates["fssai_license_number"] = *req.HotelFSSAINumber
	}

	// Optional: mark hotel configured when first updated
	if !hotel.IsConfigured && len(updates) > 0 {
		updates["is_configured"] = true
	}

	if len(updates) == 0 {
		return errors.New("no fields provided for update")
	}

	// 3️⃣ Call repository
	return r.hotelRepo.UpdateHotelFields(ctx, hotel.ID, updates)
}
