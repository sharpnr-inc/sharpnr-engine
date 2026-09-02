package orchestrator_service

import orgDepartmentService "sharpnr-engine/core/departments/service"

type departmentOrchestrator struct {
	coreDepartmentService orgDepartmentService.OrgDepartmentServiceInterface
}
