package cortexEngineService

import (
	"context"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	cortexEngineConstants "sharpnr-engine/cortex_engine/constants"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// recordAudit stamps one Cortex event onto the org's history. It is called
// inside the same transaction as the write it describes, so the two commit or
// roll back together.
//
// The resource type is fixed for now — the account is the only thing this
// module writes. It becomes a parameter, the way cables does it, once roles and
// policies land.
func (c *coretexService) recordAudit(
	ctx context.Context,
	orgID, emplID, resourceID uuid.UUID,
	action, description string,
) error {
	if c.auditLogEngine == nil {
		return nil
	}

	// Protected routes have the client IP stamped onto ctx by the auth
	// middleware; when nothing injected it the field stays empty rather than
	// failing the write.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	return c.auditLogEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &emplID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         cortexEngineConstants.AuditModule,
		Action:         action,
		ResourceType:   cortexEngineConstants.AuditResourceTypeCortexAccount,
		ResourceID:     &resourceID,
		Description:    description,
	})
}
