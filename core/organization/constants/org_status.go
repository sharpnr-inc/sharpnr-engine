package organization_constants

/* ---------------------------------- <***> --------------------------------- */

// OrgSetupStatusType represents the current configuration state
// of an organization's setup workflow.
type OrgSetupStatusType string

const (
	// OrgConfigStatusNotConfigured indicates that the organization
	// setup exists but is not yet fully configured.
	OrgConfigStatusNotConfigured OrgSetupStatusType = "NOT_CONFIGURED"

	// OrgConfigStatusConfigured indicates that the organization
	// setup has been successfully completed and configured.
	OrgConfigStatusConfigured OrgSetupStatusType = "CONFIGURED"

	// OrgConfigStatusConfigNotStarted indicates that the organization
	// setup process has not been started yet.
	OrgConfigStatusConfigNotStarted OrgSetupStatusType = "CONFIGURATION_NOT_STARTED"
)

/* ---------------------------------- <***> --------------------------------- */

type OrgStatusType string

const (
	OrgStatusActive    OrgStatusType = "ACTIVE"
	OrgStatusSuspended OrgStatusType = "SUSPENDED"
)

/* ---------------------------------- <***> --------------------------------- */
