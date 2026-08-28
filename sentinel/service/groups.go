package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelServiceErrs "sharpnr-engine/sentinel/errs"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateGroup implements [GroupsService]. A parent group, when given, must
// belong to the same org — the repo's org-scoped read is what enforces that.
func (s *sentinelService) CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, sentinelServiceErrs.GroupNameRequiredErr
	}

	// groups has no unique index on (org, name), so this check is the only
	// thing keeping names distinct.
	exists, err := s.repo.GroupExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("%w: %q", sentinelServiceErrs.GroupAlreadyExistsErr, name)
	}

	if in.ParentGroupID != nil {
		if _, err := s.repo.GetGroup(ctx, orgID, *in.ParentGroupID); err != nil {
			return uuid.Nil, sentinelServiceErrs.ParentGroupNotFoundInOrgErr
		}
	}

	group := &sentinelengine_models.Group{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Description:    in.Description,
		ParentGroupID:  in.ParentGroupID,
	}
	if err := s.repo.CreateGroup(ctx, group); err != nil {
		return uuid.Nil, err
	}
	// No engine invalidation: a new group has no policies registered against it,
	// so it changes nobody's effective policy set.
	return group.ID, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListGroups implements [GroupsService].
func (s *sentinelService) ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	groups, err := s.repo.ListGroups(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToGroupViews(groups), nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetEmployeeGroups implements [SentinelServiceInterface]: the groups an
// employee belongs to.
//
// One org-scoped JOIN. It used to fetch every group edge the employee had in
// ANY org, filter in Go, then issue one GetGroupByID per surviving edge — an
// N+1 that also depended on remembering the org check.
//
// Direct memberships only: belonging to a child group does not imply the parent.
func (s *sentinelService) GetEmployeeGroups(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	groups, err := s.repo.GetGroupsOfUser(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}
	// ToGroupViews returns an empty slice rather than nil, so this serializes
	// to [] instead of null.
	return sentinelDtos.ToGroupViews(groups), nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetGroupDetails implements [GroupService]: one group with the policies
// registered against it and the ids of its members.
//
// Member IDS, not names — sentinel has no dependency on the employee module and
// user_groups.user_id carries no FK. Resolving them is the orchestrator's job.
//
// Policies include disabled ones: an admin inspecting a group wants to see a
// policy that is registered but switched off. The engine is what filters on
// enabled.
func (s *sentinelService) GetGroupDetails(ctx context.Context, orgID, groupID uuid.UUID) (sentinelDtos.GroupDetailView, error) {
	group, err := s.repo.GetGroup(ctx, orgID, groupID)
	if err != nil {
		return sentinelDtos.GroupDetailView{}, fmt.Errorf("%w: %s", sentinelServiceErrs.GroupNotFoundInOrgErr, groupID)
	}

	policies, err := s.repo.ListPoliciesOfGroup(ctx, orgID, groupID)
	if err != nil {
		return sentinelDtos.GroupDetailView{}, err
	}

	memberIDs, err := s.repo.GetUsersOfGroup(ctx, orgID, groupID)
	if err != nil {
		return sentinelDtos.GroupDetailView{}, err
	}
	if memberIDs == nil {
		memberIDs = []uuid.UUID{} // serialize to [] rather than null
	}

	return sentinelDtos.GroupDetailView{
		Group:     sentinelDtos.ToGroupView(&group),
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		Policies:  sentinelDtos.ToPolicyViews(policies),
		MemberIDs: memberIDs,
	}, nil
}
