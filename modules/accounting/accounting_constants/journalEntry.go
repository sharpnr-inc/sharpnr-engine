package accountingconstants

type JournalEntryType string

const (
	JournalNormal         JournalEntryType = "NORMAL"
	JournalOpeningBalance JournalEntryType = "OPENING_BALANCE"
	JournalAdjustment     JournalEntryType = "ADJUSTMENT"
	JournalReversal       JournalEntryType = "REVERSAL"
)
