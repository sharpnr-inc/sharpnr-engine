package coreutilitiesapi

import (
	"net/http"
	"sharpnr-engine/core/constants"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

// GetOrgTypes returns the list of available organization types.
func GetOrgTypes(c *gin.Context) {
	orgTypes := []constants.OrgType{
		constants.Hotel,
		constants.Retail,
		constants.Salon,
		constants.Club,
		constants.Restaurant,
	}

	// ✅ Wrap orgTypes inside a map for "ORG_TYPE"
	wrappedData := map[string]any{
		"ORG_TYPE": orgTypes,
	}

	// ✅ Use typed generic DTO response
	res := dto.Response[map[string]any]{
		StatusCode: http.StatusOK,
		Message:    "List of All Organization Types",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &wrappedData, // ✅ pointer to map
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, res)
}
