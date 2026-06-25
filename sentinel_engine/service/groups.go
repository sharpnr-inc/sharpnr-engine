package sentinelservice

import (
	"context"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// CreateGroup creates a group, optionally nested under a parent. The slug is the
// group's stable Cedar entity id (policies reference Group::"<slug>"), derived
// from the name when not supplied.
func (s *sentinelService) CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, fmt.Errorf("group name is required")
	}

	slug := strings.TrimSpace(in.Slug)
	if slug == "" {
		slug = slugify(name)
	}

	group := &sentinelengine_models.Group{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Slug:           slug,
		Name:           name,
		Description:    in.Description,
		ParentGroupID:  in.ParentGroupID,
	}
	if err := s.repo.CreateGroup(ctx, group); err != nil {
		return uuid.Nil, fmt.Errorf("create group: %w", err)
	}
	return group.ID, nil
}

// ListGroups returns every group defined for an org (the group tree source).
func (s *sentinelService) ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	groups, err := s.repo.ListGroupsByOrgID(ctx, orgID)
	if err != nil {
		return nil, fmt.Errorf("list groups: %w", err)
	}
	return sentinelDtos.ToGroupViews(groups), nil
}
