package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCablesOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	cables := r.Group("/cables")
	{
		cables.GET("") // returns
		cables.POST("/start")

		cables_chan := cables.Group("/channel")
		{
			cables_chan.GET("")            // return all the channels
			cables_chan.POST("")           // creates a new channels
			cables_chan.GET("/:channelID") // returns the channels details and all the conversation and stuff
		}

	}

}
