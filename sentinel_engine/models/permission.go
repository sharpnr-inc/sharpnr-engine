package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
)

// Permission is ONE row of the access grid in the UI: "for this role, on this
// resource type, at this property, here is what's allowed (view/create/edit/
// delete)." One visible grid row == one Permission record.
//
// Scope rules:
//   - PropertyID set   -> applies only to resources at that property.
//   - PropertyID nil   -> applies to that resource type across ALL properties
//     in the org (used by the org-wide Super Admin).
//   - ResourceType "*" -> applies to every resource type (also Super Admin).
type Permission struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID uuid.UUID `gorm:"type:uuid;not null;index"`

	// PropertyID is the per-property scope. Nullable: nil means "all properties".
	PropertyID *uuid.UUID `gorm:"type:uuid;index"`

	// ResourceType is a business module like "Hotel", "Restaurant", "Bar",
	// "Kitchen", or "*" for every type.
	ResourceType string `gorm:"type:varchar(50);not null"`

	// The CRUD checkboxes from the grid.
	CanView   bool `gorm:"not null;default:false"`
	CanCreate bool `gorm:"not null;default:false"`
	CanEdit   bool `gorm:"not null;default:false"`
	CanDelete bool `gorm:"not null;default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Permission) TableName() string {
	return "policy_engine.permission"
}

// Action names — these are the strings Cedar matches on (Action::"view", ...)
// and the same values the calling service passes as EvaluationRequest.Action.
const (
	ActionView   = "view"
	ActionCreate = "create"
	ActionEdit   = "edit"
	ActionDelete = "delete"
)

// ToPolicy converts this permission row into a Cedar policy for the given role.
//
// It returns ok=false when the row grants nothing (all checkboxes off), so the
// caller can skip it. When all four boxes are on, the action list is left empty,
// which ToCedar renders as "all actions" — convenient and future-proof for the
// Super Admin.
func (p Permission) ToPolicy(roleID uuid.UUID) (Policy, bool) {
	var actions []string
	if p.CanView {
		actions = append(actions, ActionView)
	}
	if p.CanCreate {
		actions = append(actions, ActionCreate)
	}
	if p.CanEdit {
		actions = append(actions, ActionEdit)
	}
	if p.CanDelete {
		actions = append(actions, ActionDelete)
	}
	if len(actions) == 0 {
		return Policy{}, false // nothing granted
	}
	if len(actions) == 4 {
		actions = nil // all CRUD -> unconstrained action (covers future actions too)
	}

	policy := Policy{
		ID:            "perm-" + p.ID.String(),
		Effect:        Permit,
		PrincipalType: PrincipalRole,
		PrincipalID:   roleID.String(),
		Actions:       actions,
		ResourceType:  p.ResourceType, // "*" -> any type
	}

	// Per-property scope: the resource must live under this Property.
	if p.PropertyID != nil {
		policy.ResourceInType = "Property"
		policy.ResourceInID = p.PropertyID.String()
	}

	return policy, true
}
