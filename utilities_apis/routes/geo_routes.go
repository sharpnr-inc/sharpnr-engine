package utilitiesApi_routes

import (
	"sharpnr-engine/app"
	utilitiesApi_controllers "sharpnr-engine/utilities_apis/controllers"

	"github.com/gin-gonic/gin"
)

// UtilitesApiGeoRoutes registers the self-hosted geo endpoints. The paths mirror
// countrystatecity.in (https://api.countrystatecity.in/v1) so the frontend only
// needs to swap its base URL. Mounted under /api/v1, e.g. GET /api/v1/countries.
func UtilitesApiGeoRoutes(r *gin.RouterGroup, modules app.ModulesProvider) {
	ctrl := utilitiesApi_controllers.NewUtilitiesApiController(modules.UtilitiesApiService())

	r.GET("/countries", ctrl.GetAllCountries)
	r.GET("/countries/:ciso", ctrl.GetCountry)
	r.GET("/countries/:ciso/states", ctrl.GetStatesByCountry)
	r.GET("/countries/:ciso/states/:siso", ctrl.GetState)
	r.GET("/countries/:ciso/states/:siso/cities", ctrl.GetCitiesByState)
}
