package sentinelservice

import (
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// These DTOs are the shapes the React app consumes. They deliberately expose
// the structured authoring layer (managed policies, attachments) and never the
// raw Cedar text — privileges are shown as catalog-labeled permissions.

// ManagedPolicyView is a managed policy with its actions expanded into
// human-readable permissions for display.
type ManagedPolicyView struct {
	ID          uuid.UUID                     `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Effect      string                        `json:"effect"`     // permit | forbid
	System      bool                          `json:"system"`     // platform-provided (read-only in UI)
	AllActions  bool                          `json:"allActions"` // true => grants every permission (Permissions is empty; render "All permissions")
	Permissions []sentinel_actions.Permission `json:"permissions"`
}

// RoleView is the lightweight role shape for dropdowns / lists.
type RoleView struct {
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// RolePrivilegesView answers "what can this role do?" for the role detail page:
// the role, the managed policies attached to it, and the deduped union of the
// permissions those policies grant.
type RolePrivilegesView struct {
	Role                 RoleView                      `json:"role"`
	Policies             []ManagedPolicyView           `json:"policies"`
	EffectivePermissions []sentinel_actions.Permission `json:"effectivePermissions"`
}

// CreateManagedPolicyInput is the request to author a new managed policy.
type CreateManagedPolicyInput struct {
	Name         string   `json:"name" binding:"required"`
	Description  string   `json:"description"`
	Actions      []string `json:"actions" binding:"required,min=1"` // catalog action keys
	Effect       string   `json:"effect"`                           // permit (default) | forbid
	ResourceType string   `json:"resourceType"`
}

// CreateRoleInput is the request to create a role and optionally attach policies.
type CreateRoleInput struct {
	Name              string      `json:"name" binding:"required"`
	Description       string      `json:"description"`
	ManagedPolicyIDs  []uuid.UUID `json:"managedPolicyIds"` // optional initial attachments
}

// AttachPoliciesInput is the request to attach managed policies to a role.
type AttachPoliciesInput struct {
	ManagedPolicyIDs []uuid.UUID `json:"managedPolicyIds" binding:"required,min=1"`
}

// expandActions maps stored action keys to catalog permissions for display.
// Unknown keys (e.g. an action removed from the catalog but still stored) are
// surfaced with the key as the label so they remain visible, not silently dropped.
func expandActions(actions []string) []sentinel_actions.Permission {
	out := make([]sentinel_actions.Permission, 0, len(actions))
	for _, a := range actions {
		key := sentinel_actions.Name(a)
		if p, ok := sentinel_actions.Lookup(key); ok {
			out = append(out, p)
			continue
		}
		out = append(out, sentinel_actions.Permission{Key: key, Label: a})
	}
	return out
}

// toManagedPolicyView expands a stored managed policy for the UI.
func toManagedPolicyView(mp *sentinelengine_models.ManagedPolicy) ManagedPolicyView {
	return ManagedPolicyView{
		ID:          mp.ID,
		Name:        mp.Name,
		Description: mp.Description,
		Effect:      string(mp.Effect),
		System:      mp.System,
		AllActions:  mp.AllActions,
		Permissions: expandActions(mp.Actions), // empty when AllActions; UI shows "All permissions"
	}
}

// toRoleView projects a role model to its list shape.
func toRoleView(r *sentinelengine_models.Role) RoleView {
	return RoleView{ID: r.ID, Slug: r.Slug, Name: r.Name, Description: r.Description}
}
