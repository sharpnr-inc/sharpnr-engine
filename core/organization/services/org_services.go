package organization_services

import (
	"context"
	"errors"
	"fmt"
	organization_contracts "sharpnr-engine/core/organization/contracts"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This return tell org Exist or not
func (r *organizationService) IsOrgExists(ctx context.Context, rootEmail string) (bool, error) {
	// 1. Attempt to fetch the organization by email
	_, err := r.orgMainRepo.GetOrganizationByEmail(ctx, rootEmail)

	if err != nil {
		// 2. If the error is "Record Not Found", the org does NOT exist (Return false, nil)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		// 3. If it's a real database error (connection, etc.), return the error
		return false, fmt.Errorf("failed to check organization existence: %w", err)
	}

	// 4. If err is nil, it means a record was found
	return true, nil
}

//----------------------------------------------------------------------------------------------------------------//

func (s *organizationService) UpdateOrganizationName(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, name string) error {
	return s.orgMainRepo.UpdateOrgName(ctx, orgID, name)
}

/* --------------------------------- <****> --------------------------------- */

func (s *organizationService) GetOrgDetailsByID(
	ctx context.Context,
	orgID,
	emplID uuid.UUID,
) (*organization_contracts.OrganizationResult, error) {

	var orgData *organization_contracts.OrganizationResult

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		org, err := s.orgMainRepo.GetOrganizationByID(txCtx, orgID)
		if err != nil {
			return err
		}

		orgData = &organization_contracts.OrganizationResult{
			ID:        org.ID,
			SharpID:   org.SharpnrID,
			Name:      org.Name,
			RootEmail: org.RootEmail,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return orgData, nil
}

/* --------------------------------- <****> --------------------------------- */
