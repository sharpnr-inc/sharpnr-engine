package hrmsEngineService

import (
	"context"
	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"

	"github.com/google/uuid"
)

// HrmsEngineServiceInterface is the HRMS engine's single service surface.
// Attendance, payroll and recruitment methods join it here as those flows land.
type HrmsEngineServiceInterface interface {
	// StartHrmsServiceForOrg provisions the org's one HRMS account together with
	// its settings row.
	StartHrmsServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req hrmsEngineDtos.StartHrmsServiceRequestDto) error

	// GetHrmsAccountStatus is the read every HRMS-owned flow is gated on: has
	// this org started HRMS, and is that account switched on.
	GetHrmsAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (hrmsEngineDtos.HrmsAccountStatusResponse, error)

	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*hrmsEngineDtos.OrgEmployeeList, error)
}
