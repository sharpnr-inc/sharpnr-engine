package folioEngineService

import (
	"context"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	folioEngineConstants "sharpnr-engine/folio_engine/constants"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// recordAudit stamps one Folio event onto the org's history. It is called
// inside the same transaction as the write it describes, so the two commit or
// roll back together.
//
// The resource type is fixed for now — the account is the only thing this
// module writes. It becomes a parameter, the way cables does it, once documents
// and boards land.
func (f *folioService) recordAudit(
	ctx context.Context,
	orgID, emplID, resourceID uuid.UUID,
	action, description string,
) error {
	if f.auditLogEngine == nil {
		return nil
	}

	// Protected routes have the client IP stamped onto ctx by the auth
	// middleware; when nothing injected it the field stays empty rather than
	// failing the write.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	return f.auditLogEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &emplID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         folioEngineConstants.AuditModule,
		Action:         action,
		ResourceType:   folioEngineConstants.AuditResourceTypeFolioAccount,
		ResourceID:     &resourceID,
		Description:    description,
	})
}
