package auditEngineRepo

import (
	"context"
	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *auditLogRepo) CreateLog(
	ctx context.Context,
	log *auditEngineModels.AuditLog,
) error {
	// If the caller is inside a transaction the log entry joins it,
	// so a rolled-back business action never leaves a phantom audit entry.
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(log).Error
}

func (r *auditLogRepo) GetLogsByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
	query auditEngineDtos.ListLogsQuery,
) ([]auditEngineModels.AuditLog, int64, error) {
	var logs []auditEngineModels.AuditLog
	var total int64

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	base := db.WithContext(ctx).
		Model(&auditEngineModels.AuditLog{}).
		Where("organization_id = ?", orgID)

	base = applyFilters(base, query)

	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := base.
		Order("created_at DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&logs).
		Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

func (r *auditLogRepo) GetLogsByResource(
	ctx context.Context,
	orgID uuid.UUID,
	resourceType string,
	resourceID uuid.UUID,
) ([]auditEngineModels.AuditLog, error) {
	var logs []auditEngineModels.AuditLog

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("organization_id = ? AND resource_type = ? AND resource_id = ?",
			orgID, resourceType, resourceID).
		Order("created_at DESC").
		Find(&logs).
		Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func applyFilters(db *gorm.DB, query auditEngineDtos.ListLogsQuery) *gorm.DB {
	if query.Module != "" {
		db = db.Where("module = ?", query.Module)
	}
	if query.Action != "" {
		db = db.Where("action = ?", query.Action)
	}
	if query.ResourceType != "" {
		db = db.Where("resource_type = ?", query.ResourceType)
	}
	if query.ResourceID != nil {
		db = db.Where("resource_id = ?", *query.ResourceID)
	}
	if query.ActorID != nil {
		db = db.Where("actor_id = ?", *query.ActorID)
	}
	return db
}
