package employeeservice

import (
	"context"
	"errors"

	dto "sharpnr-engine/core/dtos"

	"github.com/google/uuid"
)

func (s *EmployeeServiceImpl) AddEmployeeInOrg(ctx context.Context, org_id string, req dto.CreateEmployeeInOrganisationRequestDTO) (*dto.CreateEmployeeInOrganisationResponseDTO, error) {
	if req.Email == "" || req.Name == "" || req.OrganizationID == "" {
		return nil, errors.New("all fields are required")
	}

	// business logic goes here (e.g., save to DB)
	// For now, just returning a dummy response

	resp := &dto.CreateEmployeeInOrganisationResponseDTO{
		ID:             uuid.New(), // replace with actual ID from DB
		Name:           req.Name,
		Email:          req.Email,
		OrganizationID: req.OrganizationID,
	}

	return resp, nil
}
