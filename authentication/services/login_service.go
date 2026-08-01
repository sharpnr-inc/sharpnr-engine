package authenticationEngineServices

import (
	"context"
	"fmt"
	"time"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	sentinelContexts "sharpnr-engine/sentinel_contexts"
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

	// The login route sits ahead of any auth middleware (there's no session
	// yet), so the client IP isn't stamped onto ctx by WebAuthMiddleware like
	// it is for protected routes. The controller stamps it directly before
	// calling LoginUser — see sentinelContexts.Inject in login.go.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	loginTime := time.Now().UTC()

	if err := s.auditEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: employee.OrganizationID,
		ActorID:        &employee.EmployeeID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         "authentication",
		Action:         "employee.login",
		ResourceType:   "employee",
		ResourceID:     &employee.EmployeeID,
		Description: fmt.Sprintf(
			"%s logged into the system on %s",
			employee.Name,
			loginTime.Format("2006-01-02 15:04:05 MST"),
		),
		Metadata: map[string]any{
			"email": req.Email,
			"role":  role.Name,
		},
	}); err != nil {
		return nil, err
	}

	resp := &authenticationEngineDtos.WebLoginResponse_dto{
		OrgId:      employee.OrganizationID,
		Email:      req.Email,
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Org:        authenticationEngineDtos.OrgData{Name: orgdetails.Name, SharpID: orgdetails.SharpID},
		Role:       role.Name,
	}

	return resp, nil
}
