package sentinelservice

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"

	"github.com/google/uuid"
)

// principalTypes is the set of Cedar entity types a policy's PRINCIPAL CLAUSE
// may name, with the operators valid for each.
//
// User is the only one, because User is the only entity the engine builds.
// buildEntities puts exactly two entities in the map — the principal and the
// resource — so both forms below resolve:
//
//	permit(principal is User, ...)                  any user the policy reaches
//	permit(principal == User::"<employee-id>", ...)  that one employee
//
// Role and Group are deliberately ABSENT. They were previously advertised with
// the `in` operator, which invited `permit(principal in Role::"Owner", ...)` —
// text that compiles, stores, and displays as active while matching nothing,
// because no Role or Group entity exists to resolve `in` against. Roles and
// groups are not principals in Cedar; they are ATTACHMENT TARGETS, resolved in
// SQL before evaluation. They still appear in PrincipalsView.Roles/.Groups for
// exactly that purpose.
//
// Do not add an entity type here without first building it in
// sentinelEngine.buildEntities — otherwise every policy the builder emits for
// it is silently dead.
func principalTypes() []sentinelDtos.PrincipalType {
	return []sentinelDtos.PrincipalType{
		{Type: "User", Label: "Any user", Operators: []string{"is", "=="}},
	}
}

// GetPrincipals implements [SentinelServiceInterface]: the full principal
// palette for the policy builder — static types plus the org's roles and groups.
func (s *sentinelService) GetPrincipals(ctx context.Context, orgID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	roles, err := s.repo.ListRoles(ctx, orgID)
	if err != nil {
		return sentinelDtos.PrincipalsView{}, err
	}
	groups, err := s.repo.ListGroups(ctx, orgID)
	if err != nil {
		return sentinelDtos.PrincipalsView{}, err
	}

	view := sentinelDtos.PrincipalsView{
		Types:  principalTypes(),
		Roles:  make([]sentinelDtos.PrincipalRef, 0, len(roles)),
		Groups: make([]sentinelDtos.PrincipalRef, 0, len(groups)),
		// Sentinel cannot resolve employees; the orchestrator fills this in.
		// Initialised anyway so the field never serializes as null — a
		// consumer of this view alone would otherwise crash mapping over it.
		Users: make([]sentinelDtos.PrincipalRef, 0),
	}
	for _, r := range roles {
		view.Roles = append(view.Roles, sentinelDtos.RoleRef(r))
	}
	for _, g := range groups {
		view.Groups = append(view.Groups, sentinelDtos.GroupRef(g))
	}
	return view, nil
}
