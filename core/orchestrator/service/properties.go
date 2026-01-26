package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"

	"github.com/google/uuid"
)

func (s *orchestratorService) GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error) {

	res, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	readers := make([]orchestrator_dtos.PropertiesReaderSlice, 0, len(res.Properties))

	for _, p := range res.Properties {
		readers = append(readers, orchestrator_dtos.PropertiesReaderSlice{
			ID:     p.ID,
			Name:   p.Name,
			City:   p.City,
			State:  p.State,
			Domain: p.Domain,
			Code:   p.Code,
		})
	}

	return &orchestrator_dtos.PropertiesOfOrganization_ResDto{
		Properties: readers,
	}, nil
}
