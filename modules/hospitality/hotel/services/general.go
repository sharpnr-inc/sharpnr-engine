package hotelservices

import (
	"context"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *hotelService) GetHotelsByOrgID(ctx context.Context, orgID, emplID uuid.UUID) ([]hotel_contracts.Hotels, error) {
	hotels, err := s.hotelRepo.GetHotelsByOrgID(ctx, orgID)
	if err != nil {
		// gorm.ErrRecordNotFound usually shouldn't be an "error" for a list,
		// just return an empty slice unless the business logic dictates otherwise.
		if err == gorm.ErrRecordNotFound {
			return []hotel_contracts.Hotels{}, nil
		}
		return nil, err
	}

	// Pre-allocate slice capacity for better performance
	response := make([]hotel_contracts.Hotels, 0, len(hotels))

	for _, h := range hotels {
		// Since LegalEntityID in contract is a pointer but in model it is a value
		// we create a local copy to get the address safely.
		legalEntityID := h.LegalEntityID

		response = append(response, hotel_contracts.Hotels{
			ID:             h.ID,
			Name:           h.Name,
			OrganizationID: h.OrganizationID,
			PropertyID:     h.PropertyID,
			LegalEntityID:  &legalEntityID,
			IsActive:       h.IsActive,
			IsConfigured:   h.IsConfigured,
			LastUpdatedAt:  h.UpdatedAt,
		})
	}

	return response, nil
}

/* ---------------------------------- <***> --------------------------------- */
