package sentinel_utils

// this is the package we can have for the defaults
// in this file we create a function which will create the
// default super Admin and store the polices

// Default slugs created for every new organization.
const (
	OwnerRoleSlug = "owner" // full access within the org; NOT platform-level
)

// OwnerPolicyText is the Cedar policy granting the owner role full access,
// scoped to its own tenant by the orgId guard.
const OwnerPolicyText = `permit (
  principal in Role::"owner",
  action,
  resource
) when { resource.orgId == principal.orgId };`
