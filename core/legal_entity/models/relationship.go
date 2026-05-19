package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
)

// LegalEntityRelationship models parent-child and cross-entity relationships.
// A holding company owns subsidiaries; a subsidiary may have branches.
// RelationshipType also covers JV partners and franchise arrangements.
type LegalEntityRelationship struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	ParentEntityID uuid.UUID `gorm:"type:uuid;not null;index"`
	ChildEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	// HOLDING | SUBSIDIARY | BRANCH | JV_PARTNER | FRANCHISE
	RelationshipType string `gorm:"size:20;not null"`

	// OwnershipPercent is relevant for HOLDING/SUBSIDIARY — used in
	// consolidated P&L and transfer pricing documentation.
	OwnershipPercent *float64

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityRelationship) TableName() string {
	return "legal_enitity.legal_entity_relationships"
}
