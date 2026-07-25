package property_constants

type PropertyRunningStatus string

const (
	Active    PropertyRunningStatus = "ACTIVE"
	NotActive PropertyRunningStatus = "NOT_ACTIVE"
)

type PropertyConfigurationStatus string

const (
	Configured    PropertyConfigurationStatus = "CONFIGURED"
	NotConfigured PropertyConfigurationStatus = "NOT_CONFIGURED"
)
