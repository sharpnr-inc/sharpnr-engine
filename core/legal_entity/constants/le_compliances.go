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

type LegalEntityContactType string

// Contact types
const (
	ContactTypeDirector   LegalEntityContactType = "DIRECTOR"
	ContactTypeAuthorized LegalEntityContactType = "AUTHORIZED_SIGNATORY"
	ContactTypeAccounts   LegalEntityContactType = "ACCOUNTS"
	ContactTypeCompliance LegalEntityContactType = "COMPLIANCE"
	ContactTypeOperations LegalEntityContactType = "OPERATIONS"
)
