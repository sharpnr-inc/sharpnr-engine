package sentinel_actions

// Permission catalog.
//
// This is the human-meaningful layer the UI renders. Each Permission wraps one
// Cedar action (Name) with display metadata so the React app can show a grouped
// checkbox list when an admin builds a managed policy, and so we can label the
// privileges a role/employee effectively holds.
//
// The catalog is the source of truth for "what capabilities exist". A managed
// policy stores a subset of these action Keys; we both render them (via this
// catalog) and compile them into Cedar (via the generator). Keep this list in
// sync with the action constants in core_action.go.

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

// catalog is the ordered registry. Order is the display order within a category.
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
// into display metadata.
var byKey = func() map[Name]Permission {
	m := make(map[Name]Permission, len(catalog))
	for _, p := range catalog {
		m[p.Key] = p
	}
	return m
}()

// Catalog returns the full permission catalog (the picker's data source).
func Catalog() []Permission {
	out := make([]Permission, len(catalog))
	copy(out, catalog)
	return out
}

// Lookup returns the catalog entry for an action key. ok is false for an
// unknown key (e.g. an action removed from the catalog but still referenced by
// a stored managed policy) so callers can decide how to surface it.
func Lookup(key Name) (Permission, bool) {
	p, ok := byKey[key]
	return p, ok
}

// IsValid reports whether key is a known catalog action. Use it to validate a
// managed policy's action list on write, so we never generate Cedar that
// references an action id that doesn't exist.
func IsValid(key Name) bool {
	_, ok := byKey[key]
	return ok
}
