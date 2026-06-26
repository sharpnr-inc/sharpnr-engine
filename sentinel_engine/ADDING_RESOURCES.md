# Adding a Resource, its Actions and Attributes

This guide shows how to add a new **resource** to the Sentinel policy builder so
it shows up in the frontend. Keep it simple: copy the `Property` example.

## The idea in one minute

The policy builder is built around **resources**. A resource is a thing a policy
can be about — a `Property`, an `Invoice`, an `Order`, etc. (the Cedar entity type).

Each resource has three things:

1. **Actions** — what you can do to it (`createProperty`, `editProperty`…).
2. **Attributes** — fields a policy can test in a `when {}` clause (`resource.country`…).
3. **Instances** — the real rows in the org (each `Property::"<id>"`).

The frontend asks for these in steps, so nothing is ever a giant payload:

| The FE calls | It gets |
|---|---|
| `GET /sentinel/resources` | the list of resource types (just type + label + description) |
| `GET /sentinel/resources/:type/actions` | that type's actions |
| `GET /sentinel/resources/:type/attributes` | that type's attributes |
| `GET /sentinel/resources/:type/instances` | the org's live rows of that type |

You add a resource by declaring those three things in your module and registering
them once at startup. Three steps.

---

## Step 1 — Declare the actions

In your module, make an `actions` package (e.g. `core/org_invoice/actions/invoice.go`).

```go
package orgInvoiceSentinelActions

import (
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

// Action ids are STABLE strings. Once a saved policy uses one, never rename it —
// that silently breaks the policy. Use camelCase.
const (
	ViewInvoice   sentinel_actions.Name = "viewInvoice"
	CreateInvoice sentinel_actions.Name = "createInvoice"
	VoidInvoice   sentinel_actions.Name = "voidInvoice"
)

// Actions is what the resource descriptor (Step 2) will use.
var Actions = []sentinel_resources.Action{
	{Key: ViewInvoice, Label: "View invoices", Description: "View the org's invoices."},
	{Key: CreateInvoice, Label: "Create invoice", Description: "Create a new invoice."},
	{Key: VoidInvoice, Label: "Void invoice", Description: "Void an existing invoice."},
}
```

(Real example: [core/org_property/actions/property.go](../core/org_property/actions/property.go))

---

## Step 2 — Declare the resource (descriptor + instance provider)

In your module, make a `resources` package. It holds two things.

### 2a. The Descriptor (actions + attributes)

```go
package orgInvoiceResources

import (
	orgInvoiceSentinelActions "sharpnr-engine/core/org_invoice/actions"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"
)

// ResourceType is the Cedar entity type. It MUST match the type used when you
// enforce permissions and the type written in policy text (e.g. resource is Invoice).
const ResourceType = "Invoice"

var Descriptor = sentinel_resources.Descriptor{
	Type:        ResourceType,
	Label:       "Invoice",
	Description: "An organization's invoices.",
	Actions:     orgInvoiceSentinelActions.Actions,
	Attributes: []sentinel_resources.Attribute{
		// Type is the Cedar value type: "String" | "Long" | "Boolean" | "Set" | "EntityRef".
		{Name: "status", Label: "Status", Type: "String", Description: "Invoice status (draft, paid…)."},
		{Name: "amount", Label: "Amount", Type: "Long", Description: "Invoice total."},
	},
}
```

### 2b. The Provider (lists live instances)

This fetches the real rows for one org. It uses your module's **service**, not the repo.

```go
package orgInvoiceResources

import (
	"context"

	org_invoice_services "sharpnr-engine/core/org_invoice/services"
	sentinel_resources "sharpnr-engine/sentinel_engine/resources"

	"github.com/google/uuid"
)

type invoiceResources struct {
	svc org_invoice_services.InvoiceServiceInterface
}

func NewInvoiceResources(svc org_invoice_services.InvoiceServiceInterface) *invoiceResources {
	return &invoiceResources{svc: svc}
}

func (p *invoiceResources) Type() string { return ResourceType }

func (p *invoiceResources) List(ctx context.Context, orgID uuid.UUID) ([]sentinel_resources.Ref, error) {
	invoices, err := p.svc.ListInvoices(ctx, orgID)
	if err != nil {
		return nil, err
	}
	refs := make([]sentinel_resources.Ref, 0, len(invoices))
	for _, inv := range invoices {
		refs = append(refs, sentinel_resources.Ref{ID: inv.ID, Type: ResourceType, Label: inv.Number})
	}
	return refs, nil
}

var _ sentinel_resources.Provider = (*invoiceResources)(nil)
```

(Real example: [core/org_property/resources/descriptor.go](../core/org_property/resources/descriptor.go) and
[constructor.go](../core/org_property/resources/constructor.go))

---

## Step 3 — Register it at startup

Open [app/bootstrap_sentinel_catalog.go](../app/bootstrap_sentinel_catalog.go) and add
two lines. Pass in your module's service so the provider can list instances.

```go
func registerSentinelCatalog(
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	invoiceService org_invoice_services.InvoiceServiceInterface, // <- add the service
) {
	// Property
	sentinel_resources.RegisterDescriptor(orgPropertyresources.Descriptor)
	sentinel_resources.RegisterProvider(orgPropertyresources.NewOrgPropertyResources(orgPropertyService))

	// Invoice  <- your new resource
	sentinel_resources.RegisterDescriptor(orgInvoiceResources.Descriptor)
	sentinel_resources.RegisterProvider(orgInvoiceResources.NewInvoiceResources(invoiceService))
}
```

Then update the one call in [app/bootstrap.go](../app/bootstrap.go) to pass the new
service. That's it — the four endpoints now serve your resource.

---

## Rules to remember

- **Action ids are forever.** Don't rename them once a policy uses them. Add new
  ones instead.
- **`Type` must match everywhere.** The descriptor `Type`, the policy text
  (`resource is Invoice`), and the value you pass when you enforce a permission
  must be the same string.
- **Registering a descriptor also registers its actions** as valid with the
  engine — you don't do that separately.
- **Attributes need one more wire to actually work.** Listing an attribute here
  only makes the FE show it. For a `when { resource.status == "paid" }` clause to
  evaluate, the code that calls the engine must also put `status` onto the
  resource (`AuthorizeInput.Resource.Attrs`). Add that when you start using the
  attribute in real policies.
- **Org-level actions** (things not tied to a record, like "edit org name") fit as
  an `"Organization"` resource — same three steps.
