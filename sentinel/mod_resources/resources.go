package sentinelOwnResources

import (
	sentinelservice "sharpnr-engine/sentinel/service"
	agg "sharpnr-engine/sentinel_aggregators/contracts"
)

const categoryAdministration = "Administration"

type sentinelOwnResources struct {
	sentinelService sentinelservice.SentinelServiceInterface
}

// Compile-time guard: the provider must satisfy the aggregator contract.
var _ agg.ResourceProvider = (*sentinelOwnResources)(nil)

func NewSentinelOwnResources(sentinelService sentinelservice.SentinelServiceInterface) *sentinelOwnResources {
	return &sentinelOwnResources{sentinelService: sentinelService}
}

// Resources implements [agg.ResourceProvider].
//
// All three types are organization-level, so none declares HousedIn — the
// engine stamps the Organization edge itself from the request's org.
//
// None declares attributes either. The previous version advertised `name` and
// `enabled` as conditionable, with nothing anywhere to resolve them: Cedar skips
// a policy whose attribute is missing, so a rule like `resource.name != "Owner"`
// would have read as active and matched nothing. The contract now forbids
// declaring what cannot be answered, and there is no real use case here worth
// writing a resolver for — protecting the Owner role belongs in code, not in a
// condition an admin can delete.
func (s *sentinelOwnResources) Resources() []agg.ResourceRegistration {
	return []agg.ResourceRegistration{
		agg.Standard(ResourceTypeSentinelPolicy, "Permission rules").
			Category(categoryAdministration).
			Describe("The underlying rules that grant access").
			WithActions(
				agg.Action(ActionCanEnableSentinelPolicy, "Turn rules on").
					Describe("Re-enable a rule that was switched off").
					OnInstance(),
				agg.Action(ActionCanDisableSentinelPolicy, "Turn rules off").
					Describe("Switch a rule off without deleting it").
					OnInstance(),
				agg.Action(ActionCanSetOrganizationPermissions, "Set organization-wide permissions").
					Describe("Change what everyone in the organization can do"),
			).
			ListInstances(s.listPolicies).
			Build(),

		agg.Standard(ResourceTypeSentinelRole, "Roles").
			Category(categoryAdministration).
			Describe("Job roles and what each one is allowed to do").
			WithActions(
				agg.Action(ActionCanAssignSentinelRole, "Assign roles to people").
					Describe("Decide who holds which role").
					OnInstance(),
				// The permission matrix writes through this one action. Attach
				// and detach were two actions over a register the UI no longer
				// exposes separately — a save is one operation, so it is one
				// permission.
				agg.Action(ActionCanSetSentinelRolePermissions, "Change what a role can do").
					Describe("Set the permissions granted by a role").
					OnInstance(),
			).
			ListInstances(s.listRoles).
			Build(),

		agg.Standard(ResourceTypeSentinelGroup, "Groups").
			Category(categoryAdministration).
			Describe("Teams and departments people belong to").
			WithActions(
				agg.Action(ActionCanManageSentinelGroupMembers, "Add and remove members").
					Describe("Change who belongs to a group").
					OnInstance(),
				agg.Action(ActionCanSetSentinelGroupPermissions, "Change what a group can do").
					Describe("Set the permissions granted by a group").
					OnInstance(),
			).
			ListInstances(s.listGroups).
			Build(),
	}
}
