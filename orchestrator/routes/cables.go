package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCablesOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	cables := r.Group("/cables")
	{
		// Is cables on for this org, and is the account live? The client asks
		// this before anything else, so it answers isStarted=false rather than
		// 404 for an org that never started the service.
		cables.GET("", co.GetCablesAccountStatusHandler)
		cables.POST("/start", co.StartCablesServiceHandler)

		cables_chan := cables.Group("/channel")
		{
			cables_chan.GET("", co.ListCablesChannelsHandler)
			cables_chan.POST("", co.CreateCablesChannelHandler)

			// Channel identity only for now. The conversation and its messages
			// hang off the conversationId this returns, once those tables land.
			cables_chan.GET("/:channelID", co.GetCablesChannelHandler)
		}

	}

}
