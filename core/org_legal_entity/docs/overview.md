

``` mermaid 
erDiagram

Organization ||--o{ OrganizationLegalEntity : owns
Organization ||--o{ OrganizationProperty : has

OrganizationProperty ||--o{ PropertyLegalEntity : mapped_to
OrganizationLegalEntity ||--o{ PropertyLegalEntity : operates

OrganizationProperty ||--o{ Hotel : contains
OrganizationProperty ||--o{ Restaurant : contains
OrganizationProperty ||--o{ Bar : contains

OrganizationLegalEntity ||--o{ LegalEntityTaxRegistration : registers
OrganizationLegalEntity ||--o{ LegalEntityTaxCategory : defines
OrganizationLegalEntity ||--o{ LegalEntityTaxRule : owns
LegalEntityTaxRule ||--o{ LegalEntityTaxComponent : splits

OrganizationLegalEntity ||--o{ AccountingBook : owns
AccountingBook ||--o{ AccountGroup : groups
AccountGroup ||--o{ Account : contains

Invoice ||--o{ InvoiceLine : has
InvoiceLine }o--|| LegalEntityTaxCategory : classified_as
Invoice ||--o{ InvoiceTaxLine : generates

Invoice ||--|| AccountingBook : posts_to
AccountingBook ||--o{ JournalEntry : records
JournalEntry ||--o{ JournalLine : contains
```