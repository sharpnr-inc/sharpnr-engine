package orgservices

import (
	"context"
	dto "sharpnr-engine/core/dtos"
)

func (s *orgLoginService) Login(ctx context.Context, req dto.OrgAuthRequestDTO) (*dto.OrgAuthResponseDTO, error) {

	return nil, nil

	// s.repo.FindByEmailOrPhone()
}
