package orchestrator_service

import (
	"context"
	"fmt"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeErrs "sharpnr-engine/core/employees/errs"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"strings"

	"github.com/google/uuid"
)

// employeeOrchestrator implements EmployeeOrchestrationInterface. The sentinel
// service is here because employee listings are decorated with each
// employee's role.
type employeeOrchestrator struct {
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface
	sentinelService   sentinelservice.SentinelServiceInterface

	// Owner/super-admin creation — the only flow the orchestrator-facing
	// employee service exposes.
	employeeService coreEmployeeServices.EmployeeOrchestratorServiceInterface
}

// Fetches all the employees of the organization
func (s *employeeOrchestrator) GetAllEmployeesOfOrg(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]orchestrator_dtos.EmployeeOfOrganization, error) {

	// Fetch all the employees in the organization
	empList, err := s.employeeHRService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}

	// Pre-allocate the slice to optimize memory and performance
	result := make([]orchestrator_dtos.EmployeeOfOrganization, 0, len(empList))

	for _, emp := range empList {
		// Safety check to avoid nil pointer dereference
		if emp == nil {
			continue
		}

		// Fetch the role view for this specific employee
		roleView, err := s.sentinelService.GetEmployeeRole(ctx, orgID, emp.ID)
		if err != nil {
			return nil, err
		}

		// Map the fields from OrgEmployeeList and RoleView to EmployeeOfOrganization
		mappedEmp := orchestrator_dtos.EmployeeOfOrganization{
			ID:     emp.ID,
			Name:   emp.EmployeeName,
			Status: emp.Status,
			Role: orchestrator_dtos.EmployeeRole{
				ID:          roleView.ID,
				Role:        roleView.Name, // Maps the Name from RoleView to the Role field
				Description: roleView.Description,
			},
		}

		result = append(result, mappedEmp)
	}

	return result, nil
}

/* -------------------------------------------------------------------------- */
/*                            Add a New Super Admin                           */
/* -------------------------------------------------------------------------- */

// AddNewSuperAdmin adds a second owner/super-admin to the caller's org.
//
// TEMPORARY: the account is created straight away from a password supplied in
// the request — the invitation flow this endpoint is meant to run (and whose
// wording the controller still uses) is not built yet. See the note on
// coreEmployeeServices.AddNewOwnerOrSuperAdmin for what changes when it lands.
func (s *employeeOrchestrator) AddNewSuperAdmin(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	request coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto,
) error {

	// Authorization (rule 1 in doc.go): employees are not a Cedar resource type
	// yet, so there is nothing for guard.Authorize to evaluate. Minting an owner
	// hands out full access to the whole org, so it is gated on the caller
	// already being one — the same rule the interface comment in
	// core/employees/services states.
	callerRole, err := s.sentinelService.GetEmployeeRole(ctx, orgID, emplID)
	if err != nil {
		return err
	}
	if !strings.EqualFold(callerRole.Name, sentinelservice.OwnerRoleName) {
		return fmt.Errorf("%w: %s", sentinelEngine.ErrAccessDenied, coreEmployeeErrs.NotAnOwnerErr)
	}

	// The email is the login identity and is unique across the whole system, so
	// a duplicate is rejected here rather than surfacing as a constraint
	// violation from the insert. Normalized first so the check and the row that
	// follows agree on what the address is.
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))

	exists, err := s.employeeHRService.IsUserExists(ctx, request.Email)
	if err != nil {
		return fmt.Errorf("failed to check user existence for email %s: %w", request.Email, err)
	}
	if exists {
		return fmt.Errorf("%w: %s", coreEmployeeErrs.EmployeeAlreadyExistsErr, request.Email)
	}

	// Single delegated write (rule 2 in doc.go): the employee service owns the
	// transaction that spans the employee, auth, onboarding and owner-role rows.
	if _, err := s.employeeService.AddNewOwnerOrSuperAdmin(ctx, orgID, emplID, request); err != nil {
		return err
	}

	return nil
}

/* -------------------------------------------------------------------------- */
/*                              Employee Details                              */
/* -------------------------------------------------------------------------- */

func (c *employeeOrchestrator) GetEmployeeDetails(ctx context.Context, orgID, emplID uuid.UUID, requestEmplID uuid.UUID) (orchestrator_dtos.EmployeeDetailsResponseDto, error) {
	panic("")
}

/* ---------------------------------- <***> --------------------------------- */

func (s *employeeOrchestrator) AddEmployeeInBulk(ctx context.Context, orgID, emplID uuid.UUID, request *[]orchestrator_dtos.AddEmployeeBulkRequestDto) (*orchestrator_dtos.BulkProcessResult, error) {

	// 1. Check slice length
	if request == nil || len(*request) == 0 {
		return &orchestrator_dtos.BulkProcessResult{}, nil
	}

	// 2. Extract unique, normalized emails from the payload
	emails := make([]string, 0, len(*request))
	emailToDtoMap := make(map[string]orchestrator_dtos.AddEmployeeBulkRequestDto, len(*request))

	for _, req := range *request {
		normalizedEmail := strings.ToLower(strings.TrimSpace(req.EmailAddress))
		if normalizedEmail == "" {
			continue
		}

		if _, exists := emailToDtoMap[normalizedEmail]; !exists {
			emails = append(emails, normalizedEmail)
			emailToDtoMap[normalizedEmail] = req
		}
	}

	if len(emails) == 0 {
		return &orchestrator_dtos.BulkProcessResult{}, nil
	}

	// 3. Check existing users
	existingEmailsMap := make(map[string]bool, len(emails))

	for _, email := range emails {
		exists, err := s.employeeHRService.IsUserExists(ctx, email)
		if err != nil {
			return nil, fmt.Errorf("failed to check user existence for email %s: %w", email, err)
		}

		if exists {
			existingEmailsMap[email] = true
		}
	}

	// 4. Separate valid records from skipped ones
	var validRequests []orchestrator_dtos.AddEmployeeBulkRequestDto
	var skippedEmails []string

	for email, dto := range emailToDtoMap {
		if existingEmailsMap[email] {
			skippedEmails = append(skippedEmails, email)
			continue
		}
		validRequests = append(validRequests, dto)
	}

	// 5. If no valid records remain, return early
	if len(validRequests) == 0 {
		return &orchestrator_dtos.BulkProcessResult{
			CreatedCount:  0,
			SkippedCount:  len(skippedEmails),
			SkippedEmails: skippedEmails,
		}, nil
	}

	// 6. Map to Core HR DTOs and insert into DB
	hrCoreRequests := mapToCoreHRBulkDto(validRequests)

	if err := s.employeeHRService.AddEmployeeInBulkInOrg(ctx, orgID, emplID, hrCoreRequests); err != nil {
		return nil, fmt.Errorf("failed to add bulk employees in organization: %w", err)
	}

	return &orchestrator_dtos.BulkProcessResult{
		CreatedCount:  len(validRequests),
		SkippedCount:  len(skippedEmails),
		SkippedEmails: skippedEmails,
	}, nil
}

func mapToCoreHRBulkDto(reqs []orchestrator_dtos.AddEmployeeBulkRequestDto) []coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto {
	coreReqs := make([]coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto, 0, len(reqs))

	for _, r := range reqs {
		// Handle conversion of string to *string for LastName
		var lastNamePtr *string
		if strings.TrimSpace(r.LastName) != "" {
			lastName := r.LastName
			lastNamePtr = &lastName
		}

		coreReqs = append(coreReqs, coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto{
			FirstName:    r.FirstName,
			MiddleName:   r.MiddleName,
			LastName:     lastNamePtr,
			Email:        r.EmailAddress, // mapping emailAddress -> email
			PhoneNumber:  r.PhoneNumber,
			Nickname:     r.NickName, // mapping nickName -> nickname
			EmployeeCode: r.EmployeeCode,
		})
	}

	return coreReqs
}

/* ---------------------------------- <***> --------------------------------- */
