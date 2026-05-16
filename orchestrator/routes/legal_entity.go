package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerLegalEntityRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {

	ler := r.Group("/legal-entity")
	{
		ler.GET("", co.ListAllLegalEntity)
		ler.POST("/create", co.AddLegalEntity)
		ler.GET("/:legalEntityID/details")
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
