package sentinelservice

import "fmt"

// Defaults seeded into every new organization by ProvisionOrganization. Role and
// group names are Cedar entity ids embedded in policy text, so they must stay
// stable once released.
const (
	OwnerRoleName        = "Owner"
	OwnerRoleDescription = "Full administrative access to everything in the organization"

	OwnerPolicyName        = "Owner Full Access"
	OwnerPolicyDescription = "Grants members of the Owner role every action on every resource"

	StakeholdersGroupName        = "Stakeholders"
	StakeholdersGroupDescription = "Default group for the organization's stakeholders"
)

// ownerPolicyText is the Cedar source of the seeded full-access policy: anyone
// in the Owner role may do anything.
func ownerPolicyText() string {
	return fmt.Sprintf("permit (\n  principal in Role::%q,\n  action,\n  resource\n);\n", OwnerRoleName)
}
