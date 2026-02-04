package v1

import (
	"sharpnr-engine/core/coreutilitiesapi"

	"github.com/gin-gonic/gin"
)

// UtilsRoutes registers utility endpoints (for enums, static data, etc.)
func UtilsRoutes(rg *gin.RouterGroup) {
	utils := rg.Group("/utils")
	utils.GET("/countries", coreutilitiesapi.GetCountriesHandler)
	utils.GET("/countries/:countryCode/states", coreutilitiesapi.GetStatesHandler)
	utils.GET("/countries/:countryCode/states/:stateCode/cities", coreutilitiesapi.GetCitiesHandler)
}
