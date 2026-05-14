package legal_entity_constants

type LegalEntityAddressType string

const (
	AddressTypeRegistered LegalEntityAddressType = "REGISTERED"
	AddressTypeOperating  LegalEntityAddressType = "OPERATING"
	AddressTypeGST        LegalEntityAddressType = "GST"
	AddressTypeWarehouse  LegalEntityAddressType = "WAREHOUSE"
	AddressTypeBranch     LegalEntityAddressType = "BRANCH"
)
