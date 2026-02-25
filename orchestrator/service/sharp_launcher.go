package orchestrator_service

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetOrganizationSharpLauncher returns the list of of options or the sharpnr launcher for given organizationID.
// This is used by the sharpnr launcher to determine which options to show for a given organization.
func (s *orchestratorService) GetOrganizationSharpLauncher(ctx context.Context, organizationID uuid.UUID) ([]string, error) {
	// first we need to get the organization from the database
	// Check if the organization exists
	exists, err := s.organizationService.VerifyOrgByID(ctx, organizationID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}

	// then we will dugn into the properties of the organization to fetch how many properties they have
	// then we fetch how many domain they domain if it is the hospitality domain
	// then we will use the enrch hospitality with exiting  function to determine which options to show for the sharpnr launcher

	return []string{}, nil

}
