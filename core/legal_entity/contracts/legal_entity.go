package legal_entity_contracts

import (
	"encoding/json"
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

type NewLegalEntityRequestDto struct {
	LegalName         string                                 `json:"legal_name"`
	DisplayName       string                                 `json:"display_name"`
	Type              legal_entity_constants.LegalEntityType `json:"type"`
	IncorporationDate string                                 `json:"incorporation_date"`
	FiscalYearMonth   int                                    `json:"fiscal_year_month"`
	// Address is Optional
	Addresses []CreateLegalEntityAddressRequest `json:"addresses,omitempty"`
}

type CreateLegalEntityAddressRequest struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	Landmark     string `json:"landmark"`

	AddressType legal_entity_constants.LegalEntityAddressType `json:"address_type"`

	City string `json:"city"`

	State        string `json:"state"`
	GSTStateCode string `json:"gst_state_code"`

	Country   string `json:"country"`
	Pincode   string `json:"pincode"`
	IsPrimary bool   `json:"is_primary"`
}

/* ---------------------------------- <***> --------------------------------- */

type CreateLegalEntityContactsRequestDto struct {
	Type        legal_entity_constants.LegalEntityContactType `json:"type"`
	FirstName   string                                        `json:"first_name"`
	LastName    string                                        `json:"last_name,omitempty"`
	Designation string                                        `json:"designation,omitempty"`
	Email       string                                        `json:"email,omitempty"`
	Mobile      string                                        `json:"mobile,omitempty"`
	DIN         *string                                       `json:"din,omitempty"`
	IsPrimary   bool                                          `json:"is_primary"`
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityResponseDto struct {
	ID                uuid.UUID                              `json:"id"`
	DisplayName       string                                 `json:"display_name"`
	LegalName         string                                 `json:"legal_name"`
	Status            bool                                   `json:"status"`
	EntityType        legal_entity_constants.LegalEntityType `json:"entity_type"`
	IncorporationDate time.Time                              `json:"inc_date"`
}

type LegalEntityContactResponseDto struct {
	ID          uuid.UUID                                     `json:"id"`
	Type        legal_entity_constants.LegalEntityContactType `json:"type"`
	FirstName   string                                        `json:"first_name"`
	LastName    string                                        `json:"last_name"`
	Designation string                                        `json:"designation"`
	Email       string                                        `json:"email"`
	Mobile      string                                        `json:"mobile"`
	Din         *string                                       `json:"din"`
	IsPrimary   bool                                          `json:"is_primary"`
	IsActive    bool                                          `json:"is_active"`
	AddedOn     time.Time                                     `json:"added_on"`
}

type LegalEntityBankAccountsResponseDto struct {
	ID                uuid.UUID                                         `json:"id"`
	AccountHolderName string                                            `json:"account_holder_name"`
	AccountNumber     string                                            `json:"account_number"`
	IFSC              string                                            `json:"ifsc"`
	AccountType       legal_entity_constants.LegalEntityBankAccountType `json:"account_type"`
	IsPrimary         bool                                              `json:"is_primary"`
	IsVerified        bool                                              `json:"is_verified"`
	VerifiedAt        *time.Time                                        `json:"verified_at"`
	IsActive          bool                                              `json:"is_active"`
	AddedAt           time.Time                                         `json:"added_at"`
}

type LegalEntityAddressResponseDto struct {
	ID           uuid.UUID                                     `json:"id"`
	Type         legal_entity_constants.LegalEntityAddressType `json:"type"` // e.g., "REGISTERED", "OPERATIONAL", etc.
	AddressLine1 string                                        `json:"address_line1"`
	AddressLine2 string                                        `json:"address_line2,omitempty"` // omitempty avoids cluttering the UI if blank
	Landmark     string                                        `json:"landmark,omitempty"`
	City         string                                        `json:"city"`
	State        string                                        `json:"state"`
	GSTStateCode string                                        `json:"gst_state_code,omitempty"` // Useful for tax matching in Indian setups
	Pincode      string                                        `json:"pincode"`
	Country      string                                        `json:"country"`
	IsPrimary    bool                                          `json:"is_primary"`
	PropertyID   *uuid.UUID                                    `json:"property_id"` // Stays a pointer to allow a clean null for entity-level addresses
}

type LegalEntityIdentifierResponseDto struct {
	ID         uuid.UUID                                         `json:"id"`
	Type       legal_entity_constants.LegalEntityIdentifierTypes `json:"type"` // e.g., PAN, CIN, GSTIN
	Value      string                                            `json:"value"`
	IsVerified bool                                              `json:"is_verified"`
	VerifiedAt *time.Time                                        `json:"verified_at"` // Pointers preserve null for unverified
	ExpiresAt  *time.Time                                        `json:"expires_at"`  // null for permanent IDs
}

type LegalEntityComplianceResponseDto struct {
	ID                  uuid.UUID                                     `json:"id"`
	Type                legal_entity_constants.LegalEntityLicenseType `json:"type"` // e.g., FSSAI, Trade License
	LicenseNumber       string                                        `json:"license_number"`
	IssuingAuthority    string                                        `json:"issuing_authority,omitempty"`
	ValidFrom           time.Time                                     `json:"valid_from"`
	ValidTo             *time.Time                                    `json:"valid_to"` // null = perpetual / no expiry
	RenewalReminderDays int                                           `json:"renewal_reminder_days"`
	PropertyID          *uuid.UUID                                    `json:"property_id"`            // null = entity-level, set = location-specific
	Module              *string                                       `json:"module,omitempty"`       // e.g., HOTEL, RESTAURANT, BAR
	DocumentURL         *string                                       `json:"document_url,omitempty"` // S3 / GCS path
	IsActive            bool                                          `json:"is_active"`
}

type LegalEntityTaxRegistrationResponseDto struct {
	ID                  uuid.UUID                                             `json:"id"`
	TaxType             legal_entity_constants.LegalEntityTaxType             `json:"tax_type"` // GST | VAT
	RegistrationNumber  string                                                `json:"registration_number"`
	GSTRegistrationType legal_entity_constants.LegalEntityGSTRegistrationType `json:"gst_registration_type"`
	State               string                                                `json:"state,omitempty"`
	GSTStateCode        string                                                `json:"gst_state_code,omitempty"`
	Country             string                                                `json:"country"`
	EffectiveFrom       time.Time                                             `json:"effective_from"`
	EffectiveTo         *time.Time                                            `json:"effective_to"` // null = currently active
	IsActive            bool                                                  `json:"is_active"`
}

type LegalEntityTaxComponentResponseDto struct {
	ID                 uuid.UUID                                          `json:"id"`
	ComponentCode      legal_entity_constants.LegalEntityTaxComponentType `json:"component_code"` // CGST | SGST | IGST | etc.
	Rate               float64                                            `json:"rate"`           // e.g., 9.0
	LiabilityAccountID uuid.UUID                                          `json:"liability_account_id"`
}

type LegalEntityTaxRuleResponseDto struct {
	ID                uuid.UUID                                 `json:"id"`
	TaxRegistrationID uuid.UUID                                 `json:"tax_registration_id"` // Links rule back to the specific registration context
	TaxType           legal_entity_constants.LegalEntityTaxType `json:"tax_type"`
	Rate              float64                                   `json:"rate"`           // aggregate rate e.g. 18.0
	ConditionJSON     json.RawMessage                           `json:"condition_json"` // json.RawMessage handles arbitrary dynamic JSON without parsing overhead
	MinValue          *float64                                  `json:"min_value"`
	MaxValue          *float64                                  `json:"max_value"`
	Priority          int                                       `json:"priority"`
	EffectiveFrom     time.Time                                 `json:"effective_from"`
	EffectiveTo       *time.Time                                `json:"effective_to"`
	IsActive          bool                                      `json:"is_active"`
	Components        []LegalEntityTaxComponentResponseDto      `json:"components"` // Nested breakdown components
}

type TaxCategoryAccountMappingResponseDto struct {
	ID                  uuid.UUID `json:"id"`
	AccountingBookID    uuid.UUID `json:"accounting_book_id"`
	ReceivableAccountID uuid.UUID `json:"receivable_account_id"`
	IncomeAccountID     uuid.UUID `json:"income_account_id"`
	IsActive            bool      `json:"is_active"`
}

type LegalEntityTaxCategoryResponseDto struct {
	ID              uuid.UUID                              `json:"id"`
	Name            string                                 `json:"name"`
	Code            string                                 `json:"code"` // Stable constant key (e.g., "ROOM_RENT")
	Description     string                                 `json:"description,omitempty"`
	SACCode         *string                                `json:"sac_code,omitempty"` // null if goods-based
	HSNCode         *string                                `json:"hsn_code,omitempty"` // null if service-based
	IsActive        bool                                   `json:"is_active"`
	Rules           []LegalEntityTaxRuleResponseDto        `json:"rules"`            // Slab evaluation rules nested inside
	AccountMappings []TaxCategoryAccountMappingResponseDto `json:"account_mappings"` // Financial posting links
}

type LegalEntityDetails struct {
	BasicDetails          LegalEntityResponseDto                  `json:"basic_details"`
	Contacts              []LegalEntityContactResponseDto         `json:"contacts_and_people"`
	BankAccounts          []LegalEntityBankAccountsResponseDto    `json:"bank_accounts"`
	AddressesAndOffices   []LegalEntityAddressResponseDto         `json:"addresses"`
	LicencesAndCompliance []LegalEntityComplianceResponseDto      `json:"licences_and_compliance"`
	Identifiers           []LegalEntityIdentifierResponseDto      `json:"identifiers"`
	TaxRegistrations      []LegalEntityTaxRegistrationResponseDto `json:"tax_registrations"`
	TaxCategories         []LegalEntityTaxCategoryResponseDto     `json:"tax_categories"`
}

/* ---------------------------------- <***> --------------------------------- */
