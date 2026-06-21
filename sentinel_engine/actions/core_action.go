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
)

// UID converts an action name to its Cedar EntityUID.
func (n Name) UID() types.EntityUID {
	return cedar.NewEntityUID("Action", types.String(n))
}

// this is just for example not not the part system yet.
