# Database Transaction Management (txManager) Guide

## 1. Concept Overview
The `txManager` is a structural pattern used to ensure **ACID** properties (Atomicity, Consistency, Isolation, Durability) within our Go services. 

In accounting, we often need to perform multiple inserts (Book -> Fiscal Year -> Groups -> Accounts). If any step fails, we must revert all previous steps to prevent data corruption. The `txManager` facilitates this by carrying a database transaction through the `context.Context`.

---

## 2. Architecture & Data Flow
The following diagram illustrates how the transaction is initialized in the **Service**, carried by the **Context**, and utilized by the **Repository**.



```mermaid
sequenceDiagram
    participant S as Accounting Service
    participant TX as txManager (GormTxManager)
    participant C as Context (txCtx)
    participant R as Accounting Repository
    participant DB as Database (Postgres)

    Note over S, DB: Start of Atomic Operation
    S->>TX: Call WithTx(ctx, func)
    TX->>DB: BEGIN TRANSACTION
    DB-->>TX: tx_instance (*gorm.DB)
    
    TX->>C: dbtx.InjectTx(ctx, tx_instance) 
    Note right of C: Transaction session stored in Context
    
    S->>R: CreateAccountingBook(txCtx, data)
    R->>C: dbtx.ExtractTx(txCtx)
    C-->>R: Returns tx_instance
    R->>DB: INSERT (Transaction Session)
    
    S->>R: CreateFiscalYear(txCtx, data)
    R->>C: dbtx.ExtractTx(txCtx)
    R->>DB: INSERT (Transaction Session)

    alt Success
        S-->>TX: Return nil
        TX->>DB: COMMIT
        Note over DB: All changes saved permanently
    else Failure
        S-->>TX: Return error
        TX->>DB: ROLLBACK
        Note over DB: All changes erased (Atomic Revert)
    end