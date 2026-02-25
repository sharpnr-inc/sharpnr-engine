package policies

import "embed"

//go:embed basepolicy.cedar
var InternalFS embed.FS

// GetBasePolicy reads the file from within this package
func GetBasePolicy() ([]byte, error) {
	return InternalFS.ReadFile("basepolicy.cedar")
}
