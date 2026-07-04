// Package sentinelOwnResources is the Sentinel module's own mod_resources: the
// authorization system is itself a protectable resource, so policy authoring,
// role management and group management get Cedar resource types and actions
// like any other module.
package sentinelOwnResources

// Cedar entity types of the Sentinel module's own resources.
const (
	ResourceTypeSentinelPolicy = "SentinelPolicy"
	ResourceTypeSentinelRole   = "SentinelRole"
	ResourceTypeSentinelGroup  = "SentinelGroup"
)

// Cedar action ids on SentinelPolicy.
const (
	ActionCanCreatePolicy  = "canCreatePolicy"
	ActionCanViewPolicy    = "canViewPolicy"
	ActionCanEditPolicy    = "canEditPolicy"
	ActionCanDeletePolicy  = "canDeletePolicy"
	ActionCanEnablePolicy  = "canEnablePolicy"
	ActionCanDisablePolicy = "canDisablePolicy"
)

// Cedar action ids on SentinelRole.
const (
	ActionCanCreateRole           = "canCreateRole"
	ActionCanViewRole             = "canViewRole"
	ActionCanEditRole             = "canEditRole"
	ActionCanDeleteRole           = "canDeleteRole"
	ActionCanAssignRole           = "canAssignRole"
	ActionCanAttachPolicyToRole   = "canAttachPolicyToRole"
	ActionCanDetachPolicyFromRole = "canDetachPolicyFromRole"
)

// Cedar action ids on SentinelGroup.
const (
	ActionCanCreateGroup        = "canCreateGroup"
	ActionCanViewGroup          = "canViewGroup"
	ActionCanEditGroup          = "canEditGroup"
	ActionCanDeleteGroup        = "canDeleteGroup"
	ActionCanManageGroupMembers = "canManageGroupMembers"
)
