package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel/models"
)

// These DTOs feed the frontend's Cedar policy builder — the "principal" palette.
//
// Policy applicability is ATTACHMENT-driven: the engine evaluates, for each
// user, the union of org-scoped policies and the policies attached to them
// (directly, via roles, via groups). Selecting a Role/Group in the builder
// therefore picks the ATTACHMENT TARGET for the policy — it does not write a
// `principal in Role::"..."` condition into the Cedar text (the engine builds
// no Role/Group entities, so such text would never match). Policy text targets
// the generic `principal`.

// PrincipalType is one selectable Cedar entity type the builder can scope a
// principal to, plus the operators valid for it.
type PrincipalType struct {
	Type      string   `json:"type"`      // Cedar entity type, e.g. "User", "Role", "Group"
	Label     string   `json:"label"`     // UI label
	Operators []string `json:"operators"` // valid Cedar operators, e.g. ["is","=="] or ["in"]
}

// PrincipalRef is one selectable principal instance. Value is the Cedar entity
// id baked into the policy text (slug for Role/Group), Label is the display name.
type PrincipalRef struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// PrincipalsView is the full principal palette for an org: the static entity
// types plus the org's live roles and groups.
type PrincipalsView struct {
	Types  []PrincipalType `json:"types"`
	Roles  []PrincipalRef  `json:"roles"`
	Groups []PrincipalRef  `json:"groups"`
}

// RoleRef projects a role to the ref the builder shows as an attachment target.
// Name is unique per org.
func RoleRef(r *sentinelengine_models.Role) PrincipalRef {
	return PrincipalRef{Value: r.Name, Label: r.Name}
}

// GroupRef projects a group to the ref the builder shows as an attachment
// target. Name identifies a group within an org.
func GroupRef(g *sentinelengine_models.Group) PrincipalRef {
	return PrincipalRef{Value: g.Name, Label: g.Name}
}
