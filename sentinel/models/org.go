package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrgSentinelPolicy struct {
	ID               uuid.UUID
	OrganizationID   uuid.UUID
	SentinelPolicyID uuid.UUID // this is ID of Sentinel Policy

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (OrgSentinelPolicy) TableName() string {
	return "sentinel_engine.org_sentinel_policies"
}
