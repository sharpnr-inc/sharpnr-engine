package signupservice

import (
	"context"
	dto "sharpnr-engine/core/dtos"
)

func (s *signupService) CompleteOrgRegistration(
	ctx context.Context,
	req dto.Org_Complete_Regesistration_REQ_DTO,
	org_email, org_sharpnr_id string,
) (*dto.Org_Complete_Regesistration_RESPONSE_DTO, string, error) {

	return nil, "", nil
}
