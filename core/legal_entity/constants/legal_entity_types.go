package legal_entity_constants

// BusinessRegistrationType represents types of business registrations in India
type LegalEntityType string

const (
	// Sole Proprietorship
	SoleProprietorship LegalEntityType = "SOLE_PROPRIETORSHIP"

	// Partnership Firm
	Partnership LegalEntityType = "PARTNERSHIP"

	// Limited Liability Partnership (LLP)
	LimitedLiabilityPartnership LegalEntityType = "LIMITED_LIABILITY_PARTNERSHIP"

	// Private Limited Company
	PrivateLimitedCompany LegalEntityType = "PRIVATE_LIMITED_COMPANY"

	// Public Limited Company
	PublicLimitedCompany LegalEntityType = "PUBLIC_LIMITED_COMPANY"

	// One Person Company (OPC)
	OnePersonCompany LegalEntityType = "ONE_PERSON_COMPANY"

	// Section 8 Company (Non-Profit Organization)
	Section8Company LegalEntityType = "SECTION_8_COMPANY"

	// Cooperative Society
	CooperativeSociety LegalEntityType = "COOPERATIVE_SOCIETY"

	// Trust (NGO / Charitable)
	Trust LegalEntityType = "TRUST"

	// Society (NGO / Charitable)
	Society LegalEntityType = "SOCIETY"
)

// Label returns a human-readable presentation string for UI components
func (t LegalEntityType) Label() string {
	switch t {
	case SoleProprietorship:
		return "Sole Proprietorship"
	case Partnership:
		return "Partnership Firm"
	case LimitedLiabilityPartnership:
		return "Limited Liability Partnership (LLP)"
	case PrivateLimitedCompany:
		return "Private Limited Company"
	case PublicLimitedCompany:
		return "Public Limited Company"
	case OnePersonCompany:
		return "One Person Company (OPC)"
	case Section8Company:
		return "Section 8 Company (Non-Profit)"
	case CooperativeSociety:
		return "Cooperative Society"
	case Trust:
		return "Trust (NGO / Charitable)"
	case Society:
		return "Society (NGO / Charitable)"
	default:
		return string(t)
	}
}

type EntityTypeResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// AllBusinessRegistrationTypes returns all supported registration types
// User for the API to return the List
func AllBusinessRegistrationTypes() []LegalEntityType {
	return []LegalEntityType{
		SoleProprietorship,
		Partnership,
		LimitedLiabilityPartnership,
		PrivateLimitedCompany,
		PublicLimitedCompany,
		OnePersonCompany,
		Section8Company,
		CooperativeSociety,
		Trust,
		Society,
	}
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityGSTRegistrationType string

// GST registration types
const (
	GSTTypeRegular     LegalEntityGSTRegistrationType = "REGULAR"
	GSTTypeComposition LegalEntityGSTRegistrationType = "COMPOSITION"
	GSTTypeSEZ         LegalEntityGSTRegistrationType = "SEZ"
	GSTTypeISD         LegalEntityGSTRegistrationType = "ISD" // Input Service Distributor
	GSTTypeCasual      LegalEntityGSTRegistrationType = "CASUAL"
)

func (t LegalEntityGSTRegistrationType) Label() string {
	switch t {
	case GSTTypeRegular:
		return "Regular Taxpayer"
	case GSTTypeComposition:
		return "Composition Scheme"
	case GSTTypeSEZ:
		return "SEZ Developer / Unit"
	case GSTTypeISD:
		return "Input Service Distributor (ISD)"
	case GSTTypeCasual:
		return "Casual Taxable Person"
	default:
		return string(t)
	}
}

// AllGSTRegistrationTypes returns all supported GST registration categories
func AllGSTRegistrationTypes() []LegalEntityGSTRegistrationType {
	return []LegalEntityGSTRegistrationType{
		GSTTypeRegular,
		GSTTypeComposition,
		GSTTypeSEZ,
		GSTTypeISD,
		GSTTypeCasual,
	}
}

type GSTRegistrationTypeResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}
