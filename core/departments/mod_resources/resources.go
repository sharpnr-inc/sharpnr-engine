package orgDepartmentResources

import (
	agg "sharpnr-engine/sentinel_aggregators/contracts"
)

// Resources implements [agg.ResourceProvider].
//
// Department is ORGANIZATION-LEVEL: an org-level department hangs off the
// organization directly, so there is no HousedIn here — the engine stamps the
// Organization edge itself from the request's org. (The other department track,
// core.cost_center_departments, belongs to its cost center and would declare
// one; it is a separate type and not registered here.)
//
// The four generated CRUD actions are exactly what the module enforces, so
// there is nothing extra to declare. No attributes either: declaring one
// obliges this module to answer for it, and there is no condition worth writing
// over a department beyond scoping to specific ones — which the instance picker
// below already does.
func (s *orgDepartmentResources) Resources() []agg.ResourceRegistration {
	return []agg.ResourceRegistration{
		agg.Standard(ResourceTypeDepartment, "Departments").
			Category("Organization").
			Describe("Teams the organization is divided into").
			ListInstances(s.listDepartments).
			Build(),
	}
}
