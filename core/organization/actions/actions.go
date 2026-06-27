// Package organizationSentinalActions declares the org-level sentinel actions —
// the permissions that govern administration of the organization itself,
// including who may manage its authorization (policies, roles, attachments).
//
// These are the actions enforced on the "Organization" resource (see
// core/organization/resources). The route layer gates the /sentinel endpoints
// with them, so only the owner — or whoever the owner grants them to — can
// administer authorization.
//
// Action ids are stable Cedar strings; renaming one breaks any stored policy
// that references it.
package organizationSentinalActions

import (
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

const (
	// Sentinel administration.
	ViewSentinel sentinel_actions.Name = "viewSentinel" // read policies/roles/the builder
	CreatePolicy sentinel_actions.Name = "createPolicy"
	CreateRole   sentinel_actions.Name = "createRole"
	AttachPolicy sentinel_actions.Name = "attachPolicy" // attach a policy to a role
	DetachPolicy sentinel_actions.Name = "detachPolicy" // detach a policy from a role
)

// Actions is this module's contribution to the Organization resource descriptor.
var Actions = []sentinel_resources.Action{
	{Key: ViewSentinel, Label: "View authorization", Description: "View the org's policies, roles, and the policy builder."},
	{Key: CreatePolicy, Label: "Create policy", Description: "Create authorization policies."},
	{Key: CreateRole, Label: "Create role", Description: "Create roles."},
	{Key: AttachPolicy, Label: "Attach policy to role", Description: "Group policies under a role."},
	{Key: DetachPolicy, Label: "Detach policy from role", Description: "Remove a policy from a role."},
}
