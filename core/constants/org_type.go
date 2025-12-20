package constants

// OrgType represents the type of an organization.
type OrgType string

const (
	Hospitality OrgType = "HOSPITALITY"
	Retail      OrgType = "RETAIL"
	Salon       OrgType = "SALON"
)

// IsValid checks if the organization type is valid.
func (ot OrgType) IsValid() bool {
	switch ot {
	case Hospitality, Retail, Salon:
		return true
	}
	return false
}
