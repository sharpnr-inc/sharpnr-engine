package sentinel_actions

import "sync"

// Permission catalog.
//
// This is the human-meaningful layer the UI renders. Each Permission wraps one
// Cedar action (Name) with display metadata so the React app can show a grouped
// checkbox list when an admin builds a managed policy, and so we can label the
// privileges a role/employee effectively holds.
//
// The catalog is a registry: it ships with a built-in seed (the example domains
// below) and each feature module contributes its own actions via Register at
// startup (the composition root). That keeps action ownership with the module
// that enforces them, while the engine still validates and the FE still renders
// against one aggregated list.

// Category groups related permissions for display (the section headers in the
// permission picker). Add new categories here as new domains gain actions.
type Category string

const (
	CategoryProperties Category = "Properties"
	CategoryDocuments  Category = "Documents" // example domain; not wired into the app yet
)

// Permission is one catalog entry: a Cedar action plus how to present it.
type Permission struct {
	Key         Name     `json:"key"`         // = the Cedar action id (Name)
	Label       string   `json:"label"`       // short UI label, e.g. "Create property"
	Description string   `json:"description"` // one-line help text
	Category    Category `json:"category"`    // grouping for the UI
}

// registryMu guards catalog and byKey. Registration happens at startup; reads
// (Catalog/Lookup/IsValid) happen on the request path, so reads take RLock.
var registryMu sync.RWMutex

// catalog is the ordered registry. Order is the display order within a category.
// Property actions live here as a transitional seed; they are owned by the
// org_property module's actions package and will move there once registration is
// wired in the composition root.
var catalog = []Permission{
	// Properties
	{Key: ViewProperty, Label: "View properties", Description: "View the organization's properties.", Category: CategoryProperties},
	{Key: CreateProperty, Label: "Create property", Description: "Add a new property to the organization.", Category: CategoryProperties},
	{Key: AddBusinessModulesInProperty, Label: "Add business modules", Description: "Add business-operation modules (hotel, bar, restaurant…) to a property.", Category: CategoryProperties},

	// Documents (example domain — actions exist but are not yet enforced anywhere)
	{Key: ViewDocument, Label: "View documents", Description: "View documents.", Category: CategoryDocuments},
	{Key: EditDocument, Label: "Edit documents", Description: "Edit documents.", Category: CategoryDocuments},
	{Key: DeleteDocument, Label: "Delete documents", Description: "Delete documents.", Category: CategoryDocuments},
}

// byKey indexes the catalog for O(1) lookup when expanding stored action keys
// into display metadata. Rebuilt by Register on every contribution.
var byKey = buildIndex(catalog)

func buildIndex(perms []Permission) map[Name]Permission {
	m := make(map[Name]Permission, len(perms))
	for _, p := range perms {
		m[p.Key] = p
	}
	return m
}

// Register contributes module-owned permissions to the catalog. It is idempotent
// per Key: re-registering a Key overwrites its metadata in place (keeping display
// order) rather than duplicating it, so a module and the transitional seed can
// declare the same action without producing two picker rows. Call at startup,
// before the catalog is served.
func Register(perms ...Permission) {
	registryMu.Lock()
	defer registryMu.Unlock()
	for _, p := range perms {
		if _, exists := byKey[p.Key]; exists {
			for i := range catalog {
				if catalog[i].Key == p.Key {
					catalog[i] = p
					break
				}
			}
		} else {
			catalog = append(catalog, p)
		}
		byKey[p.Key] = p
	}
}

// Catalog returns the full permission catalog (the picker's data source).
func Catalog() []Permission {
	registryMu.RLock()
	defer registryMu.RUnlock()
	out := make([]Permission, len(catalog))
	copy(out, catalog)
	return out
}

// Lookup returns the catalog entry for an action key. ok is false for an
// unknown key (e.g. an action removed from the catalog but still referenced by
// a stored managed policy) so callers can decide how to surface it.
func Lookup(key Name) (Permission, bool) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	p, ok := byKey[key]
	return p, ok
}

// IsValid reports whether key is a known catalog action. Use it to validate a
// managed policy's action list on write, so we never generate Cedar that
// references an action id that doesn't exist.
func IsValid(key Name) bool {
	registryMu.RLock()
	defer registryMu.RUnlock()
	_, ok := byKey[key]
	return ok
}
