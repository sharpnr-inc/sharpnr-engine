package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID     uuid.UUID
	UserID    uuid.UUID
	Action    string
	Entity    string
	EntityID  uuid.UUID
	Timestamp time.Time
}
