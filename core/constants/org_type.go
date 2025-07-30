package constants

// OrgType represents the type of an organization.
type OrgType string

const (
	Hotel      OrgType = "HOTEL"
	Retail     OrgType = "RETAIL"
	Salon      OrgType = "SALON"
	Club       OrgType = "CLUB"
	Restaurant OrgType = "RESTAURANT"
)

// IsValid checks if the organization type is valid.
func (ot OrgType) IsValid() bool {
	switch ot {
	case Hotel, Retail, Salon, Club, Restaurant:
		return true
	}
	return false
}
