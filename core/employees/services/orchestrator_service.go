package coreEmployeeServices

import (
	"context"
	"fmt"
	"strings"
	"time"

	coreEmployeeConstants "sharpnr-engine/core/employees/constants"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeErrs "sharpnr-engine/core/employees/errs"
	coreEmployeeModels "sharpnr-engine/core/employees/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/shared/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddNewOwnerOrSuperAdmin implements [EmployeeOrchestratorServiceInterface]:
// create a second owner/super-admin in an org that already has one.
//
// TEMPORARY IMPLEMENTATION — the intended flow is invitation-based, the same
// shape as AddEmployeeInBulkInOrg: create the employee PENDING, issue a
// provisioning token, mail the invitee a link, and let them set their own
// password; the account only becomes ACTIVE once they complete it. That is what
// the endpoint's response message ("invitation sent") is written for.
//
// Until the invite mail is wired up, the caller passes the password in the
// request and the account is created ACTIVE and provisioned right here, so a
// second owner can be added at all. When the real flow lands: drop Password
// from AddNewSuperAdminInOrganizationRequestDto, create the employee PENDING
// with a provisioning record instead of an EmployeeAuth row, and move the owner
// role assignment to the moment the invitee finishes provisioning.
//
// The whole thing is one transaction: an employee row with no auth row cannot
// sign in, and one with no Owner role is a super-admin in name only — neither
// is recoverable without manual surgery.
func (s *employeeService) AddNewOwnerOrSuperAdmin(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	req coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto,
) (uuid.UUID, error) {

	firstName := strings.TrimSpace(req.FirstName)
	lastName := strings.TrimSpace(req.LastName)
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if firstName == "" {
		return uuid.Nil, coreEmployeeErrs.EmployeeFirstNameRequiredErr
	}
	if email == "" {
		return uuid.Nil, coreEmployeeErrs.EmployeeEmailRequiredErr
	}
	if req.Password == "" {
		return uuid.Nil, coreEmployeeErrs.EmployeePasswordRequiredErr
	}

	// TEMP: hashed here because the request carries the password. In the invite
	// flow the hash is produced when the invitee sets it, not now.
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to hash super admin password: %w", err)
	}

	newEmployeeID := uuid.New()
	now := time.Now()

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		employee := &coreEmployeeModels.Employee{
			ID:             newEmployeeID,
			OrganizationID: orgID,
			FirstName:      firstName,
			LastName:       lastName,
			Email:          email,
			Status:         coreEmployeeConstants.EmployeeStatusActive,
			CreatedAt:      now,
			UpdatedAt:      now,
		}
		if err := s.employeeRepo.CreateEmployee(txCtx, employee); err != nil {
			return err
		}

		auth := &coreEmployeeModels.EmployeeAuth{
			ID:            uuid.New(),
			EmployeeID:    newEmployeeID,
			Username:      email,
			Email:         email,
			PasswordHash:  hashedPassword,
			IsActive:      true,
			IsProvisioned: true,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		if err := s.employeeRepo.CreateEmployeeAuth(txCtx, auth); err != nil {
			return err
		}

		onboarding := &coreEmployeeModels.EmployeeOnboarding{
			ID:         uuid.New(),
			EmployeeID: newEmployeeID,
			Status:     coreEmployeeConstants.EmployeeOnboardingActive,
			CreatedAt:  now,
		}
		if err := s.employeeRepo.CreateEmployeeOnBoarding(txCtx, onboarding); err != nil {
			return err
		}

		// Authorization: the employee is an owner only because of this edge —
		// the Owner role carries the full-access policy seeded at org signup.
		return s.sentinelService.AssignOwnerRole(txCtx, orgID, newEmployeeID)
	})
	if err != nil {
		return uuid.Nil, err
	}

	return newEmployeeID, nil
}
