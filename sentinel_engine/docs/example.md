# Sentinel Engine — A Working Example

This page shows, in simple steps, how Cedar is wired into Sentinel and how to
use it. The code lives in:

- `models/policy.go` — how a rule is stored and turned into Cedar text
- `engine/sentinel.go` — the engine that asks Cedar the yes/no question
- `engine/sentinel_test.go` — runnable examples (allow, deny, forbid, context)

---

## The mental model

Cedar answers one question: **"Can this _principal_ do this _action_ on this
_resource_?"** It says **allow** or **deny** based on the policies you give it.

Two rules to remember:

1. **Default deny** — if no policy says `permit`, the answer is **deny**.
2. **Forbid wins** — if any policy says `forbid`, the answer is **deny**, even
   if another policy says `permit`.

---

## The three pieces you provide

To answer a question, Cedar needs three things. Sentinel gathers them for you:

| Piece | What it is | Where it comes from |
| --- | --- | --- |
| **Policies** | The rules (`permit` / `forbid`) | `PolicyStore.PoliciesFor(...)` |
| **Entities** | The principal and the roles/groups it belongs to | `PolicyStore.MembershipsOf(...)` |
| **Request** | The actual principal + action + resource being checked | The caller (`EvaluationRequest`) |

The `PolicyStore` interface is the seam between Cedar and your database. You
implement it once (against `sentinel_engine/repo`) and the engine handles the
rest.

```go
type PolicyStore interface {
    PoliciesFor(ctx, principalType, principalID string) ([]models.Policy, error)
    MembershipsOf(ctx, principalType, principalID string) ([]Membership, error)
}
```

---

## How a policy becomes Cedar text

You usually write a policy with plain fields:

```go
models.Policy{
    ID:            "engineer-read-docs",
    Effect:        models.Permit,
    PrincipalType: models.PrincipalRole,
    PrincipalID:   "engineer",
    Actions:       []string{"read", "list"},
    ResourceType:  "Document",
}
```

`ToCedar()` turns that into the Cedar language:

```cedar
// Policy: engineer-read-docs
permit (
  principal in Role::"engineer",
  action in [Action::"read", Action::"list"],
  resource is Document
);
```

> **Why `in` and not `==`?** The request's principal is always the *User*
> (e.g. `User::"alice"`). `principal in Role::"engineer"` means "alice is, or
> belongs to, the engineer role." Sentinel records alice's roles/groups as Cedar
> "parents," which is what makes role/group policies apply to a user.

If a rule is too complex for the fields (e.g. it needs a `when {}` condition),
write Cedar by hand in the `RawCedar` field and it's used as-is.

---

## Using the engine

```go
// 1. Build the engine once at startup with your repo-backed store.
sentinel := engine.Boot(myPolicyStore)

// 2. Ask a question before running business logic.
decision, err := sentinel.Evaluate(ctx, engine.EvaluationRequest{
    PrincipalType: "User",
    PrincipalID:   "alice",
    Action:        "read",
    ResourceType:  "Document",
    ResourceID:    "report-q3",
})
if err != nil {
    return err // store/parse failure — fail closed
}
if !decision.Allowed {
    return errors.New("forbidden") // decision.Reason explains why
}
// ... allowed: proceed
```

---

## What happens inside `Evaluate` (the 5 steps)

1. **Load** the principal's policies and memberships from the `PolicyStore`.
2. **Build a Cedar policy set** by joining every policy's `ToCedar()` text.
3. **Build the entity graph** — the user plus its roles/groups as parents.
4. **Build the Cedar request** — principal, action, resource, and optional context.
5. **Authorize** — Cedar returns `allow`/`deny` plus a diagnostic, which we map
   to `Decision{Allowed, Reason}`.

---

## Attribute-based rules (ABAC) with context

Pass extra facts in `Context` and reference them in a policy as `context.<key>`:

```go
sentinel.Evaluate(ctx, engine.EvaluationRequest{
    PrincipalType: "User", PrincipalID: "alice",
    Action: "read", ResourceType: "Document", ResourceID: "report-q3",
    Context: map[string]any{"mfa": true},
})
```

```cedar
permit (
  principal in Role::"engineer",
  action == Action::"read",
  resource is Document
) when { context.mfa == true };
```

(Context currently supports `string`, `bool`, and `int` values.)

---

## Run the examples

```bash
go test ./sentinel_engine/engine/ -v
```

You'll see four passing cases: a role grant, a default deny, forbid-overrides-permit,
and a context-gated (MFA) rule.

---

## What's left to wire (your next step)

The engine is complete and tested. To take it to production you implement
`PolicyStore` against your real data:

- `PoliciesFor` → load the principal's roles & groups, collect their policies, return them.
- `MembershipsOf` → return those same roles & groups as `Membership{Type, ID}`.

That work belongs in `sentinel_engine/repo` and `sentinel_engine/service`,
which currently hold the interfaces/stubs. The in-memory `memStore` in the test
file is a complete template to copy from.
```
