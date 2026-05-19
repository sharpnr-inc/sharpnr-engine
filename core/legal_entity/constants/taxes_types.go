package legal_entity_constants

type LegalEntityTaxType string

// Tax types
const (
	TaxTypeGST LegalEntityTaxType = "GST"
	TaxTypeVAT LegalEntityTaxType = "VAT"
)

type LegalEntityTaxComponentType string

// Tax component codes
const (
	ComponentCGST LegalEntityTaxComponentType = "CGST"
	ComponentSGST LegalEntityTaxComponentType = "SGST"
	ComponentIGST LegalEntityTaxComponentType = "IGST"
	ComponentUGST LegalEntityTaxComponentType = "UGST"
	ComponentCESS LegalEntityTaxComponentType = "CESS"
	ComponentVAT  LegalEntityTaxComponentType = "VAT"
)

type LegalEntityBankAccountType string

// Bank account types
const (
	AccountTypeCurrent LegalEntityBankAccountType = "CURRENT"
	AccountTypeSavings LegalEntityBankAccountType = "SAVINGS"
	AccountTypeODCC    LegalEntityBankAccountType = "ODCC"
)
