package legal_entity_constants

// BusinessRegistrationType represents types of business registrations in India
type BusinessRegistrationType string

const (
	// Sole Proprietorship
	SoleProprietorship BusinessRegistrationType = "SOLE_PROPRIETORSHIP"

	// Partnership Firm
	Partnership BusinessRegistrationType = "PARTNERSHIP"

	// Limited Liability Partnership (LLP)
	LimitedLiabilityPartnership BusinessRegistrationType = "LIMITED_LIABILITY_PARTNERSHIP"

	// Private Limited Company
	PrivateLimitedCompany BusinessRegistrationType = "PRIVATE_LIMITED_COMPANY"

	// Public Limited Company
	PublicLimitedCompany BusinessRegistrationType = "PUBLIC_LIMITED_COMPANY"

	// One Person Company (OPC)
	OnePersonCompany BusinessRegistrationType = "ONE_PERSON_COMPANY"

	// Section 8 Company (Non-Profit Organization)
	Section8Company BusinessRegistrationType = "SECTION_8_COMPANY"

	// Cooperative Society
	CooperativeSociety BusinessRegistrationType = "COOPERATIVE_SOCIETY"

	// Trust (NGO / Charitable)
	Trust BusinessRegistrationType = "TRUST"

	// Society (NGO / Charitable)
	Society BusinessRegistrationType = "SOCIETY"
)

// AllBusinessRegistrationTypes returns all supported registration types
func AllBusinessRegistrationTypes() []BusinessRegistrationType {
	return []BusinessRegistrationType{
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
