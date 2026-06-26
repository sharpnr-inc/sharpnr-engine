package sentinelservice

import (
	"context"
	"fmt"

	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"

	"github.com/google/uuid"
)

// principalTypes is the static set of Cedar entity types the policy builder can
// scope a principal to. The type strings come from the engine's exported entity
// names so the builder advertises exactly what the evaluator resolves. Selecting
// a specific user (== User::"<uuid>") is deferred, so User exposes only "is".
func principalTypes() []sentinelDtos.PrincipalType {
	return []sentinelDtos.PrincipalType{
		{Type: sentinelengine.EntityUser, Label: "User", Operators: []string{"is"}},
		{Type: sentinelengine.EntityRole, Label: "Role", Operators: []string{"in"}},
		{Type: sentinelengine.EntityGroup, Label: "Group", Operators: []string{"in"}},
	}
}

// GetPrincipals returns the org's principal palette for the builder: the static
// entity types plus the org's roles and groups (keyed by the slug the Cedar text
// embeds).
func (s *sentinelService) GetPrincipals(ctx context.Context, orgID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	roles, err := s.repo.ListRolesByOrgID(ctx, orgID)
	if err != nil {
		return sentinelDtos.PrincipalsView{}, fmt.Errorf("list roles for principals: %w", err)
	}
	groups, err := s.repo.ListGroupsByOrgID(ctx, orgID)
	if err != nil {
		return sentinelDtos.PrincipalsView{}, fmt.Errorf("list groups for principals: %w", err)
	}

	roleRefs := make([]sentinelDtos.PrincipalRef, 0, len(roles))
	for _, r := range roles {
		roleRefs = append(roleRefs, sentinelDtos.RoleRef(r))
	}
	groupRefs := make([]sentinelDtos.PrincipalRef, 0, len(groups))
	for _, g := range groups {
		groupRefs = append(groupRefs, sentinelDtos.GroupRef(g))
	}

	return sentinelDtos.PrincipalsView{
		Types:  principalTypes(),
		Roles:  roleRefs,
		Groups: groupRefs,
	}, nil
}

// ListResources returns every registered resource descriptor — the resource
// types a policy can be scoped to, with their actions and attributes.
func (s *sentinelService) ListResources() []sentinel_resources.Descriptor {
	return sentinel_resources.Descriptors()
}

// GetResourceActions returns the actions declared for one resource type. Errors
// for an unregistered type so a bad path param surfaces instead of an empty list.
func (s *sentinelService) GetResourceActions(resourceType string) ([]sentinel_resources.Action, error) {
	d, ok := sentinel_resources.DescriptorFor(resourceType)
	if !ok {
		return nil, fmt.Errorf("unknown resource type %q", resourceType)
	}
	return d.Actions, nil
}

// GetResourceAttributes returns the ABAC attributes declared for one resource
// type. Errors for an unregistered type.
func (s *sentinelService) GetResourceAttributes(resourceType string) ([]sentinel_resources.Attribute, error) {
	d, ok := sentinel_resources.DescriptorFor(resourceType)
	if !ok {
		return nil, fmt.Errorf("unknown resource type %q", resourceType)
	}
	return d.Attributes, nil
}

// ListResourceInstances returns the org's live instances of one resource type,
// for the builder's `resource in <Type>::"<id>"` picker. Errors if no provider is
// registered for the type (a bad type surfaces instead of an empty list).
func (s *sentinelService) ListResourceInstances(ctx context.Context, orgID uuid.UUID, resourceType string) ([]sentinel_resources.Ref, error) {
	return sentinel_resources.List(ctx, resourceType, orgID)
}
