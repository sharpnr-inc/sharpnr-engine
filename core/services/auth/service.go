package services

import (
	"errors"

	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/interfaces"
	"sharpnr-engine/core/mappers"
	"sharpnr-engine/shared/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// This Struct Return the Repository Interface or Contract
// Which have 4 Signature functions
type organisationService struct {
	repo interfaces.CoreOrganizationRepository
}

// This service for the implemeting the Auth and Registration for the Organisations
func CreateNewOrganization(repo interfaces.CoreOrganizationRepository) *organisationService {
	return &organisationService{repo: repo}
}

func (s *organisationService) RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error) {
	org := mappers.ToOrganizationModel(req)
	org.SharpnrOrgUUID = uuid.New()
	org.OrgOprationType = "kakajs"

	sharpID, err := utils.GenerateUniqueSharpnrOrgID(s.repo)
	if err != nil {
		return nil, "", err
	}

	org.SharpnrOrgID = sharpID

	hashedPasword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to Hash the Password")
	}

	org.Password = string(hashedPasword)

	hashedMasterPasword, err := bcrypt.GenerateFromPassword([]byte(req.MasterPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to Hash the Password")
	}

	org.MasterPassword = string(hashedMasterPasword)

	if err := s.repo.Create(org); err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWTToken(org.SharpnrOrgUUID, org.Name)
	if err != nil {
		return nil, "", err
	}

	respDTO := mappers.ToOrganizationRegistrationResponseDTO(org)
	return &respDTO, token, nil

}
