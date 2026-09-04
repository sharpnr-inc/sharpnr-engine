package cablesEngineRoutes

import (
	"sharpnr-engine/app"
	cablesEngineControllers "sharpnr-engine/cables_engine/controllers"

	"github.com/gin-gonic/gin"
)

func ResgisterCablesEngineRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,

) {

	cablesController := cablesEngineControllers.NewCablesEngineController(modules.CablesEngineService())

	cables := v1.Group("/cables")
	{
		cables.GET("/status", cablesController.GetCablesAccountStatus)

		// The client's first call: identity, roster and sidebar in one round
		// trip, plus the roster reconcile. Everything else assumes it has run.
		cables.GET("/bootstrap", cablesController.GetBootstrapHandler)

		channels := cables.Group("/channels")
		{
			channels.GET("", cablesController.ListChannelsHandler)
			channels.POST("", cablesController.CreateChannelHandler)
			channels.POST("/:channelID/join", cablesController.JoinChannelHandler)
		}

		// Channels and DMs are the same thing below this line — both are
		// conversations, so one set of message endpoints serves both.
		conversations := cables.Group("/conversations")
		{
			conversations.GET("", cablesController.ListConversationsHandler)
			conversations.POST("/direct", cablesController.OpenDirectConversationHandler)

			conversations.GET("/:conversationID/messages", cablesController.GetMessagesHandler)
			conversations.POST("/:conversationID/messages", cablesController.SendMessageHandler)
			conversations.POST("/:conversationID/read", cablesController.MarkReadHandler)
		}
	}

}
