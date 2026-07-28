package coreEmployeeServices

import (
	"context"
	"fmt"
	"net/url"
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

	provisions := make([]coreEmployeeModels.EmployeeProvisioning, 0, len(req))

	for _, emp := range req {
		lastName := ""
		if emp.LastName != nil {
			lastName = *emp.LastName
		}

		// 1. Create Provisioning Token
		token, err := s.jwtManager.IssueEmployeeProvisioningToken(emp.FirstName, lastName, emp.Email)
		if err != nil {
			return fmt.Errorf("failed to issue provisioning token for %s: %w", emp.Email, err)
		}

		// 2. Prepare Provisioning Model Record
		// (Assuming token is stored directly in TokenHash or hashed if your architecture requires it)
		provision := coreEmployeeModels.EmployeeProvisioning{
			ID:         uuid.New(),
			EmployeeID: emplID,
			Email:      emp.Email,
			TokenHash:  token,
			ExpiresAt:  time.Now().Add(24 * 7 * time.Hour), // Set standard expiration duration (e.g., 7 days)
			CreatedAt:  time.Now(),
		}
		provisions = append(provisions, provision)

		// 3. Construct URL
		actionURL := fmt.Sprintf(
			"https://auth.sharpnr.com/emp/prov/%s+%s+%s+%s",
			url.QueryEscape(token),
			url.QueryEscape(emp.FirstName),
			url.QueryEscape(lastName),
			url.QueryEscape(emp.Email),
		)
		// 4. Enqueue Email Notification
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

	// 5. Save all provisions in bulk
	if err := s.employeeRepo.CreateEmployeeProvision(ctx, orgID, provisions); err != nil {
		return fmt.Errorf("failed to persist employee provisions: %w", err)
	}

	return nil
}
