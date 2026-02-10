# Accounting 101 for Developers

## 1. The Golden Rule: The Accounting Equation
Every transaction must keep this equation in balance. If it doesn't, your "Journal Entry" is invalid.

$$Assets = Liabilities + Equity$$

* **Assets:** What you **own** (Cash, Inventory, Equipment).
* **Liabilities:** What you **owe** (Loans, Unpaid Bills).
* **Equity:** What is **left over** for the owners.

---

## 2. The DEAD COIL Cheat Sheet
Use this to determine if a value goes in the **Debit** (left) or **Credit** (right) column.

| Type | Account | Increase | Decrease |
| :--- | :--- | :--- | :--- |
| **D** | **D**ebit | — | — |
| **E** | **E**xpenses | Debit | Credit |
| **A** | **A**ssets | Debit | Credit |
| **D** | **D**rawings | Debit | Credit |
| | | | |
| **C** | **C**redit | — | — |
| **O** | **O**wner's Equity| Credit | Debit |
| **I** | **I**ncome (Revenue)| Credit | Debit |
| **L** | **L**iabilities | Credit | Debit |

---

## 3. The Anatomy of a Transaction
In database terms, a transaction consists of one **Header** (Journal Entry) and at least two **Lines** (Journal Lines).

### Rules for Journal Lines:
1. **Double Entry:** You cannot have a single line. You need at least one Debit and one Credit.
2. **Zero Sum:** The sum of all Debits minus the sum of all Credits must equal zero.
   $$\sum Debits - \sum Credits = 0$$



---

## 4. Common Developer Pitfalls
* **Data Types:** Never use `float` for currency. Use `decimal`, `numeric`, or `bigint` (storing cents) to avoid rounding errors.
* **Immutability:** Once a Journal Entry is "Posted," it should generally not be deleted. To fix an error, you issue a "Reversing Entry."
* **The "Sign" Convention:** * In the DB, many devs store everything as a positive number and use a `side` column (`DEBIT` or `CREDIT`).
    * Alternatively, some store Debits as positive and Credits as negative. 

---

## 5. Typical Flow
1. **Transaction:** Customer buys a coffee for $5.
2. **Journal Entry:** * Debit **Cash** (Asset increases) $5.00
   * Credit **Sales Revenue** (Income increases) $5.00
3. **Ledger:** These lines are "posted" to their respective accounts to update the balance.