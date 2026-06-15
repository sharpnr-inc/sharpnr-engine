package legal_entity_contracts

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"
)

type UpdateLegalEntityNamesRequestDto struct {
	LegalName         *string
	DisplayName       *string
	IncorporationDate *time.Time
	FYStartMonth      *int
}

/* ---------------------------------- <***> --------------------------------- */

type UpdateLegalEntityAddressRequestDto struct {
	Type         *string
	AddressLine1 *string
	AddressLine2 *string
	Landmark     *string
	City         *string
	State        *string

	GSTStateCode *string

	Pincode *string
	Country *string
}

/* ---------------------------------- <***> --------------------------------- */

type UpdateLegalEntityBankAccountDetailsRequest struct {
	AccountHolderName *string                                            `json:"account_holder_name"`
	AccountNumber     *string                                            `json:"account_number"`
	IFSC              *string                                            `json:"ifsc"`
	BankName          *string                                            `json:"bank_name"`
	BranchName        *string                                            `json:"branch_name"`
	MICRCode          *string                                            `json:"micr_code"`
	AccountType       *legal_entity_constants.LegalEntityBankAccountType `json:"account_type"`
	IsPrimary         *bool                                              `json:"is_primary"`
	IsActive          *bool                                              `json:"is_active"`
}
