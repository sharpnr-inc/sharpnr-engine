package orgCostCenterService

import (
	"context"
	"fmt"
	"strings"
	"time"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	orgCostCenterConstants "sharpnr-engine/core/cost_center/constants"
	orgCostCenterDtos "sharpnr-engine/core/cost_center/dtos"
	orgCostCenterModels "sharpnr-engine/core/cost_center/models"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// toView maps a stored cost center onto the wire shape the frontend reads.
//
// The timestamps are strings on the DTO, so they are rendered in RFC3339 —
// unambiguous about the offset, which matters for a module whose whole point is
// that sites live in different timezones.
func toView(costCenter orgCostCenterModels.OrgCostCenter) orgCostCenterDtos.CostCenterResponseDto {
	return orgCostCenterDtos.CostCenterResponseDto{
		ID:          costCenter.ID,
		Name:        costCenter.Name,
		Description: costCenter.Description,
		Code:        costCenter.Code,
		Address:     costCenter.AddressLine,
		City:        costCenter.City,
		State:       costCenter.State,
		Contry:      costCenter.Country,
		CreatedAt:   costCenter.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   costCenter.UpdatedAt.Format(time.RFC3339),
	}
}

/* ---------------------------------- <***> --------------------------------- */

// deriveCostCenterCode builds the short human key from the name ("Bangalore
// Office" -> "BANGALORE_OFFICE") for callers that do not supply one.
//
// Uniqueness is not this function's job: the caller checks it against the org,
// which is where a collision can actually be seen.
func deriveCostCenterCode(name string) string {
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
	if len(code) > orgCostCenterConstants.MaxCostCenterCodeLength {
		code = code[:orgCostCenterConstants.MaxCostCenterCodeLength]
	}
	return code
}

/* ---------------------------------- <***> --------------------------------- */

// recordAudit stamps one cost center event onto the org's history. It is called
// inside the same transaction as the write it describes, so the two commit or
// roll back together.
func (c *orgCostCenterServiceConstructor) recordAudit(
	ctx context.Context,
	orgID, emplID, costCenterID uuid.UUID,
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
		Module:         orgCostCenterConstants.AuditModule,
		Action:         action,
		ResourceType:   orgCostCenterConstants.AuditResourceType,
		ResourceID:     &costCenterID,
		Description:    description,
		Metadata:       metadata,
	})
}

/* ---------------------------------- <***> --------------------------------- */

func describeCostCenter(verb string, costCenter orgCostCenterModels.OrgCostCenter) string {
	return fmt.Sprintf("Cost center %s (%s) was %s", costCenter.Name, costCenter.Code, verb)
}
