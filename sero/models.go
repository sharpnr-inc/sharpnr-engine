package sero

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeroOutbox struct {
	ID        uuid.UUID           `gorm:"primaryKey"`
	EventName SeroEngineEventType `gorm:"index"`
	Payload   []byte              `gorm:"type:jsonb"`
	Status    string              `gorm:"default:'PENDING';index"`
	Retries   int                 `gorm:"default:0"`
	LastError string              `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SeroOutbox) TableName() string {
	return "sero_engine.sero_outbox"
}
