package barservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *barServices) GetBarDetailsConfig(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.BAR_UpdateBarDetails_Response, error) {
	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return nil, err
	}

	if bar == nil {
		return nil, errors.New("bar not found in the DB")
	}

	res := &hospitalitydtos.BAR_UpdateBarDetails_Response{
		BarDisplayName:                 bar.DisplayName,
		BarLegalName:                   bar.LegalName,
		OpeningTime:                    bar.OpeningTime,
		ClosingTime:                    bar.ClosingTime,
		BarID:                          bar.ID,
		OrgId:                          bar.OrganizationID,
		IsKitchen:                      bar.IsKitchen,
		KitchenId:                      *bar.KitchenID,
		IsShareLegalDetailsWithMainOrg: bar.IsShareLegalDetailsWithMainOrg,
	}

	return res, nil

}
