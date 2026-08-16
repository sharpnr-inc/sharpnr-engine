package coreEmployeeServices

import (
	"context"
	"fmt"
	"net/url"
	coreEmployeeConstants "sharpnr-engine/core/employees/constants"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeModels "sharpnr-engine/core/employees/models"
	"sharpnr-engine/sharp_mailer_engine"
	"time"

	"github.com/google/uuid"
)

func (s *employeeService) AddEmployeeInBulkInOrg(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req []coreEmployeeContracts.AddEmployeeInBulkInOrganizationRequestDto) error {
	if len(req) == 0 {
		return nil
	}

	employees := make([]coreEmployeeModels.Employee, 0, len(req))
	provisions := make([]coreEmployeeModels.EmployeeProvisioning, 0, len(req))

	for _, emp := range req {
		lastName := ""
		if emp.LastName != nil {
			lastName = *emp.LastName
		}

		middleName := ""
		if emp.MiddleName != nil {
			middleName = *emp.MiddleName
		}

		nickname := ""
		if emp.Nickname != nil {
			nickname = *emp.Nickname
		}

		employeeCode := ""
		if emp.EmployeeCode != nil {
			employeeCode = *emp.EmployeeCode
		}

		// 1. Prepare New Employee Record (Status: PENDING)
		employeeID := uuid.New()
		newEmployee := coreEmployeeModels.Employee{
			ID:             employeeID,
			OrganizationID: orgID,
			FirstName:      emp.FirstName,
			MiddleName:     middleName,
			LastName:       lastName,
			NickName:       nickname,
			Email:          emp.Email,
			PhoneNumber:    emp.PhoneNumber,
			Status:         coreEmployeeConstants.EmployeeStatusPending,
			EmployeeCode:   employeeCode,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		employees = append(employees, newEmployee)

		// 2. Issue Provisioning Token
		token, err := s.jwtManager.IssueEmployeeProvisioningToken(emp.FirstName, lastName, emp.Email)
		if err != nil {
			return fmt.Errorf("failed to issue provisioning token for %s: %w", emp.Email, err)
		}

		// 3. Prepare Provisioning Record Linked to the newly generated Employee ID
		provision := coreEmployeeModels.EmployeeProvisioning{
			ID:         uuid.New(),
			EmployeeID: employeeID,
			Email:      emp.Email,
			TokenHash:  token,
			ExpiresAt:  time.Now().Add(24 * 7 * time.Hour), // 7 days expiration
			CreatedAt:  time.Now(),
		}
		provisions = append(provisions, provision)

		// 4. Construct Safely Encoded URL
		actionURL := fmt.Sprintf(
			"https://auth.sharpnr.com/emp/prov/%s+%s+%s+%s",
			url.QueryEscape(token),
			url.QueryEscape(emp.FirstName),
			url.QueryEscape(lastName),
			url.QueryEscape(emp.Email),
		)

		// 5. Enqueue Notification Job
		s.mailerEngine.Enqueue(sharp_mailer_engine.NewNotificationJob(
			emp.Email,
			"Set Up Your Employee Account",
			sharp_mailer_engine.NotificationData{
				Title:      "Welcome to the Organization",
				Subtitle:   fmt.Sprintf("Hello %s,", emp.FirstName),
				Message:    "You have been added to the organization. Please click the button below to set up your account password.",
				ActionURL:  actionURL,
				ActionText: "Set Up Account",
				FooterNote: "This link will expire in 7 days.",
			},       
		))
	}
                                        
                         
	return nil
}
