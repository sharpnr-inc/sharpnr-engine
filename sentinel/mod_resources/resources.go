package sentinelOwnResources

import (
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
)

type sentinelOwnResources struct {
	sentinelService sentinelservice.SentinelServiceInterface
}

// Compile-time guard: the provider must satisfy the aggregator contract.
// var _ sentinelAggregatorContracts.ResourceProvider = (*sentinelOwnResources)(nil)

func NewSentinelOwnResources(sentinelService sentinelservice.SentinelServiceInterface) *sentinelOwnResources {
	return &sentinelOwnResources{sentinelService: sentinelService}
}

// Resources implements [sentinelAggregatorContracts.ResourceProvider].
func (s *sentinelOwnResources) Resources() []sentinelAggregatorContracts.ResourceDescriptor {
	return []sentinelAggregatorContracts.ResourceDescriptor{
		{
			Type:        ResourceTypeSentinelPolicy,
			Label:       "Authorization policy",
			Description: "A Cedar policy stored by the Sentinel authorization system",
			Actions: []sentinelAggregatorContracts.ActionDescriptor{
				{Name: ActionCanCreatePolicy, Label: "Create policy", Description: "Author a new authorization policy"},
				{Name: ActionCanViewPolicy, Label: "View policy", Description: "View policies and their Cedar text", RequiresInstance: true},
				{Name: ActionCanEditPolicy, Label: "Edit policy", Description: "Edit a policy's Cedar text or metadata", RequiresInstance: true},
				{Name: ActionCanDeletePolicy, Label: "Delete policy", Description: "Delete a policy", RequiresInstance: true},
				{Name: ActionCanEnablePolicy, Label: "Enable policy", Description: "Enable a disabled policy", RequiresInstance: true},
				{Name: ActionCanDisablePolicy, Label: "Disable policy", Description: "Disable a policy without deleting it", RequiresInstance: true},
			},
			Attributes: []sentinelAggregatorContracts.AttributeDescriptor{
				{Name: "name", Type: "string", Description: "Policy name"},
				{Name: "scope", Type: "string", Description: "Classification bucket: organization | role | group | user"},
				{Name: "enabled", Type: "bool", Description: "Whether the policy is active in the engine"},
			},
		},
		{
			Type:        ResourceTypeSentinelRole,
			Label:       "Role",
			Description: "A named role that policies reference and users are assigned to",
			Actions: []sentinelAggregatorContracts.ActionDescriptor{
				{Name: ActionCanCreateRole, Label: "Create role", Description: "Create a new role"},
				{Name: ActionCanViewRole, Label: "View role", Description: "View roles and their privileges", RequiresInstance: true},
				{Name: ActionCanEditRole, Label: "Edit role", Description: "Edit a role's details", RequiresInstance: true},
				{Name: ActionCanDeleteRole, Label: "Delete role", Description: "Delete a role", RequiresInstance: true},
				{Name: ActionCanAssignRole, Label: "Assign role", Description: "Assign a role to an employee", RequiresInstance: true},
				{Name: ActionCanAttachPolicyToRole, Label: "Attach policy", Description: "Group a policy under a role", RequiresInstance: true},
				{Name: ActionCanDetachPolicyFromRole, Label: "Detach policy", Description: "Remove a policy from a role", RequiresInstance: true},
			},
			Attributes: []sentinelAggregatorContracts.AttributeDescriptor{
				{Name: "name", Type: "string", Description: "Role name (the Cedar entity id)"},
			},
		},
		{
			Type:        ResourceTypeSentinelGroup,
			Label:       "Group",
			Description: "A named group users belong to; groups can nest",
			Actions: []sentinelAggregatorContracts.ActionDescriptor{
				{Name: ActionCanCreateGroup, Label: "Create group", Description: "Create a new group"},
				{Name: ActionCanViewGroup, Label: "View group", Description: "View groups and their members", RequiresInstance: true},
				{Name: ActionCanEditGroup, Label: "Edit group", Description: "Edit a group's details", RequiresInstance: true},
				{Name: ActionCanDeleteGroup, Label: "Delete group", Description: "Delete a group", RequiresInstance: true},
				{Name: ActionCanManageGroupMembers, Label: "Manage members", Description: "Add or remove group members", RequiresInstance: true},
			},
			Attributes: []sentinelAggregatorContracts.AttributeDescriptor{
				{Name: "name", Type: "string", Description: "Group name (the Cedar entity id)"},
			},
		},
	}
}

// ListInstances implements [sentinelAggregatorContracts.ResourceProvider]: the
// org's live policies, roles or groups as selectable policy targets.

// func (s *sentinelOwnResources) ListInstances(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) ([]sentinelAggregatorContracts.ResourceInstance, error) {
// 	switch resourceType {

// 	case ResourceTypeSentinelPolicy:
// 		policies, err := s.sentinelService.ListPolicies(ctx, orgID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(policies))
// 		for _, p := range policies {
// 			instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: p.ID.String(), Label: p.Name})
// 		}
// 		return instances, nil

// 	case ResourceTypeSentinelRole:
// 		roles, err := s.sentinelService.ListRoles(ctx, orgID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(roles))
// 		for _, r := range roles {
// 			instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: r.ID.String(), Label: r.Name})
// 		}
// 		return instances, nil

// 	case ResourceTypeSentinelGroup:
// 		groups, err := s.sentinelService.ListGroups(ctx, orgID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(groups))
// 		for _, g := range groups {
// 			instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: g.ID.String(), Label: g.Name})
// 		}
// 		return instances, nil

// 	default:
// 		return nil, fmt.Errorf("sentinel resources: unknown resource type %q", resourceType)
// 	}
// }
