package hrmsEngineService

import (
	"context"
	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"

	"github.com/google/uuid"
)

// HrmsEngineServiceInterface is the HRMS engine's single service surface.
// Attendance, payroll and recruitment methods join it here as those flows land.
type HrmsEngineServiceInterface interface {
	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*hrmsEngineDtos.OrgEmployeeList, error)
}
