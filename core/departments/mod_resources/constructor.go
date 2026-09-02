package orgDepartmentResources

import (
	orgDepartmentService "sharpnr-engine/core/departments/service"
	agg "sharpnr-engine/sentinel_aggregators/contracts"
)

type orgDepartmentResources struct {
	departmentService orgDepartmentService.OrgDepartmentServiceInterface
}

// Compile-time guard: the provider must satisfy the aggregator contract.
var _ agg.ResourceProvider = (*orgDepartmentResources)(nil)

func NewOrgDepartmentResources(departmentService orgDepartmentService.OrgDepartmentServiceInterface) *orgDepartmentResources {
	return &orgDepartmentResources{departmentService: departmentService}
}
