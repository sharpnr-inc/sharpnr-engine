package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// =============================================================================
// BANK ACCOUNTS
// =============================================================================

// LegalEntityBankAccount stores banking details used for settlements,
// vendor payments, NACH mandates, and TDS/TCS filings.
//
// AccountNumber should be encrypted at rest — store only the encrypted
// ciphertext here. The service layer decrypts on demand.
// IsVerified flags that a penny-drop or NACH verification succeeded.
type LegalEntityBankAccount struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`

	AccountHolderName string `gorm:"size:355;not null"`
	AccountNumber     string `gorm:"size:512;not null"` // encrypted ciphertext
	IFSC              string `gorm:"size:11;not null"`  // always 11 chars
	BankName          string `gorm:"size:100"`
	BranchName        string `gorm:"size:100"`
	MICRCode          string `gorm:"size:9"` // 9-digit MICR for cheque processing

	AccountType string `gorm:"size:15;not null;default:'CURRENT'"` // see AccountType constants

	// IsPrimary marks the default account for outgoing payments.
	IsPrimary bool `gorm:"default:false;not null"`

	IsVerified bool `gorm:"default:false;not null"`
	VerifiedAt *time.Time

	IsActive  bool `gorm:"default:true;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LegalEntityBankAccount) TableName() string {
	return "legal_entity.legal_entity_bank_account"
}
