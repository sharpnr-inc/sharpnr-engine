package sentinelservice

// Well-known slugs and labels seeded at provisioning. The owner role's slug must
// stay "owner" because the default owner policy (default_policies.OwnerPolcy)
// references Role::"owner" by text, and the engine resolves roles by slug.
const (
	ownerRoleSlug   = "owner"
	ownerRoleName   = "Owner"
	ownerRoleDesc   = "Full access within the organization."
	ownerPolicyName = "Owner — Full Access"
	ownerPolicyDesc = "Grants the owner role full access to every resource in the organization."

	stakeholdersGroupSlug = "stakeholders"
	stakeholdersGroupName = "Stakeholders"
	stakeholdersGroupDesc = "Top-level group every organization starts with."
)
