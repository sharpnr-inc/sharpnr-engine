package sentinelservice

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"

	"github.com/google/uuid"
)

// principalTypes is the static half of the principal palette: the Cedar entity
// types a policy's principal can be scoped to, with the operators valid for
// each. Role/Group only make sense with `in` (membership); User supports `is`
// (any user) — the FE defers `== User::"<id>"` for now.
func principalTypes() []sentinelDtos.PrincipalType {
	return []sentinelDtos.PrincipalType{
		{Type: "User", Label: "User", Operators: []string{"is", "=="}},
		{Type: "Role", Label: "Role", Operators: []string{"in"}},
		{Type: "Group", Label: "Group", Operators: []string{"in"}},
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
	}
	for _, r := range roles {
		view.Roles = append(view.Roles, sentinelDtos.RoleRef(r))
	}
	for _, g := range groups {
		view.Groups = append(view.Groups, sentinelDtos.GroupRef(g))
	}
	return view, nil
}
