# Legal Entity & Tax Engine — Architecture & Design Guide

## 1. What is a Legal Entity?

A **Legal Entity** represents a legally registered business unit that:
- Owns **tax registrations**
- Is responsible for **tax compliance**
- Files **returns**
- Owns **revenue, tax liability, and accounting books**

In real life, this maps to:
- A company
- A partnership
- A proprietorship
- A registered branch

> ⚠️ **Important**  
> An **Organization** is just a business account / tenant.  
> A **Legal Entity** is what governments recognize.

---

## 2. Why Legal Entity is Needed

An organization can have:
- Multiple properties
- Multiple domains (Hotel, Restaurant, Bar)
- Multiple GST registrations
- Different tax rates per service
- Different accounting books

Example:
- One company runs:
  - Hotel under GSTIN A
  - Restaurant under GSTIN B
  - Bar under GSTIN C

All under **one Organization**, but **three Legal Entities**.

---

## 3. High-Level Ownership Model

Organization
└── LegalEntity
├── Tax Registrations
├── Tax Categories
├── Tax Rules
├── Tax Components
├── Compliance
└── Accounting Mapping



**Key Rule**  
👉 Tax logic is NEVER owned by modules (Hotel / Restaurant / Bar).  
👉 Tax logic is NEVER owned by Accounting.  
👉 Tax logic is ALWAYS owned by Legal Entity.

---

## 4. Core Legal Entity Tables

---

### 4.1 OrganizationLegalEntity

**Purpose:**  
Represents a legally registered business.

```go
OrganizationLegalEntity
