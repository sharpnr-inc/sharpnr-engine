package sharpnruirenderer

import (
	"context"
	"fmt"
	"sharpnr-engine/core/constants"
	"sort"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func (o *orgUIRendererRepository) CreateOrgUIRenderer(ctx context.Context, orgId uuid.UUID, orgConfig []string, orgType constants.OrgType) (datatypes.JSON, error) {
	fmt.Println("Incoming orgConfig:", orgConfig)

	sort.Strings(orgConfig)

	// Join slice to create a unique key
	key := ""
	for i, v := range orgConfig {
		if i > 0 {
			key += "_"
		}
		key += v
	}

	fmt.Println("Combination key:", key)

	var (
		uiJSON datatypes.JSON
		err    error
	)

	// --- Switch based on combination ---
	switch key {

	case "BAR":
		fmt.Println("Only BAR module active")

		uiJSON, err = BarRendererFactory(ctx, orgId, orgType)
		fmt.Println(uiJSON)

	case "HOTEL":
		fmt.Println("Only HOTEL module active")
		// TODO: generate HOTEL-only renderer JSON

	case "RESTAURANT":
		fmt.Println("Only RESTAURANT module active")
		// TODO: generate RESTAURANT-only renderer JSON

	case "BAR_HOTEL":
		fmt.Println("BAR + HOTEL")
		// TODO: generate BAR + HOTEL renderer JSON

	case "BAR_RESTAURANT":
		fmt.Println("BAR + RESTAURANT")
		// TODO: generate BAR + RESTAURANT renderer JSON

	case "HOTEL_RESTAURANT":
		fmt.Println("HOTEL + RESTAURANT")
		// TODO: generate HOTEL + RESTAURANT renderer JSON

	case "BAR_HOTEL_RESTAURANT":
		fmt.Println("All three modules active")
		// TODO: generate full renderer JSON

	default:
		fmt.Println("Unknown combination or empty orgConfig")
		return nil, fmt.Errorf("invalid org configuration")
	}

	// Placeholder empty JSON (replace with actual render tree)
	if err != nil {
		return nil, err
	}

	// ✅ Return the actual render tree, not placeholder
	return uiJSON, nil
}
