package coreutilitiesapi

import (
	"net/http"

	"sharpnr-engine/core/constants"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// GetOrgTypes returns the list of available organization types.
func GetOrgTypes(c *gin.Context) {
	orgTypes := []constants.OrgType{
		constants.Hospitality,
		constants.Retail,
		constants.Salon,
	}

	// Wrap inside a key for frontend contract
	wrappedData := map[string]any{
		"ORG_TYPE": orgTypes,
	}

	utils.Success(
		c,
		http.StatusOK,
		"List of All Organization Types",
		wrappedData,
	)
}
