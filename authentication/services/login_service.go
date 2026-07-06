package authenticationEngineServices

import (
	"context"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
)

// LoginUser implements AuthenticationServiceInterface.
func (s *authenticationService) LoginUser(
	ctx context.Context,
	req authenticationEngineDtos.LoginRequest_dto,
) (*authenticationEngineDtos.WebLoginResponse_dto, error) {

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

	resp := &authenticationEngineDtos.WebLoginResponse_dto{
		OrgId:      employee.OrganizationID,
		Email:      employee.Name,
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Org:        authenticationEngineDtos.OrgData{Name: orgdetails.Name, SharpID: orgdetails.SharpID},
		Role:       role.Name,
	}

	return resp, nil
}
