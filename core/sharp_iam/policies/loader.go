package policies

import "embed"

//go:embed requirements.cedar
var InternalFS embed.FS

// GetBasePolicy reads the file from within this package
func GetBasePolicy() ([]byte, error) {
	return InternalFS.ReadFile("requirements.cedar")
}
