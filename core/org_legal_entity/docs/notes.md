So this is legal entity of the organization, it can also be referred as the Tax Engine in sharpnr, the module decide what tax to applied on what product, on service

Organization
 └── Legal Entity
       ├── Tax Registrations
       ├── Tax Categories
       ├── Tax Rules
       └── Compliance


Order / Charge
   ↓
InvoiceEngine
   ↓
Tax Engine (Legal Entity)
   ↓
Tax Breakdown (CGST, SGST, VAT, etc.)
   ↓
Accounting Engine
   ↓
Ledger Entries


Organization
└── LegalEntity
├── Tax Registrations
├── Tax Categories
├── Tax Rules
├── Tax Components
├── Compliance
└── Accounting Mapping