package signupservice

import (
	"context"
	dto "sharpnr-engine/core/dtos"
)

func (s *signupService) VerifyOrgRegistration(ctx context.Context, tokn string, code string) (*dto.OrgVerifyOTP_Res_Dto, bool, error) {
	return nil, false, nil
}
