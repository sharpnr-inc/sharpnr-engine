package legal_entity_constants

type LegalEntityTaxType string

// Tax types
const (
	TaxTypeGST LegalEntityTaxType = "GST"
	TaxTypeVAT LegalEntityTaxType = "VAT"
)

type LegalEntityTaxTypeResponse struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// Label returns the official format name of the tax system
func (t LegalEntityTaxType) Label() string {
	switch t {
	case TaxTypeGST:
		return "Goods and Services Tax (GST)"
	case TaxTypeVAT:
		return "Value Added Tax (VAT)"
	default:
		return string(t)
	}
}

// Description explains the tax framework for frontend tooltips or context hints
func (t LegalEntityTaxType) Description() string {
	switch t {
	case TaxTypeGST:
		return "Comprehensive, multi-stage, destination-based indirect tax system levied on the supply of goods and services."
	case TaxTypeVAT:
		return "Consumption tax assessed on the value added in each production and distribution stage of goods and services."
	default:
		return ""
	}
}

// AllTaxTypes returns all supported legal tax types
func AllTaxTypes() []LegalEntityTaxType {
	return []LegalEntityTaxType{
		TaxTypeGST,
		TaxTypeVAT,
	}
}

/* --------------------------------- <*****> -------------------------------- */

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

type LegalEntityTaxComponentTypeResponse struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// Label returns the standard legal name of the tax component
func (c LegalEntityTaxComponentType) Label() string {
	switch c {
	case ComponentCGST:
		return "Central Goods and Services Tax (CGST)"
	case ComponentSGST:
		return "State Goods and Services Tax (SGST)"
	case ComponentIGST:
		return "Integrated Goods and Services Tax (IGST)"
	case ComponentUGST:
		return "Union Territory Goods and Services Tax (UTGST)"
	case ComponentCESS:
		return "Compensation Cess"
	case ComponentVAT:
		return "Value Added Tax Component"
	default:
		return string(c)
	}
}

// Description explains the destination or allocation of the tax component breakdown
func (c LegalEntityTaxComponentType) Description() string {
	switch c {
	case ComponentCGST:
		return "The central government share of tax collected on intra-state supplies of goods or services."
	case ComponentSGST:
		return "The state government share of tax collected on intra-state supplies of goods or services."
	case ComponentIGST:
		return "Tax levied on all inter-state supplies of goods or services, collected by the center and shared with consumption states."
	case ComponentUGST:
		return "Tax levied on supplies of goods or services within Union Territories without an independent legislature."
	case ComponentCESS:
		return "Additional levy over standard tax rates applied primarily on luxury, demerit, or sin goods to compensate states for revenue loss."
	case ComponentVAT:
		return "The direct single-tier provincial or state component evaluated during transactions under a VAT regime."
	default:
		return ""
	}
}

// AllTaxComponentTypes returns all supported tax components
func AllTaxComponentTypes() []LegalEntityTaxComponentType {
	return []LegalEntityTaxComponentType{
		ComponentCGST,
		ComponentSGST,
		ComponentIGST,
		ComponentUGST,
		ComponentCESS,
		ComponentVAT,
	}
}

type LegalEntityBankAccountType string

// Bank account types
const (
	AccountTypeCurrent LegalEntityBankAccountType = "CURRENT"
	AccountTypeSavings LegalEntityBankAccountType = "SAVINGS"
	AccountTypeODCC    LegalEntityBankAccountType = "ODCC"
)

type LegalEntityBankAccountTypeResponse struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// Label returns the commercial naming for the bank account classification
func (b LegalEntityBankAccountType) Label() string {
	switch b {
	case AccountTypeCurrent:
		return "Current Account"
	case AccountTypeSavings:
		return "Savings Account"
	case AccountTypeODCC:
		return "Overdraft / Cash Credit (OD/CC)"
	default:
		return string(b)
	}
}

// Description explains the banking operational parameters for business use cases
func (b LegalEntityBankAccountType) Description() string {
	switch b {
	case AccountTypeCurrent:
		return "High-volume transactional account tailored for business entities with no limits on deposits or withdrawals."
	case AccountTypeSavings:
		return "Interest-bearing deposit account meant primarily for retail, individual transactions, or proprietary businesses with limits on free transactional volume."
	case AccountTypeODCC:
		return "Credit facilities mapping either an overdraft ceiling or cash credit line against business inventory/receivables to bridge working capital."
	default:
		return ""
	}
}

// AllBankAccountTypes returns all supported business banking account options
func AllBankAccountTypes() []LegalEntityBankAccountType {
	return []LegalEntityBankAccountType{
		AccountTypeCurrent,
		AccountTypeSavings,
		AccountTypeODCC,
	}
}
