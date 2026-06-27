package sentinelengine

import (
	"context"
	"fmt"

	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"
)

// buildEntities assembles the principal's membership graph and the target
// resource into an EntityMap. Both principal and resource carry an orgId
// attribute so the `resource.orgId == principal.orgId` tenant guard holds.
//
// Rebuilt on every authorization call (unlike the policy set, which is cached),
// because membership and principal attributes change constantly.
func (s *Sentinel) buildEntities(ctx context.Context, in AuthorizeInput) (types.EntityMap, error) {

	em := types.EntityMap{}

	orgID := in.OrgID.String()
	orgUID := cedar.NewEntityUID(EntityOrganization, types.String(orgID))
	em[orgUID] = types.Entity{UID: orgUID}

	roleSlugs, err := s.repo.UserRoleSlugs(ctx, in.OrgID, in.UserID)
	if err != nil {
		return nil, fmt.Errorf("load roles: %w", err)
	}

	groups, err := s.repo.UserGroupClosure(ctx, in.OrgID, in.UserID)
	if err != nil {
		return nil, fmt.Errorf("load groups: %w", err)
	}
	attrs, err := s.repo.UserAttributes(ctx, in.OrgID, in.UserID)
	if err != nil {
		return nil, fmt.Errorf("load principal attributes: %w", err)
	}

	userParents := []types.EntityUID{orgUID}

	for _, slug := range roleSlugs {
		ruid := cedar.NewEntityUID(EntityRole, types.String(slug))
		em[ruid] = types.Entity{UID: ruid, Parents: cedar.NewEntityUIDSet(orgUID)}
		userParents = append(userParents, ruid)
	}

	for _, g := range groups {
		guid := cedar.NewEntityUID(EntityGroup, types.String(g.Slug))
		parents := []types.EntityUID{orgUID}
		if g.ParentSlug != "" {
			parents = append(parents, cedar.NewEntityUID(EntityGroup, types.String(g.ParentSlug)))
		}
		em[guid] = types.Entity{UID: guid, Parents: cedar.NewEntityUIDSet(parents...)}
		if g.Direct {
			userParents = append(userParents, guid)
		}
	}

	// Principal (the employee). orgId is engine-owned (the tenant guard relies on
	// it) and cannot be overridden by a loaded attribute; everything else from
	// UserAttributes is merged in so policies can read `principal.<attr>`.
	principalAttrs := cedar.RecordMap{"orgId": types.String(orgID)}
	for k, v := range attrs {
		if k == "orgId" {
			continue
		}
		principalAttrs[types.String(k)] = v
	}
	userUID := cedar.NewEntityUID(EntityUser, types.String(in.UserID.String()))
	em[userUID] = types.Entity{
		UID:        userUID,
		Parents:    cedar.NewEntityUIDSet(userParents...),
		Attributes: cedar.NewRecord(principalAttrs),
	}

	// Resource, with orgId merged into any caller-supplied attributes.
	resAttrs := cedar.RecordMap{"orgId": types.String(orgID)}
	for k, v := range in.Resource.Attrs {
		resAttrs[types.String(k)] = v
	}
	resUID := cedar.NewEntityUID(types.EntityType(in.Resource.Type), types.String(in.Resource.ID))
	em[resUID] = types.Entity{
		UID:        resUID,
		Parents:    cedar.NewEntityUIDSet(in.Resource.Parents...),
		Attributes: cedar.NewRecord(resAttrs),
	}

	return em, nil
}
