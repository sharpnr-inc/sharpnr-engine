package sentinelOwnResources

import (
	"context"

	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorTypes "sharpnr-engine/sentinel_aggregators/types"

	"github.com/google/uuid"
)

type sentinelOwnResources struct {
	sentinelService sentinelservice.SentinelServiceInterface
}

// Compile-time guard: the provider must satisfy the aggregator contract.
var _ sentinelAggregatorContracts.ResourceProvider = (*sentinelOwnResources)(nil)

func NewSentinelOwnResources(sentinelService sentinelservice.SentinelServiceInterface) *sentinelOwnResources {
	return &sentinelOwnResources{sentinelService: sentinelService}
}

// Resources implements [sentinelAggregatorContracts.ResourceProvider]: the
// authorization system is itself protectable, so it registers three resource
// types like any other module. Each carries its own lister — there is no
// resource-type switch anywhere in this file.
func (s *sentinelOwnResources) Resources() []sentinelAggregatorContracts.ResourceRegistration {
	return []sentinelAggregatorContracts.ResourceRegistration{
		{
			Descriptor: sentinelAggregatorContracts.ResourceDescriptor{
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
					{Name: "name", Type: sentinelAggregatorTypes.AttrString, Description: "Policy name"},
					{Name: "enabled", Type: sentinelAggregatorTypes.AttrBool, Description: "Whether the policy is active in the engine"},
				},
			},
			ListInstances: s.listPolicies,
		},
		{
			Descriptor: sentinelAggregatorContracts.ResourceDescriptor{
				Type:        ResourceTypeSentinelRole,
				Label:       "Role",
				Description: "A named role that policies are registered against and users are assigned to",
				Actions: []sentinelAggregatorContracts.ActionDescriptor{
					{Name: ActionCanCreateRole, Label: "Create role", Description: "Create a new role"},
					{Name: ActionCanViewRole, Label: "View role", Description: "View roles and their privileges", RequiresInstance: true},
					{Name: ActionCanEditRole, Label: "Edit role", Description: "Edit a role's details", RequiresInstance: true},
					{Name: ActionCanDeleteRole, Label: "Delete role", Description: "Delete a role", RequiresInstance: true},
					{Name: ActionCanAssignRole, Label: "Assign role", Description: "Assign a role to an employee", RequiresInstance: true},
					{Name: ActionCanAttachPolicyToRole, Label: "Attach policy", Description: "Register a policy against a role", RequiresInstance: true},
					{Name: ActionCanDetachPolicyFromRole, Label: "Detach policy", Description: "Remove a policy from a role", RequiresInstance: true},
				},
				Attributes: []sentinelAggregatorContracts.AttributeDescriptor{
					{Name: "name", Type: sentinelAggregatorTypes.AttrString, Description: "Role name"},
				},
			},
			ListInstances: s.listRoles,
		},
		{
			Descriptor: sentinelAggregatorContracts.ResourceDescriptor{
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
					{Name: "name", Type: sentinelAggregatorTypes.AttrString, Description: "Group name"},
				},
			},
			ListInstances: s.listGroups,
		},
	}
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelOwnResources) listPolicies(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceInstance, error) {
	policies, err := s.sentinelService.ListPolicies(ctx, orgID)
	if err != nil {
		return nil, err
	}
	instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(policies))
	for _, p := range policies {
		instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: p.ID.String(), Label: p.Name})
	}
	return instances, nil
}

func (s *sentinelOwnResources) listRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceInstance, error) {
	roles, err := s.sentinelService.ListRoles(ctx, orgID)
	if err != nil {
		return nil, err
	}
	instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(roles))
	for _, r := range roles {
		instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: r.ID.String(), Label: r.Name})
	}
	return instances, nil
}

func (s *sentinelOwnResources) listGroups(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceInstance, error) {
	groups, err := s.sentinelService.ListGroups(ctx, orgID)
	if err != nil {
		return nil, err
	}
	instances := make([]sentinelAggregatorContracts.ResourceInstance, 0, len(groups))
	for _, g := range groups {
		instances = append(instances, sentinelAggregatorContracts.ResourceInstance{ID: g.ID.String(), Label: g.Name})
	}
	return instances, nil
}
