package organization

import "github.com/gin-gonic/gin"

func RegisterOrganisationRoutes(rg *gin.RouterGroup) {
    orgGroup := rg.Group("/organisation")

    // Directly use controller funcs since same package
    orgGroup.GET("/", GetOrganisations)
    orgGroup.GET("/:id", GetOrganisationByID)
}