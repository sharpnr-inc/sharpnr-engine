package legal_entity_constants

type LegalEntityLicenseType string

// License / compliance types
const (
	LicenseFSSAI     LegalEntityLicenseType = "FSSAI"
	LicenseLiquor    LegalEntityLicenseType = "LIQUOR"
	LicenseFire      LegalEntityLicenseType = "FIRE"
	LicenseHealth    LegalEntityLicenseType = "HEALTH"
	LicenseTrade     LegalEntityLicenseType = "TRADE"
	LicenseShopEstab LegalEntityLicenseType = "SHOP_ESTAB"
	LicensePCB       LegalEntityLicenseType = "PCB" // Pollution Control Board
	LicenseDrug      LegalEntityLicenseType = "DRUG"
	LicensePF        LegalEntityLicenseType = "PF" // Provident Fund
	LicenseESIC      LegalEntityLicenseType = "ESIC"
	LicensePT        LegalEntityLicenseType = "PT" // Professional Tax
)

type LegalEntityLicenseTypeResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Label returns a human-readable presentation string for licenses
func (t LegalEntityLicenseType) Label() string {
	switch t {
	case LicenseFSSAI:
		return "FSSAI (Food Safety License)"
	case LicenseLiquor:
		return "Liquor License"
	case LicenseFire:
		return "Fire Safety NOC"
	case LicenseHealth:
		return "Health Trade License"
	case LicenseTrade:
		return "Trade License"
	case LicenseShopEstab:
		return "Shops & Establishment Registration"
	case LicensePCB:
		return "Pollution Control Board NOC"
	case LicenseDrug:
		return "Drug License"
	case LicensePF:
		return "Provident Fund (PF) Registration"
	case LicenseESIC:
		return "ESIC Registration"
	case LicensePT:
		return "Professional Tax (PT) Registration"
	default:
		return string(t)
	}
}

// AllLicenseTypes returns all supported license categories
func AllLicenseTypes() []LegalEntityLicenseType {
	return []LegalEntityLicenseType{
		LicenseFSSAI,
		LicenseLiquor,
		LicenseFire,
		LicenseHealth,
		LicenseTrade,
		LicenseShopEstab,
		LicensePCB,
		LicenseDrug,
		LicensePF,
		LicenseESIC,
		LicensePT,
	}
}

/* ----------------------------- Contact Types ----------------------------- */

type LegalEntityContactType string

const (
	ContactTypeDirector   LegalEntityContactType = "DIRECTOR"
	ContactTypeAuthorized LegalEntityContactType = "AUTHORIZED_SIGNATORY"
	ContactTypeAccounts   LegalEntityContactType = "ACCOUNTS"
	ContactTypeCompliance LegalEntityContactType = "COMPLIANCE"
	ContactTypeOperations LegalEntityContactType = "OPERATIONS"
)

type LegalEntityContactTypeResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Label returns a human-readable presentation string for contacts
func (t LegalEntityContactType) Label() string {
	switch t {
	case ContactTypeDirector:
		return "Director / Partner"
	case ContactTypeAuthorized:
		return "Authorized Signatory"
	case ContactTypeAccounts:
		return "Accounts / Finance Team"
	case ContactTypeCompliance:
		return "Compliance / Legal Officer"
	case ContactTypeOperations:
		return "Operations Manager"
	default:
		return string(t)
	}
}

// AllContactTypes returns all supported contact types
func AllContactTypes() []LegalEntityContactType {
	return []LegalEntityContactType{
		ContactTypeDirector,
		ContactTypeAuthorized,
		ContactTypeAccounts,
		ContactTypeCompliance,
		ContactTypeOperations,
	}
}
