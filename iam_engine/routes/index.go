package iamEngineRoutes

import (
	"sharpnr-engine/app"
	iamEngineControllers "sharpnr-engine/iam_engine/controllers"
	"sharpnr-engine/pkg/cookieManager"

	"github.com/gin-gonic/gin"
)

func RegisterIamEngineRoutes(v1 *gin.RouterGroup, cookies cookieManager.Manager, modules app.ModulesProvider) {

	iamController := iamEngineControllers.NewIamEngineServiceController(modules.IamEngineService())

	iamRoutes := v1.Group("/iam")
	{
		iamRoutes.GET("/profile", iamController.GetMyProfile)
	}
}
