package orgservices

import (
	"errors"
	"time"

	"github.com/google/uuid"
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/interfaces"
	"github.com/sharpnr-inc/sharpnr-engine/mappers"
	"github.com/sharpnr-inc/sharpnr-engine/utils"
	"golang.org/x/crypto/bcrypt"
)

type organizationService struct {
	repo interfaces.OrganizationRepository
}

func NewOrganizationService(repo interfaces.OrganizationRepository) *organizationService {
	return &organizationService{repo: repo}
}

func (s *organizationService) RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error) {
	org := mappers.ToOrganizationModel(req)

	// Generate UUID for internal use
	org.SharpnrOrgUUID = uuid.New()

	// Generate unique SharpnrOrgID (16 char string, e.g. SHR0ABCDE12345678)
	sharpID, err := utils.GenerateUniqueSharpnrOrgID(s.repo)
	if err != nil {
		return nil, "", err
	}
	org.SharpnrOrgID = sharpID

	// Set timestamps
	org.CreatedAt = time.Now().Unix()
	org.UpdatedAt = org.CreatedAt

	// Hash passwords
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}
	org.Password = string(hashedPassword)

	hashedMasterPassword, err := bcrypt.GenerateFromPassword([]byte(req.MasterPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to hash master password")
	}
	org.MasterPassword = string(hashedMasterPassword)

	// Persist to DB
	if err := s.repo.Create(org); err != nil {
		return nil, "", err
	}

	// Generate JWT token using UUID & org name (you implement this)
	token, err := utils.GenerateJWTToken(org.SharpnrOrgUUID, org.Name)
	if err != nil {
		return nil, "", err
	}

	respDTO := mappers.ToOrganizationRegistrationResponseDTO(org)
	return &respDTO, token, nil
}
