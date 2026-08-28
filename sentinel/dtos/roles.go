package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// These DTOs are the shapes the React app consumes for role management.

// CreateRoleInput is the create-role payload. PolicyIDs is an optional set of
// existing policies to attach to the role at creation time.
type CreateRoleInput struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	PolicyIDs   *[]uuid.UUID `json:"policyIds,omitempty"`
}

// RoleView is the lightweight role shape for dropdowns / lists.
type RoleView struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// RolePrivilegesView is a role plus the policies attached to it — what the UI
// shows when an admin inspects a role.
type RolePrivilegesView struct {
	Role     RoleView     `json:"role"`
	Policies []PolicyView `json:"policies"`
}

// AttachPoliciesInput is the payload to group existing policies under a role.
type AttachPoliciesInput struct {
	PolicyIDs []uuid.UUID `json:"policyIds"`
}

// ToRoleView projects a role model to its list shape.
func ToRoleView(r *sentinelengine_models.Role) RoleView {
	return RoleView{ID: r.ID, Name: r.Name, Description: r.Description}
}

// ToRoleViews projects a slice of role models to list shapes.
func ToRoleViews(rs []*sentinelengine_models.Role) []RoleView {
	out := make([]RoleView, 0, len(rs))
	for _, r := range rs {
		out = append(out, ToRoleView(r))
	}
	return out
}

type UpdateRoleDetailsRequestDto struct {
	Name        *string
	Description *string
}
