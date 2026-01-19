package coreutilitiesapi

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

type SubModule struct {
	ID          string `json:"id"`
	ModuleName  string `json:"module_name"`
	Description string `json:"description"`
}

func GetSubModules(c *gin.Context) {
	module := c.Param("module")

	var data []SubModule

	switch module {
	case "hospitality":
		data = []SubModule{
			{
				ID:          "1",
				ModuleName:  "HOTEL",
				Description: "Comprehensive hotel management module for handling reservations, room inventory, guest check-ins/outs, housekeeping, billing, and reporting.",
			},
			{
				ID:          "2",
				ModuleName:  "RESTAURANT",
				Description: "End-to-end restaurant operations module for managing tables, menus, orders, kitchen display systems (KDS), billing, and online ordering integration.",
			},
			{
				ID:          "3",
				ModuleName:  "PUB/CLUB",
				Description: "Entertainment and nightlife operations module covering guest entry, table and bottle service, membership management, event scheduling, and POS integration.",
			},
			{
				ID:          "4",
				ModuleName:  "BAR",
				Description: "Bar and lounge management system focusing on inventory tracking, liquor control, bartender sales monitoring, and daily revenue insights.",
			},
			{
				ID:          "5",
				ModuleName:  "CLOUD KITCHEN/TAKEAWAYS",
				Description: "Specialized cloud kitchen and takeaway module designed for multi-brand virtual kitchens—supports order aggregation, delivery partner APIs, kitchen workflows, and real-time analytics.",
			},
		}

	case "education":
		data = []SubModule{
			{
				ID:          "1",
				ModuleName:  "SCHOOL",
				Description: "School management module",
			},
			{
				ID:          "2",
				ModuleName:  "COLLEGE",
				Description: "College operations module",
			},
		}

	default:
		utils.Failure(
			c,
			http.StatusNotFound,
			utils.CodeNotFound,
			errors.New("module not found"),
		)
		return
	}

	// Unified success response
	utils.Success(
		c,
		http.StatusOK,
		"List of All SubModules",
		map[string]any{
			"module":     module,
			"submodules": data,
		},
	)
}
