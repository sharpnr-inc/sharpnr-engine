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
// belong to the same org.
func (s *sentinelService) CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return uuid.Nil, sentinelServiceErrs.GroupNameRequiredErr
	}

	exists, err := s.repo.GroupExists(ctx, orgID, name)
	if err != nil {
		return uuid.Nil, err
	}
	if exists {
		return uuid.Nil, fmt.Errorf("%w: %q", sentinelServiceErrs.GroupAlreadyExistsErr, name)
	}

	if in.ParentGroupID != nil {
		parent, err := s.repo.GetGroupByID(ctx, *in.ParentGroupID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("parent group: %w", err)
		}
		if parent.OrganizationID != orgID {
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

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelService) GetEmployeeGroups(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.GroupView, error) {

	// Fetch all group assignments for the given employee
	userGroups, err := s.repo.GetGroupsAssignedToUser(ctx, emplID)
	if err != nil {
		return nil, err
	}

	// Initialize as an empty slice so it serializes to [] instead of null in JSON
	result := make([]sentinelDtos.GroupView, 0)

	for _, userGroup := range userGroups {
		// Enforce tenant scoping by filtering by organization ID
		if userGroup.OrganizationID == orgID {
			groupDetails, err := s.repo.GetGroupByID(ctx, userGroup.GroupID)
			if err != nil {
				// Fail fast if a group cannot be fetched, or log and continue depending on requirements
				return nil, err
			}

			// Map the sentinelengine_models.Group to sentinelDtos.GroupView
			groupView := sentinelDtos.GroupView{
				ID:            groupDetails.ID,
				Name:          groupDetails.Name,
				Description:   groupDetails.Description,
				ParentGroupID: groupDetails.ParentGroupID, // Directly maps *uuid.UUID
			}

			result = append(result, groupView)
		}
	}

	return result, nil
}
