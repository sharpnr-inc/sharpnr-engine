package sharp_iam

import (
	"fmt"
	"strings"
)

// PolicyRequest represents what the Frontend sends to HR
type PolicyRequest struct {
	Domain       string   `json:"domain"`        // e.g., "Hospitality"
	PrincipalID  string   `json:"principal_id"`  // e.g., "emp_123"
	ResourceType string   `json:"resource_type"` // e.g., "Module"
	ResourceID   string   `json:"resource_id"`   // e.g., "Kitchen_01"
	Actions      []string `json:"actions"`       // e.g., ["CreateOrder", "EditMenu"]
}

// BuildCedarPolicy turns the UI request into a Cedar string
func BuildCedarPolicy(req PolicyRequest) string {
	// Format actions: ["A", "B"] -> Action::"A", Action::"B"
	var actionStrings []string
	for _, a := range req.Actions {
		actionStrings = append(actionStrings, fmt.Sprintf("%s::Action::\"%s\"", req.Domain, a))
	}

	actionsFormatted := strings.Join(actionStrings, ", ")

	// Build the permit statement
	// We use the Domain as a Namespace for Principal and Resource as well
	return fmt.Sprintf(
		"permit(\n  principal == %s::User::\"%s\",\n  action in [%s],\n  resource == %s::%s::\"%s\"\n);",
		req.Domain, req.PrincipalID,
		actionsFormatted,
		req.Domain, req.ResourceType, req.ResourceID,
	)
}
