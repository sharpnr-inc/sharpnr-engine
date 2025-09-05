package interfaces

import (
	"context"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

//////////////////////////
// ORGANIZATION REGISTRATION
//////////////////////////

// OrgRegistrationService defines methods for registering an organization
type OrganizationSignUpService interface {
	// InitiateOrgRegistration is the first step to start the registration process
	InitiateOrgRegistration(ctx context.Context, req dto.OrgInitialRegistration_Request_Dto) (*dto.OrgInitialRegistration_Response_Dto, string, error)

	// VerifyOrgRegistration verifies the OTP/token sent to the organization
	VerifyOrgRegistration(ctx context.Context, token, code string) (*dto.OrgVerifyOTP_Res_Dto, bool, error)

	// Optional: CompleteOrgRegistration finalizes the registration
	// CompleteOrgRegistration(token string) (*dto.OrgRegistrationComplete_Response_Dto, error)
}

// IOrgRegistrationRepository defines data access methods for organization registration
type OrganizationRepository interface {
	// CreateOrgInMainDB saves the organization in the main database
	CreateOrgInMainDB(ctx context.Context, tempOrg *models.Organization) error

	// CreateOrgInRedisDB stores temporary registration data in Redis
	CreateOrgInRedisDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error

	// EmailOrPhoneExistsInRedis checks if email or phone exists in Redis
	EmailOrPhoneExistsInRedis(ctx context.Context, email string, phone string) (bool, error)

	// EmailOrPhoneExistsInMainDB checks if email or phone exists in main DB
	EmailOrPhoneExistsInMainDB(ctx context.Context, email string, phone string) (bool, error)

	// GetOrgFromRedisDB retrieves temporary registration data from Redis
	GetOrgFromRedisDB(ctx context.Context, email string) (*models.TempOrgRegistrationData, error)

	// FindBySharpnrOrgID checks if an organization exists by its Sharpnr ID
	FindBySharpnrOrgID(id string) (bool, error)

	//It will get find you the orgId
	GetOrgUUID_By_SharpnrID(sharpnr_id string) (uuid.UUID, error)
}
