package legal_entity_constants

type LegalEntityIdentifierTypes string

const (
	IdentifierPAN   LegalEntityIdentifierTypes = "PAN"
	IdentifierCIN   LegalEntityIdentifierTypes = "CIN"
	IdentifierTAN   LegalEntityIdentifierTypes = "TAN"
	IdentifierUDYAM LegalEntityIdentifierTypes = "UDYAM"
	IdentifierIEC   LegalEntityIdentifierTypes = "IEC"   // Import Export Code
	IdentifierDPIIT LegalEntityIdentifierTypes = "DPIIT" // Startup India
	IdentifierLLPIN LegalEntityIdentifierTypes = "LLPIN" // LLP Identification Number
)

// LegalEntityIdentifierTypeResponse encapsulates identity field requirements
type LegalEntityIdentifierTypeResponse struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// Label returns the official legal name of the identity registration
func (t LegalEntityIdentifierTypes) Label() string {
	switch t {
	case IdentifierPAN:
		return "Permanent Account Number (PAN)"
	case IdentifierCIN:
		return "Corporate Identification Number (CIN)"
	case IdentifierTAN:
		return "Tax Deduction Account Number (TAN)"
	case IdentifierUDYAM:
		return "Udyam MSME Registration"
	case IdentifierIEC:
		return "Import Export Code (IEC)"
	case IdentifierDPIIT:
		return "DPIIT Startup Recognition"
	case IdentifierLLPIN:
		return "LLP Identification Number (LLPIN)"
	default:
		return string(t)
	}
}

// Description explains the regulatory purpose for frontend tooltips or context hints
func (t LegalEntityIdentifierTypes) Description() string {
	switch t {
	case IdentifierPAN:
		return "Mandatory 10-digit alphanumeric identifier issued by the Income Tax Department for direct tax compliance tracking."
	case IdentifierCIN:
		return "Unique 21-digit identifier issued by the Ministry of Corporate Affairs (MCA) upon incorporation of a company."
	case IdentifierTAN:
		return "10-digit alphanumeric number required by all businesses responsible for deducting or collecting tax at source (TDS/TCS)."
	case IdentifierUDYAM:
		return "Government registration for Micro, Small, and Medium Enterprises to qualify for priority banking and industrial benefits."
	case IdentifierIEC:
		return "Registration issued by the DGFT required for any business importing or exporting physical goods or commercial services."
	case IdentifierDPIIT:
		return "Recognition certificate issued under the Startup India initiative granting specific tax holiday exemptions."
	case IdentifierLLPIN:
		return "Unique 7-character registration key issued by the MCA to legally identify Limited Liability Partnerships."
	default:
		return ""
	}
}

// AllIdentifierTypes returns all supported legal identity types
func AllIdentifierTypes() []LegalEntityIdentifierTypes {
	return []LegalEntityIdentifierTypes{
		IdentifierPAN,
		IdentifierCIN,
		IdentifierTAN,
		IdentifierUDYAM,
		IdentifierIEC,
		IdentifierDPIIT,
		IdentifierLLPIN,
	}
}
