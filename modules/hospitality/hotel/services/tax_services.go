package hotelservices

import (
	"context"
	"errors"
	"math"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"time"

	"github.com/google/uuid"
)

func (s *hotelService) CreateTaxRule(ctx context.Context, orgId uuid.UUID, req hoteldtos.CreateHotelTaxRuleRequest) error {

	return s.hotelRepo.WithTx(ctx, func(txRepo hotelrepo.HotelRepositoryInterface) error {
		hotel, err := txRepo.GetHotelByOrgID(ctx, orgId)
		if err != nil {
			return err
		}

		if req.TaxRate <= 0 {
			return errors.New("tax rate must be greater than zero")
		}

		if req.EffectiveFrom.IsZero() {
			return errors.New("effectiveFrom is required")
		}

		if req.ConditionType == hospitalityconstants.TaxConditionPriceSlab {
			if req.MinValue == nil && req.MaxValue == nil {
				return errors.New("either minValue or maxValue must be provided for PRICE_SLAB")
			}
			if req.MinValue != nil && req.MaxValue != nil &&
				*req.MinValue > *req.MaxValue {
				return errors.New("minValue cannot be greater than maxValue")
			}
		}

		if len(req.Components) == 0 {
			return errors.New("at least one tax component is required")
		}

		var compSum float64
		for _, c := range req.Components {
			if c.Rate <= 0 {
				return errors.New("tax component rate must be greater than zero")
			}
			compSum += c.Rate
		}

		if math.Abs(compSum-req.TaxRate) > 0.01 {
			return errors.New("sum of tax components must equal total tax rate")
		}

		taxRule := &hospitalitymodels.HotelTaxRule{
			ID:            uuid.New(),
			HotelID:       hotel.ID,
			RuleName:      req.RuleName,
			TaxType:       req.TaxType,
			ConditionType: req.ConditionType,
			MinValue:      req.MinValue,
			MaxValue:      req.MaxValue,
			TaxRate:       req.TaxRate,
			EffectiveFrom: req.EffectiveFrom,
			EffectiveTo:   req.EffectiveTo,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := txRepo.CreateTaxRuleOfHotel(ctx, taxRule); err != nil {
			return err
		}

		for _, comp := range req.Components {
			taxComp := &hospitalitymodels.HotelTaxComponent{
				ID:             uuid.New(),
				HotelTaxRuleID: taxRule.ID,
				ComponentType:  comp.ComponentType,
				Rate:           comp.Rate,
				CreatedAt:      time.Now(),
			}

			if err := txRepo.AddTaxComponentInTaxRule(ctx, taxComp); err != nil {
				return err
			}
		}

		return nil
	})
}
