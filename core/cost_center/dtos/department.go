package orgCostCenterDtos

// CreateCostCenterDepartmentRequestDto is the payload for a new department
// inside one cost center — the Bangalore office's Finance team.
//
// The cost center it belongs to comes from the request path, not from here: a
// department is created under a site, and a body that could name a different
// one would be a second, contradictory source of truth.
type CreateCostCenterDepartmentRequestDto struct {
	Name        string `json:"name" validate:"required" binding:"required"`
	Description string `json:"description,omitempty"`

	// Code is the short human key ("FIN"), unique within the cost center. It is
	// optional on the wire: when it is empty the service derives one from the
	// name, because the column is NOT NULL.
	Code string `json:"code,omitempty"`

	// DepartmentHead is the employee who runs this department at this location,
	// as a UUID string. Empty means nobody is in charge yet.
	DepartmentHead string `json:"department_head,omitempty"`
}
