package orgDepartmentService

import (
	"context"
	"fmt"
	"strings"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	orgDepartmentsConstants "sharpnr-engine/core/departments/constants"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentModels "sharpnr-engine/core/departments/models"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// toView maps a stored department onto the wire shape the frontend reads.
func toView(department orgDepartmentModels.OrgDeparment) orgDepartmentsDtos.DepartmentViewResponseDto {
	view := orgDepartmentsDtos.DepartmentViewResponseDto{
		ID:          department.ID,
		Name:        department.Name,
		Description: department.Description,
		Code:        department.Code,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
	}

	if department.HeadEmployeeID != nil {
		head := department.HeadEmployeeID.String()
		view.DepartmentHead = &head
	}

	return view
}

/* ---------------------------------- <***> --------------------------------- */

// deriveDepartmentCode builds the short human key from the department name
// ("Front Office" -> "FRONT_OFFICE") for callers that do not supply one.
//
// Uniqueness is not this function's job: the caller checks it against the org,
// which is where a collision can actually be seen.
func deriveDepartmentCode(name string) string {
	var b strings.Builder

	for _, r := range strings.ToUpper(strings.TrimSpace(name)) {
		switch {
		case r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-' || r == '_':
			// Collapse runs of separators into a single underscore.
			if code := b.String(); code != "" && !strings.HasSuffix(code, "_") {
				b.WriteRune('_')
			}
		}
	}

	code := strings.Trim(b.String(), "_")
	if len(code) > orgDepartmentsConstants.MaxDepartmentCodeLength {
		code = code[:orgDepartmentsConstants.MaxDepartmentCodeLength]
	}
	return code
}

/* ---------------------------------- <***> --------------------------------- */

// recordAudit stamps one department event onto the org's history. It is called
// inside the same transaction as the write it describes, so the two commit or
// roll back together.
func (c *coreOrgDepartmentService) recordAudit(
	ctx context.Context,
	orgID, emplID, departmentID uuid.UUID,
	action, description string,
	metadata map[string]any,
) error {
	if c.auditEngine == nil {
		return nil
	}

	// Protected routes have the client IP stamped onto ctx by the auth
	// middleware; when nothing injected it the field stays empty rather than
	// failing the write.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	return c.auditEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &emplID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         orgDepartmentsConstants.AuditModule,
		Action:         action,
		ResourceType:   orgDepartmentsConstants.AuditResourceType,
		ResourceID:     &departmentID,
		Description:    description,
		Metadata:       metadata,
	})
}

/* ---------------------------------- <***> --------------------------------- */

// changedFields is the metadata payload for an update: what the caller
// actually asked to change, so the audit trail says more than "it changed".
func changedFields(req orgDepartmentsDtos.UpdateDepartmentRequestDto) map[string]any {
	changes := map[string]any{}

	if req.Name != nil {
		changes["name"] = *req.Name
	}
	if req.Description != nil {
		changes["description"] = *req.Description
	}
	if req.Code != nil {
		changes["code"] = *req.Code
	}
	if req.DepartmentHead != nil {
		changes["department_head"] = *req.DepartmentHead
	}

	return changes
}

/* ---------------------------------- <***> --------------------------------- */

func describeDepartment(verb string, department orgDepartmentModels.OrgDeparment) string {
	return fmt.Sprintf("Department %s (%s) was %s", department.Name, department.Code, verb)
}
