package sentinelservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreateGroup implements [GroupsService]. A parent group, when given, must
// belong to the same org.
func (s *sentinelService) CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, errors.New("group name is required")
	}

	exists, err := s.repo.GroupExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("a group named %q already exists", name)
	}

	if in.ParentGroupID != nil {
		parent, err := s.repo.GetGroupByID(ctx, *in.ParentGroupID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("parent group: %w", err)
		}
		if parent.OrganizationID != orgID {
			return uuid.Nil, errors.New("parent group not found in this organization")
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
	return group.ID, nil
}

// ListGroups implements [GroupsService].
func (s *sentinelService) ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	groups, err := s.repo.ListGroupsByOrganizationID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	return sentinelDtos.ToGroupViews(groups), nil
}
