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

	// tableRes := make([]hospitalitydtos.BarTablesResponse, 0)

	// for _, t := range bar.Tables {
	// 	tableRes = append(tableRes, hospitalitydtos.BarTablesResponse{
	// 		TableName:   t.Name,
	// 		TableId:     t.ID,
	// 		TableStatus: t.Status,
	// 	})
	// }

	// flooresp := make([]hospitalitydtos.BarFloorResponse, 0)

	// for _, bf := range bar.Floors {
	// 	flooresp = append(flooresp, hospitalitydtos.BarFloorResponse{
	// 		FloorId: bf.ID,
	// 		// FloorName: bf.FloorNumber,
	// 		Tables: tableRes,
	// 	})
	// }

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
		DisplayName:                    bar.DisplayName,
		LegalName:                      bar.LegalName,
		CinNumber:                      bar.CINNumber,
		FssaiNumber:                    bar.FSSAINumber,
		GstNumber:                      bar.GSTNumber,
		IsMultifloor:                   bar.IsMultifloor,
		AddressLine:                    bar.AddressLine,
		City:                           bar.City,
		State:                          bar.State,
		Country:                        bar.Country,
		Pincode:                        bar.Pincode,
	}

	return res, nil

}
