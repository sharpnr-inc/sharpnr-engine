package accountingconstants

type JournalEntryType string

const (
	JournalNormal         JournalEntryType = "NORMAL"
	JournalOpeningBalance JournalEntryType = "OPENING_BALANCE"
	JournalAdjustment     JournalEntryType = "ADJUSTMENT"
	JournalReversal       JournalEntryType = "REVERSAL"
)

func (j JournalEntryType) IsValid() bool {
	switch j {
	case JournalNormal,
		JournalOpeningBalance,
		JournalAdjustment,
		JournalReversal:
		return true
	default:
		return false
	}
}
