package legal_entity_constants

type LegalEntityAddressType string

const (
	AddressTypeRegistered LegalEntityAddressType = "REGISTERED"
	AddressTypeOperating  LegalEntityAddressType = "OPERATING"
	AddressTypeGST        LegalEntityAddressType = "GST"
	AddressTypeWarehouse  LegalEntityAddressType = "WAREHOUSE"
	AddressTypeBranch     LegalEntityAddressType = "BRANCH"
)

// Label returns a human-readable presentation string for UI components
func (t LegalEntityAddressType) Label() string {
	switch t {
	case AddressTypeRegistered:
		return "Registered Office"
	case AddressTypeOperating:
		return "Operating / Corporate Office"
	case AddressTypeGST:
		return "GST Registered Address"
	case AddressTypeWarehouse:
		return "Warehouse / Storage Facility"
	case AddressTypeBranch:
		return "Branch Office"
	default:
		return string(t)
	}
}

// AllAddressTypes returns all supported address classifications
func AllAddressTypes() []LegalEntityAddressType {
	return []LegalEntityAddressType{
		AddressTypeRegistered,
		AddressTypeOperating,
		AddressTypeGST,
		AddressTypeWarehouse,
		AddressTypeBranch,
	}
}

type LegalEntityAddressTypeResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}
