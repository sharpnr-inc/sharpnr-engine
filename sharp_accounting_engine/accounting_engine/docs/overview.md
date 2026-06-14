Organization
└── LegalEntity
└── AccountingBook
├── FiscalYear
│ └── JournalEntry
│ └── JournalLine
│ └── Account
│
└── AccountGroup
└── Account

🔴 Important clarification

AccountGroup does NOT live under FiscalYear

Accounts are NOT recreated per year

Fiscal year only affects balances, not structure

ASSET ✅ money lives here — Cash, Bank, etc.
LIABILITY ❌ this is what you OWE — Accounts Payable, Loans
EXPENSE ❌ this records costs — Rent, Salaries
EQUITY ❌ this is owner's stake — Capital, Retained Earnings
INCOME ❌ this records revenue — Sales, Interest earned
