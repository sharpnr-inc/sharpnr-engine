package sentinel_utils

import "github.com/google/uuid"

// this is the package we can have for the defaults
// in this file we create a function which will create the
// default super Admin and store the polices

// Default slugs created for every new organization.
const (
	OwnerRoleSlug = "owner" // full access within the org; NOT platform-level
)

// SystemFullAccessPolicyID is the well-known id of the shared "Full Access"
// system managed policy (organization_id NULL), seeded in migration 000032 and
// attached to every org's owner role at provisioning. Fixed (app-chosen, not
// DB-generated) so provisioning can reference it without a lookup by name.
var SystemFullAccessPolicyID = uuid.MustParse("f0000000-0000-4000-8000-000000000001")

// OwnerPolicyText is the Cedar policy granting the owner role full access,
// scoped to its own tenant by the orgId guard.
const OwnerPolicyText = `permit (
  principal in Role::"owner",
  action,
  resource
) when { resource.orgId == principal.orgId };`
