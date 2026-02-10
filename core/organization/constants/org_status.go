package organization_constants

type OrganizationStatus string

const (
	OrgNotConfigured OrganizationStatus = "NOT_CONFIGURED"
	OrgActive        OrganizationStatus = "ACTIVE"
	OrgSuspended     OrganizationStatus = "SUSPENDED"
)
