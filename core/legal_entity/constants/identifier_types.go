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

