package v1

import (
	orchestratorcontroller "sharpnr-engine/core/controllers/orchestratorController"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	orchestrator_org_service "sharpnr-engine/core/services/orchestratorOrg_service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CoreOrgRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	newOrchestratorRepo := organizationrepositories.NewOrganizationRepo(db, redis)
	newOrchestratorService := orchestrator_org_service.NewOrchestratorOrgService(newOrchestratorRepo)
	newOrchestratorController := orchestratorcontroller.NewOrchestratorOrganizationController(newOrchestratorService)

	coreOrg := v1.Group("/orchestrator/org/")

	coreOrg.GET("get-details", newOrchestratorController.GetDetails)
	coreOrg.POST("update")
	coreOrg.POST("add-bank-accounts")
}
