package hotelservices

import (
	"context"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
)

type TaxCalculationResult struct {
	TotalTax   float64
	Components []TaxComponentResult
}

type TaxComponentResult struct {
	ComponentType hospitalityconstants.TaxComponentType
	Rate          float64
	Amount        float64
}

func (s *hotelService) CalculateTaxForBooking(
	ctx context.Context,
	hotelID uuid.UUID,
	baseAmount float64,
) (*TaxCalculationResult, error) {

	// 1️⃣ Decide tax type
	// (for now: ROOM_GST, later you can split per charge)
	taxType := hospitalityconstants.TaxTypeRoomGST

	// 2️⃣ Fetch active tax rules
	rules, err := s.hotelRepo.GetActiveTaxRulesByHotelAndType(
		ctx,
		hotelID,
		taxType,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	if len(rules) == 0 {
		// No tax applicable
		return &TaxCalculationResult{
			TotalTax:   0,
			Components: []TaxComponentResult{},
		}, nil
	}

	// 3️⃣ Pick matching slab
	var matchedRule *hotel_models.HotelTaxRule

	for _, rule := range rules {

		// Min check
		if rule.MinValue != nil && baseAmount < *rule.MinValue {
			continue
		}

		// Max check
		if rule.MaxValue != nil && baseAmount > *rule.MaxValue {
			continue
		}

		matchedRule = &rule
		break
	}

	if matchedRule == nil {
		// No slab matched → no tax
		return &TaxCalculationResult{
			TotalTax:   0,
			Components: []TaxComponentResult{},
		}, nil
	}

	// 4️⃣ Fetch components (CGST / SGST / IGST)
	components, err := s.hotelRepo.GetTaxComponentsByRuleID(
		ctx,
		matchedRule.ID,
	)
	if err != nil {
		return nil, err
	}

	// 5️⃣ Calculate tax
	result := &TaxCalculationResult{
		Components: []TaxComponentResult{},
	}

	for _, comp := range components {

		compTax := (baseAmount * comp.Rate) / 100

		result.Components = append(result.Components, TaxComponentResult{
			ComponentType: comp.ComponentType,
			Rate:          comp.Rate,
			Amount:        compTax,
		})

		result.TotalTax += compTax
	}

	return result, nil
}
