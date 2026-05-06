# Sentinel Engine — Introduction

> Version: `v0.0.1` · Module: `github.com/sharpnr/sentinel_engine`

---

## What is Sentinel Engine?

Sentinel Engine is SharpNR's internal authorization engine. It is injected into all SharpNR service modules and is responsible for one thing: **deciding whether a principal (user, group, or role) is allowed to perform an action on a resource.**

It is conceptually similar to AWS IAM — you define policies, attach them to principals, and the engine evaluates them at request time. Under the hood it uses [Cedar](https://www.cedarpolicy.com/), Amazon's open-source policy language and evaluation engine, via the [`cedar-go`](https://github.com/cedar-policy/cedar-go) bindings.

---

## Why Cedar?

Cedar was built specifically for authorization. It is:

- **Formally verified** — the evaluation algorithm has mathematical correctness guarantees
- **Extremely fast** — policy evaluation is near-instant even with large resource trees
- **Expressive** — supports attribute-based access control (ABAC) alongside role-based (RBAC)
- **Auditable** — every decision includes a diagnostic trace of which policy matched

---

## Core Concepts

| Concept       | Description                                                                     |
| ------------- | ------------------------------------------------------------------------------- |
| **Principal** | The entity requesting access — a `User`, `Group`, or `Role`                     |
| **Action**    | The operation being attempted — e.g. `read`, `write`, `delete`                  |
| **Resource**  | The target of the action — e.g. `Document::"report-q3"`                         |
| **Policy**    | A rule that permits or forbids a principal from performing actions on resources |
| **Context**   | Optional key-value data passed at evaluation time for dynamic conditions        |

---

## How Sentinel Fits Into SharpNR

Every SharpNR service that needs access control receives a `SentinelService` instance at startup via constructor injection. When a request comes in, the service calls `sentinel.Evaluate(...)` before executing any business logic.

```
HTTP Request
     │
     ▼
SharpNR Service
     │
     ├── sentinel.Evaluate(req) ──► Sentinel Engine ──► Cedar Evaluator
     │         │                                              │
     │    Allowed / Denied ◄────────────────────────── Policy Match
     │
     ▼
Business Logic (only if allowed)
```

Sentinel is **not** an HTTP service itself — it is a Go package that runs in-process with each service. There is no extra network hop.

---

## Quick Start

```go
import (
    "github.com/sharpnr/sentinel_engine/engine"
    "github.com/sharpnr/sentinel_engine/service"
)

// Boot the engine (do this once at service startup)
sentinel := engine.Boot()

// Evaluate a permission
result, err := sentinel.Evaluate(ctx, service.EvaluationRequest{
    PrincipalType: "User",
    PrincipalID:   "alice",
    Action:        "read",
    ResourceType:  "Document",
    ResourceID:    "report-q3",
})

if !result.Allowed {
    return errors.New("forbidden")
}
```

---

## Related Docs

- [How the Engine Works](./engine.md)
- [Models Reference](./models.md)
- [Service & Injection Guide](./service.md)
