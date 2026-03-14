package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerLegalEntityRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {

	ler := r.Group("/legal-entity")
	{
		ler.GET("", co.ListAllLegalEntity)
		ler.GET("/overview")
		ler.POST("/create", co.AddLegalEntity)
		ler.POST("/update")

		letr := ler.Group("/tax-registration/:legalEntityID")
		{
			letr.POST("/add", co.AddTaxRegistrationLegalEntity)
			letr.GET("", co.ListAllTaxRegistrationOfLegalEntity) // this returns the all tax registration
		}

		letc := ler.Group("/tax-category/:legalEntityID")
		{
			letc.POST("/add", co.AddTaxCategoryInLegalEntity)
			letc.GET("", co.ListAllTaxCategoryOfLegalEntry) // this will return all the Tax Category of the of Legal Entity

		}
	}
}
