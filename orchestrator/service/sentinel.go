package orchestrator_service

import (
	"context"

	coreEmployeeServices "sharpnr-engine/core/employees/services"
	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// Sentinel (authorization) orchestration. These methods are thin pass-throughs
// to the Sentinel service — the orchestrator is the HTTP facade the frontend
// talks to, and authorization authoring lives in sentinel_engine. orgID/emplID
// are carried for tenant scope and future per-call authorization, even where a
// delegate doesn't need them yet.

// sentinelOrchestrator implements SentinelOrchestrationInterface.
//
// guard is the seam for the missing authoring authorization: policy/role
// creation and attachment must call guard.Authorize before delegating (rule 1
// in doc.go). The checks are not wired yet — do not ship policy authoring to
// production before they are.
type sentinelOrchestrator struct {
	guard              *orchestrator_authz.Guard
	sentinelService    sentinelservice.SentinelServiceInterface
	sentinelAggregator sentinelAggregatorService.SentinelAggregatorServiceInterface

	// employeeHRService resolves member ids to names. Sentinel stores group
	// membership as bare employee ids with no FK and no knowledge of the
	// employee module; composing the two is exactly this layer's job.
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface
}

func (s *sentinelOrchestrator) ListPrincipals(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.PrincipalsView, error) {
	return s.sentinelService.GetPrincipals(ctx, orgID)
}

func (s *sentinelOrchestrator) ListAuthzResources(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorService.ResourceSummaryView, error) {
	return s.sentinelAggregator.ListResources(ctx, orgID, emplID)
}

func (s *sentinelOrchestrator) GetAuthzResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (sentinelAggregatorService.ResourceDetailsView, error) {
	return s.sentinelAggregator.GetResourceDetails(ctx, orgID, emplID, resourceType)
}

func (s *sentinelOrchestrator) GetAuthzContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error) {
	return s.sentinelAggregator.GetContextPalette(ctx, orgID, emplID)
}

func (s *sentinelOrchestrator) CreatePolicy(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error) {
	return s.sentinelService.CreatePolicy(ctx, orgID, in)
}

func (s *sentinelOrchestrator) ListPolicies(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.PolicyView, error) {
	return s.sentinelService.ListPolicies(ctx, orgID)
}

func (s *sentinelOrchestrator) GetPolicy(ctx context.Context, orgID, emplID, policyID uuid.UUID) (sentinelDtos.PolicyView, error) {
	return s.sentinelService.GetPolicy(ctx, orgID, policyID)
}

func (s *sentinelOrchestrator) ListRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.RoleView, error) {
	return s.sentinelService.ListRoles(ctx, orgID)
}

func (s *sentinelOrchestrator) CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error) {
	return s.sentinelService.CreateRole(ctx, orgID, emplID, in)
}

func (s *sentinelOrchestrator) GetRolePrivileges(ctx context.Context, orgID, emplID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error) {
	return s.sentinelService.GetRolePrivileges(ctx, orgID, roleID)
}

func (s *sentinelOrchestrator) AttachPoliciesToRole(ctx context.Context, orgID, emplID, roleID uuid.UUID, in sentinelDtos.AttachPoliciesInput) error {
	return s.sentinelService.AttachPolicies(ctx, orgID, emplID, roleID, in.PolicyIDs)
}

func (s *sentinelOrchestrator) DetachPolicyFromRole(ctx context.Context, orgID, emplID, roleID, policyID uuid.UUID) error {
	return s.sentinelService.DetachPolicy(ctx, orgID, roleID, policyID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelOrchestrator) ListSentinelUserGroups(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.GroupView, error) {
	return s.sentinelService.ListGroups(ctx, orgID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelOrchestrator) CreateNewSentinelGroup(ctx context.Context, orgID, emplID uuid.UUID, req *sentinelDtos.CreateGroupsRequestDto) error {
	_, err := s.sentinelService.CreateGroup(ctx, orgID, *req)
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetSentinelGroupDetail returns one group with its policies and its members.
//
// Two domains meet here: sentinel owns the group, its policy registrations and
// the ids of its members; the employee module owns who those ids are. Sentinel
// is asked first, then the ids are resolved to names.
func (s *sentinelOrchestrator) GetSentinelGroupDetail(ctx context.Context, orgID, emplID, groupID uuid.UUID) (sentinelDtos.GroupDetailsResponseDto, error) {
	detail, err := s.sentinelService.GetGroupDetails(ctx, orgID, groupID)
	if err != nil {
		return sentinelDtos.GroupDetailsResponseDto{}, err
	}

	members, err := s.resolveGroupMembers(ctx, orgID, emplID, detail.MemberIDs)
	if err != nil {
		return sentinelDtos.GroupDetailsResponseDto{}, err
	}

	policies := make([]sentinelDtos.GroupPolicesResponseDto, 0, len(detail.Policies))
	for _, p := range detail.Policies {
		policies = append(policies, sentinelDtos.GroupPolicesResponseDto{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return sentinelDtos.GroupDetailsResponseDto{
		ID:          detail.Group.ID,
		Name:        detail.Group.Name,
		Description: detail.Group.Description,
		CreatedAt:   detail.CreatedAt,
		UpdatedAt:   detail.UpdatedAt,
		User:        members,
		Polices:     policies,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// resolveGroupMembers turns sentinel's member ids into {id, name} pairs.
//
// It indexes the org's employees once rather than fetching each member, so this
// stays a single employee-module call regardless of group size. The trade is
// that it reads the whole org roster to render one group; if that becomes hot,
// the employee module wants a "list employees by ids" method and this becomes
// one narrow query.
//
// An id with no matching employee is kept with an empty name rather than
// dropped: a membership row pointing at a removed employee is a real state, and
// hiding it would make the group look smaller than it is.
func (s *sentinelOrchestrator) resolveGroupMembers(ctx context.Context, orgID, emplID uuid.UUID, memberIDs []uuid.UUID) ([]sentinelDtos.GroupUserResponseDto, error) {
	members := make([]sentinelDtos.GroupUserResponseDto, 0, len(memberIDs))
	if len(memberIDs) == 0 {
		return members, nil
	}
	if s.employeeHRService == nil {
		// Fail soft: the group and its policies are still worth returning.
		for _, id := range memberIDs {
			members = append(members, sentinelDtos.GroupUserResponseDto{ID: id})
		}
		return members, nil
	}

	employees, err := s.employeeHRService.ListOrgAllEmployees(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}
	byID := make(map[uuid.UUID]string, len(employees))
	for _, e := range employees {
		if e != nil {
			byID[e.ID] = e.EmployeeName
		}
	}

	for _, id := range memberIDs {
		members = append(members, sentinelDtos.GroupUserResponseDto{ID: id, Name: byID[id]})
	}
	return members, nil
}
