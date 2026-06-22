package sentinel_actions

import (
	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"
)

// Name is an application action identifier. Values map 1:1 to Cedar Action ids.
type Name string

const (
	ViewDocument   Name = "viewDocument"
	EditDocument   Name = "editDocument"
	DeleteDocument Name = "deleteDocument"
	ViewRepo       Name = "viewRepo"

	// Property actions.
	ViewProperty   Name = "viewProperty"
	CreateProperty Name = "createProperty"
)

// UID converts an action name to its Cedar EntityUID.
// Helper so call sites don't hand-build the EntityUID each time.
func Action(name string) types.EntityUID {
	return cedar.NewEntityUID("Action", types.String(name))
}

// this is just for example not not the part system yet.
