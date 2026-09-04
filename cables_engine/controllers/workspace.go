package cablesEngineControllers

import (
	"net/http"
	"strconv"

	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// GetBootstrapHandler godoc
// @Summary      Bootstrap the cables workspace
// @Description  Returns the caller's cables identity, the org roster and the sidebar in one round trip, and reconciles the roster so an employee hired since the last visit gets a cables identity. Messages are not included — they load per conversation.
// @Tags         Cables
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=cablesEngineDtos.BootstrapResponseDto} "Workspace loaded"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      409     {object}  utils.Response "Cables is not started or is switched off"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /cables/bootstrap [get]
func (c *cablesEngineController) GetBootstrapHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	workspace, err := c.cablesService.Bootstrap(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Workspace loaded", workspace)
}

/* ---------------------------------- <***> --------------------------------- */

// ListConversationsHandler godoc
// @Summary      List the caller's conversations
// @Description  The sidebar on its own — channels and DMs the caller is a member of, with unread counts, ordered by most recent activity.
// @Tags         Cables
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=[]cablesEngineDtos.ConversationSummaryDto} "Conversations fetched"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /cables/conversations [get]
func (c *cablesEngineController) ListConversationsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	conversations, err := c.cablesService.ListConversations(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Conversations fetched", conversations)
}

/* ---------------------------------- <***> --------------------------------- */

// OpenDirectConversationHandler godoc
// @Summary      Open a direct message
// @Description  Gets or creates the DM between the caller and one other member. Idempotent — the same pair always resolves to the same conversation, so a double-click cannot create two.
// @Tags         Cables
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                                true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      cablesEngineDtos.OpenDirectConversationRequestDto     true  "The other member"
// @Success      200      {object}  utils.Response{data=cablesEngineDtos.ConversationSummaryDto} "Conversation opened"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      404      {object}  utils.Response "Member not found"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /cables/conversations/direct [post]
func (c *cablesEngineController) OpenDirectConversationHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.OpenDirectConversationRequestDto](ctx)
	if !ok {
		return
	}

	conversation, err := c.cablesService.OpenDirectConversation(ctx.Request.Context(), orgID, emplID, req.MemberID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Conversation opened", conversation)
}

/* ---------------------------------- <***> --------------------------------- */

// GetMessagesHandler godoc
// @Summary      Read a page of messages
// @Description  Newest first. `before` is a seq cursor, not an offset — pass the previous page's nextCursor to walk backwards through history.
// @Tags         Cables
// @Produce      json
// @Param        Cookie          header    string  true   "Session cookie containing org_id and empl_id"
// @Param        conversationID  path      string  true   "Conversation id"
// @Param        before          query     int     false  "Return messages with a seq lower than this"
// @Param        limit           query     int     false  "Page size (default 50, max 100)"
// @Success      200             {object}  utils.Response{data=cablesEngineDtos.MessagePageResponseDto} "Messages fetched"
// @Failure      403             {object}  utils.Response "Not a member of this conversation"
// @Failure      404             {object}  utils.Response "Conversation not found"
// @Router       /cables/conversations/{conversationID}/messages [get]
func (c *cablesEngineController) GetMessagesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	conversationID, ok := c.resolveConversationID(ctx)
	if !ok {
		return
	}

	// Both are optional and both are advisory — a malformed value falls back to
	// the default rather than 400ing a read.
	beforeSeq, _ := strconv.ParseInt(ctx.Query("before"), 10, 64)
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	page, err := c.cablesService.GetMessages(ctx.Request.Context(), orgID, emplID, conversationID, beforeSeq, limit)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Messages fetched", page)
}

/* ---------------------------------- <***> --------------------------------- */

// SendMessageHandler godoc
// @Summary      Send a message
// @Description  Writes the message and then fans it out over the websocket. Pass a clientMessageId so a retry after a dropped connection resolves to the row already written instead of double-posting.
// @Tags         Cables
// @Accept       json
// @Produce      json
// @Param        Cookie          header    string                                   true  "Session cookie containing org_id and empl_id"
// @Param        conversationID  path      string                                   true  "Conversation id"
// @Param        payload         body      cablesEngineDtos.SendMessageRequestDto   true  "The message"
// @Success      201             {object}  utils.Response{data=cablesEngineDtos.MessageResponseDto} "Message sent"
// @Failure      400             {object}  utils.Response "Validation error"
// @Failure      403             {object}  utils.Response "Not a member of this conversation"
// @Router       /cables/conversations/{conversationID}/messages [post]
func (c *cablesEngineController) SendMessageHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	conversationID, ok := c.resolveConversationID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.SendMessageRequestDto](ctx)
	if !ok {
		return
	}

	message, err := c.cablesService.SendMessage(ctx.Request.Context(), orgID, emplID, conversationID, *req)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Message sent", message)
}

/* ---------------------------------- <***> --------------------------------- */

// MarkReadHandler godoc
// @Summary      Move the read cursor
// @Description  Clears the sidebar badge up to the given seq. The cursor never travels backwards server-side, so two tabs reporting out of order is harmless.
// @Tags         Cables
// @Accept       json
// @Produce      json
// @Param        Cookie          header    string                                 true  "Session cookie containing org_id and empl_id"
// @Param        conversationID  path      string                                 true  "Conversation id"
// @Param        payload         body      cablesEngineDtos.MarkReadRequestDto    true  "Read position"
// @Success      200             {object}  utils.Response "Read position updated"
// @Failure      403             {object}  utils.Response "Not a member of this conversation"
// @Router       /cables/conversations/{conversationID}/read [post]
func (c *cablesEngineController) MarkReadHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	conversationID, ok := c.resolveConversationID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.MarkReadRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.cablesService.MarkRead(ctx.Request.Context(), orgID, emplID, conversationID, req.Seq); err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Read position updated")
}

/* ---------------------------------- <***> --------------------------------- */

// ListChannelsHandler godoc
// @Summary      Browse channels
// @Description  Every channel in the org's cables account, for the channel browser. The caller's own channels come from the conversation list instead.
// @Tags         Cables
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=[]cablesEngineDtos.ChannelResponseDto} "Channels fetched"
// @Router       /cables/channels [get]
func (c *cablesEngineController) ListChannelsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	channels, err := c.cablesService.ListChannels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Channels fetched", channels)
}

/* ---------------------------------- <***> --------------------------------- */

// CreateChannelHandler godoc
// @Summary      Create a channel
// @Description  Creates the channel, the conversation its messages live in, and the creator's owner membership, in one transaction.
// @Tags         Cables
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                      true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      cablesEngineDtos.CreateChannelRequestDto    true  "Channel to create"
// @Success      201      {object}  utils.Response "Channel created"
// @Failure      409      {object}  utils.Response "A channel with this name already exists"
// @Router       /cables/channels [post]
func (c *cablesEngineController) CreateChannelHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.CreateChannelRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.cablesService.AddNewChannel(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Channel created")
}

/* ---------------------------------- <***> --------------------------------- */

// JoinChannelHandler godoc
// @Summary      Join a public channel
// @Description  Adds the caller to a public channel and returns its sidebar row. Idempotent, and rejoining keeps the read state you had.
// @Tags         Cables
// @Produce      json
// @Param        Cookie     header    string  true  "Session cookie containing org_id and empl_id"
// @Param        channelID  path      string  true  "Channel id"
// @Success      200        {object}  utils.Response{data=cablesEngineDtos.ConversationSummaryDto} "Channel joined"
// @Failure      403        {object}  utils.Response "The channel is private"
// @Failure      404        {object}  utils.Response "Channel not found"
// @Router       /cables/channels/{channelID}/join [post]
func (c *cablesEngineController) JoinChannelHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	channelID, ok := c.resolveUUIDParam(ctx, "channelID")
	if !ok {
		return
	}

	conversation, err := c.cablesService.JoinChannel(ctx.Request.Context(), orgID, emplID, channelID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Channel joined", conversation)
}
