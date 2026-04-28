package hr_routes

import (
	"sharpnr-engine/app"

	"github.com/gin-gonic/gin"
)

func RegisterHRRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	hr := v1.Group("/sharp-hr")
	registerEmployeeRoutes(hr, modules)
}
