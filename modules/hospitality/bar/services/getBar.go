package barservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *barServices) GetBarDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetBarDetails_RESPONSE, error) {
	// 1️⃣ Fetch Bar for this Organization
	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return nil, err
	}
	if bar == nil {
		return nil, errors.New("bar not found for this organization")
	}

	response := &hospitalitydtos.GetBarDetails_RESPONSE{
		BarId:   bar.ID.String(),
		BarName: bar.Name,
		Floors:  []hospitalitydtos.BarFloorResponse{},
		Tables:  []hospitalitydtos.BarTablesResponse{},
	}

	// 2️⃣ Case A: Bar has Floors
	if len(bar.Floors) > 0 {
		for _, floor := range bar.Floors {
			floorResponse := hospitalitydtos.BarFloorResponse{
				FloorId:   floor.ID,
				FloorName: floor.Name,
				Tables:    []hospitalitydtos.BarTablesResponse{},
			}

			for _, table := range floor.Tables {
				floorResponse.Tables = append(floorResponse.Tables, hospitalitydtos.BarTablesResponse{
					TableId:     table.ID,
					TableName:   table.Name,
					TableStatus: table.Status,
				})
			}
			response.Floors = append(response.Floors, floorResponse)
		}
	}

	// 3️⃣ Case B: Bar has no Floors, Tables directly under Bar
	if len(bar.Floors) == 0 && len(bar.Tables) > 0 {
		for _, table := range bar.Tables {
			response.Tables = append(response.Tables, hospitalitydtos.BarTablesResponse{
				TableId:     table.ID,
				TableName:   table.Name,
				TableStatus: table.Status,
			})
		}
	}

	return response, nil
}
