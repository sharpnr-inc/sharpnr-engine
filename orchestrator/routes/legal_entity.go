package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerLegalEntityRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {

	ler := r.Group("/legal-entity")
	{
		ler.GET("", co.ListAllLegalEntity) // fetches the all the legal Entity of the Organizations
		ler.GET("/:legalEntityID/details", co.GetLegalEntityDetailsHandler) // return the legal entity details by ID
		ler.POST("/create", co.AddLegalEntity) // creates a new Legal entity in the organization
		ler.POST("/:legalEntityID/add-address")
		ler.POST("/:legalEntityID/add-bankaccount")
		ler.POST("/:legalEntityID/add-identities")

		le_utils := ler.Group("/utils")
		{
			le_utils.GET("/entity_types", co.GetEntityTypes)
			le_utils.GET("/gst_types", co.GetGSTRegistrationTypes)
			le_utils.GET("/address_types", co.GetLegalEntityAddressTypes)
			le_utils.GET("/license_types", co.GetLegalEntityLicenseTypes)
			le_utils.GET("/contact_types", co.GetLegalEntityContactTypes)
			le_utils.GET("/identifier_types", co.GetLegalEntityIdentifierTypes)
		}
	}
}
