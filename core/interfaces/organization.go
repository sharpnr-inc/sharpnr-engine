package interfaces

import (
	"context"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
)

// This is Interface for the Registering the Organization
// This will have 4 funcs
type OrgRegistrationService interface {
	//Fisrt Step to Initialising the Registration Process
	InitiateOrgRegistration(ctx context.Context, req dto.OrgInitialRegistration_Request_Dto) (*dto.OrgInitialRegistration_Response_Dto, string, error)
	VerifyOrgRegistration(ctx context.Context, token, code string) (*dto.OrgVerifyOTP_Res_Dto, bool, error)
	// CompleteOrgRegistration(token string) (*dto.OrgRegistrationComplete_Response_Dto, error)
}

type IOrgRegistrationRepository interface {
	CreateOrgInMainDB(ctx context.Context, tempOrg *models.Organization) error
	CreateOrgInRedisDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error

	EmailOrPhoneExistsInRedis(ctx context.Context, email string, phone string) (bool, error)

	EmailOrPhoneExistsInMainDB(ctx context.Context, email string, phone string) (bool, error)

	GetOrgFromRedisDB(ctx context.Context, email string) (*models.TempOrgRegistrationData, error)

	FindBySharpnrOrgID(id string) (bool, error)
}

type IOrgLogin interface {
}
