package hr_employee_services

import (
	"context"
	sharphrdtos "sharpnr-engine/modules/hr/hr_dtos"
	"sharpnr-engine/shared/mailer"

	"github.com/google/uuid"
)

func (s *hrEmployeeService) CreateNewEmployee(
	ctx context.Context,
	orgID uuid.UUID,
	req sharphrdtos.CreateEmployeeInOrganization_RequestDTO,
) error {

	_, err := mailer.SendVerificationLinkToEmployee(req.Email, "thisistest")
	if err != nil {
		return err
	}
	return nil
}
