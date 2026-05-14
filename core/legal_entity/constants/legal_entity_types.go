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

