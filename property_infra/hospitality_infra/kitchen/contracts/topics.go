package kitchen_contracts

import (
	"fmt"

	"github.com/google/uuid"
)

// KotsTopic is the SSE topic carrying a kitchen's full KOT list. Publishers
// (services) and subscribers (controllers) must both build the topic here so
// the two strings can never drift apart.
func KotsTopic(kitchenID uuid.UUID) string {
	return fmt.Sprintf("kitchen_%s", kitchenID)
}
