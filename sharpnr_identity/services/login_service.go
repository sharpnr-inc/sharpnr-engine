package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_dtos "sharpnr-engine/sharpnr_identity/dtos"
)

// LoginUser implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) LoginUser(
	ctx context.Context,
	req sharpnr_identity_auth_dtos.LoginRequest_dto,
) (*sharpnr_identity_auth_dtos.WebLoginResponse_dto, error) {

	employee, err := s.employeeService.AuthenticateEmployee(
		ctx,
		req.Email,
		req.Password,
	)
	if err != nil {
		// propagate auth error (already sanitized)
		return nil, err
	}

	orgdetails, err := s.orgService.GetOrgDetailsByID(ctx, employee.OrganizationID, employee.EmployeeID)
	if err != nil {
		return nil, err
	}

	role, err := s.sentinelService.GetEmployeeRole(ctx, employee.OrganizationID, employee.EmployeeID)
	if err != nil {
		return nil, err
	}

	resp := &sharpnr_identity_auth_dtos.WebLoginResponse_dto{
		OrgId:      employee.OrganizationID,
		Email:      employee.Name,
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Org:        sharpnr_identity_auth_dtos.OrgData{Name: orgdetails.Name, SharpID: orgdetails.SharpID},
		Role:       role.Name,
	}

	return resp, nil
}
