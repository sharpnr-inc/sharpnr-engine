package sentinelDtos

import (
	sentinelengine_models "sharpnr-engine/sentinel/models"
)

// These DTOs feed the frontend's Cedar policy builder — the "principal" palette.
//
// The palette answers two DIFFERENT questions, and keeping them apart is the
// point of this file:
//
//  1. Types — what may appear in the policy's PRINCIPAL CLAUSE. Only User: it
//     is the only entity the engine builds.
//
//  2. Roles / Groups / Users — what the policy may be ATTACHED TO. Picking one
//     writes nothing into the Cedar text; it chooses which register the policy
//     is recorded in, and that is what decides who it reaches.
//
// Policy applicability is attachment-driven: the engine evaluates, per user,
// the union of the org-wide policies and those registered against them
// directly, via a role, or via a group. So policy text normally targets the
// generic `principal` and lets attachment decide applicability — which is
// exactly what the seeded Owner policy does.

// PrincipalType is one selectable Cedar entity type the builder can scope a
// principal to, plus the operators valid for it.
type PrincipalType struct {
	Type      string   `json:"type"`      // Cedar entity type, e.g. "User", "Role", "Group"
	Label     string   `json:"label"`     // UI label
	Operators []string `json:"operators"` // valid Cedar operators, e.g. ["is","=="] or ["in"]
}

// PrincipalRef is one selectable principal instance.
//
// Value is always the entity's UUID, never its name. It is what the caller
// hands back to an attach endpoint (POST /roles/:roleID/policies,
// AttachPolicyToUser, ...), all of which take ids. Names are not usable as
// identifiers anyway: groups carry no unique index on (org, name), and two
// employees can share a display name outright.
//
// Label is the display name, and is for humans only.
type PrincipalRef struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// PrincipalsView is the full principal palette for an org: the static entity
// types plus every live attachment target — roles, groups and users.
//
// The three lists mirror three of the four attachment registers, so the builder
// can offer each path. (The fourth, the organization itself, needs no picker —
// there is exactly one org.)
//
// Users is populated by the ORCHESTRATOR, not the sentinel service: sentinel
// stores employee ids with no FK and has no dependency on the employee module,
// so it cannot resolve them. See sentinelOrchestrator.ListPrincipals.
type PrincipalsView struct {
	Types  []PrincipalType `json:"types"`
	Roles  []PrincipalRef  `json:"roles"`
	Groups []PrincipalRef  `json:"groups"`
	Users  []PrincipalRef  `json:"users"`
}

// RoleRef projects a role to the ref the builder shows as an attachment target.
func RoleRef(r *sentinelengine_models.Role) PrincipalRef {
	return PrincipalRef{Value: r.ID.String(), Label: r.Name}
}

// GroupRef projects a group to the ref the builder shows as an attachment
// target.
func GroupRef(g *sentinelengine_models.Group) PrincipalRef {
	return PrincipalRef{Value: g.ID.String(), Label: g.Name}
}
