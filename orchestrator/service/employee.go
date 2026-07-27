package orchestrator_service

import (
	"context"
	"fmt"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	sentinelservice "sharpnr-engine/sentinel/service"
	"strings"

	"github.com/google/uuid"
)

// employeeOrchestrator implements EmployeeOrchestrationInterface. The sentinel
// service is here because employee listings are decorated with each
// employee's role.
type employeeOrchestrator struct {
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface
	sentinelService   sentinelservice.SentinelServiceInterface

	// Wired for the upcoming owner-management flows; not consumed yet.
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

func (s *employeeOrchestrator) AddNewSuperAdmin(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	request coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto,
) error {

	// step 1 first we will check if these user is in other companies
	// step 2 then we create a temporary

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
