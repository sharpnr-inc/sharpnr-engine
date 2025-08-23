package orgservices

import (
	"context"
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/utils"
	"time"

	"github.com/google/uuid"
)

func (s *orgRegistrationService) VerifyOrgRegistration(ctx context.Context, tokn string, code string) (*dto.OrgVerifyOTP_Res_Dto, bool, error) {
	// for Now Token is Email

	validateOtp, err := s.repo.EmailOrPhoneExistsInRedis(ctx, tokn, code)
	if err != nil {
		return nil, false, fmt.Errorf("failed to check redis: %w", err)
	}

	// Then We need to move this Org Data in the Main DB and Remove from Redis
	// Before Moving to the Data to MainDB we need to Assign the UUid and SharpnrOrgID
	// Then we Will Push That Data in the Main DB

	if !validateOtp {
		return nil, false, fmt.Errorf("invalid otp for %s", tokn)
	}

	orgData, err := s.repo.GetOrgFromRedisDB(ctx, tokn)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get org data from redis: %w", err)
	}

	orgID, err := utils.GenerateUniqueSharpnrOrgID(s.repo)
	if err != nil {
		return nil, false, fmt.Errorf("failed to generate org id: %w", err)
	}

	org := &models.Organization{
		ID:            uuid.New(),
		SharpnrOrgID:  orgID,
		Name:          orgData.Name,
		RootEmail:     orgData.Email,
		ContactNumber: orgData.Phone,
		// set required flags
		ISRegistered: true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.CreateOrgInMainDB(ctx, org); err != nil {
		return nil, false, fmt.Errorf("failed to save org in DB: %w", err)
	}

	res := &dto.OrgVerifyOTP_Res_Dto{
		SharpnrOrgID: org.SharpnrOrgID,
		Name:         org.Name,
	}

	return res, true, nil

}
