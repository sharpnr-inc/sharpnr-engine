package auditEngineRepo

import "gorm.io/gorm"

type auditLogRepo struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *auditLogRepo {
	return &auditLogRepo{
		db: db,
	}
}
