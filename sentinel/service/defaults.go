package sentinelservice

// Defaults seeded into every new organization by ProvisionOrganization.
const (
	OwnerRoleName        = "Owner"
	OwnerRoleDescription = "Full administrative access to everything in the organization"

	OwnerPolicyName        = "Owner Full Access"
	OwnerPolicyDescription = "Grants members of the Owner role every action on every resource"

	StakeholdersGroupName        = "Stakeholders"
	StakeholdersGroupDescription = "Default group for the organization's stakeholders"
)

// ownerPolicyText is the Cedar source of the seeded full-access policy. Policy
// applicability is attachment-driven: this text permits everything for whoever
// it applies to, and it applies to owners because ProvisionOrganization attaches
// it to the Owner role (scope=role) — the text itself names no principal.
func ownerPolicyText() string {
	return "permit (\n  principal,\n  action,\n  resource\n);\n"
}
