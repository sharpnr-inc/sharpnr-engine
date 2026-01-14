package hotelservices

import (
	"context"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
)

func (s *hotelService) SearchGuests(ctx context.Context, orgId uuid.UUID, qry string) ([]hoteldtos.HotelGuestResponseDTO, error) {
	hotel, err := s.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Search guests scoped to hotel
	guests, err := s.hotelRepo.SearchHotelGuest(
		ctx,
		hotel.ID,
		qry,
	)
	if err != nil {
		return nil, err
	}

	// 3️⃣ Map model → DTO
	result := make([]hoteldtos.HotelGuestResponseDTO, 0, len(guests))
	for _, g := range guests {
		result = append(result, hoteldtos.HotelGuestResponseDTO{
			ID:        g.ID,
			Name:      g.Name,
			Mobile:    g.Mobile,
			DOB:       g.DOB,
			Gender:    g.Gender,
			AadhaarId: g.AadhaarId,
		})
	}

	// 4️⃣ Return DTOs
	return result, nil
}

func (s *hotelService) CreateGuestInHotel(
	ctx context.Context,
	orgId uuid.UUID,
	req hoteldtos.CreateHotelGuestsRequest,
) error {
	hotel, err := s.hotelRepo.GetHotelByOrgID(ctx, orgId)
	if err != nil {
		return err
	}

	guest := &hotel_models.HotelGuest{
		ID:        uuid.New(),
		HotelId:   hotel.ID,
		OrgId:     orgId,
		Name:      req.Name,
		Mobile:    req.Mobile, // ✅ FIXED
		Email:     req.Email,
		DOB:       req.DOB,
		AadhaarId: req.AadhaarId,
		Gender:    "M/F",
	}

	return s.hotelRepo.CreatePrimaryGuest(ctx, guest)
}
