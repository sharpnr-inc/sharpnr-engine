# SERE Workflow Engine & Transaction Management Guide

## 1. Introduction: The "Atomic" Problem
In a complex system (like an ERP or Hospitality app), a single user action often triggers multiple steps across different modules.

**Example: "Onboarding a New Organization"**
1.  Create the **Organization** record in the database.
2.  Create the **Accounting Book** & Ledgers.
3.  Create the **Admin User**.

**The Risk:** If Step 1 succeeds but Step 2 fails, you have a "Zombie Organization" with no accounting ledgers. The system is broken.

**The Solution:** We use **SERE (Workflow Engine)** combined with a **txManager (Transaction Manager)**. This ensures that **all steps** happen inside a single "Safety Bubble" (Database Transaction). If *any* step fails, *everything* is erased (Rolled Back).

---

## 2. How It Works (The Architecture)

We don't want our Event Engine (`SERE`) to depend on our Database logic (`GORM`). That would be "messy code." Instead, we use a **Bridge Pattern**.

1.  **The Bridge (`UseTxExecutor`):** We teach SERE *how* to start a transaction, but we keep the actual database code outside of SERE.
2.  **The Carrier (`Context`):** The database transaction is hidden inside the Go `context.Context`. We pass this context from handler to handler.
3.  **The Workflow:** A list of functions (Handlers) that run in order. They share data and the transaction.

### Visualizing the Flow



```mermaid
graph TD
    subgraph "Orchestrator Service"
        Start[Start Onboarding] --> CallSERE[Call SERE.ExecuteTx]
    end

    subgraph "Infrastructure (The Bridge)"
        CallSERE --> TxMgr[txManager: Start DB Transaction]
        TxMgr --> Inject[Inject Transaction into Context]
        Inject --> SERE_Loop[SERE: Run Handlers Loop]
    end

    subgraph "The Workflow (Inside Transaction)"
        SERE_Loop --> H1[Handler 1: Create Org]
        H1 -->|Writes to DB| DB[(Database)]
        H1 -->|Passes Data| H2[Handler 2: Create Accounting]
        H2 -->|Writes to DB| DB
    end

    H2 --> Decision{Success?}
    Decision -- Yes --> Commit[Commit Transaction (Save All)]
    Decision -- No --> Rollback[Rollback Transaction (Undo All)]

sequenceDiagram
    participant Orch as Orchestrator
    participant SERE as SERE Engine
    participant TX as txManager
    participant H as Handlers (Org/Accounting)
    participant DB as Database

    Orch->>SERE: ExecuteTx(ctx, "ONBOARD_ORG", data)
    
    Note over SERE, TX: SERE uses the "Bridge" function
    SERE->>TX: txManager.WithTx(ctx)
    TX->>DB: BEGIN TRANSACTION
    TX-->>SERE: Context with Tx Injected (txCtx)

    loop Every Handler
        SERE->>H: Run Handler(wf) using txCtx
        H->>DB: Insert Data (using txCtx)
        
        alt Handler Fails
            H-->>SERE: Return Error
            SERE-->>TX: Return Error
            TX->>DB: ROLLBACK (Everything is deleted)
            TX-->>Orch: Return Error
        end
    end

    SERE-->>TX: Return Nil (Success)
    TX->>DB: COMMIT (Data is permanent)
    TX-->>Orch: Success