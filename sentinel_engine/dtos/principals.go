package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
)

// These DTOs feed the frontend's Cedar policy builder — the "principal" palette.
// A policy's principal scope is written one of four ways, and each draws on a
// different part of this view:
//
//	principal is User                    -> a PrincipalType with operator "is"
//	principal == User::"<employee-uuid>" -> (deferred) a specific user
//	principal in Role::"<slug>"          -> a PrincipalRef from Roles
//	principal in Group::"<slug>"         -> a PrincipalRef from Groups
//
// Role/Group are embedded in the Cedar text by SLUG (Value), while the UI shows
// the human Label.

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

// RoleRef projects a role to the Cedar-slug-keyed ref the builder embeds.
func RoleRef(r *sentinelengine_models.Role) PrincipalRef {
	return PrincipalRef{Value: r.Slug, Label: r.Name}
}

// GroupRef projects a group to the Cedar-slug-keyed ref the builder embeds.
func GroupRef(g *sentinelengine_models.Group) PrincipalRef {
	return PrincipalRef{Value: g.Slug, Label: g.Name}
}
