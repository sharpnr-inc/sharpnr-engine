package cablesEngineService

import (
	"context"
	"errors"
	"strings"
	"time"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetMessages implements [CablesEngineServiceInterface]: one page of history,
// newest first.
//
// beforeSeq is a cursor, not an offset. History grows at the end being read, so
// OFFSET would skip or repeat rows whenever somebody posts mid-scroll.
func (c *cablesService) GetMessages(
	ctx context.Context,
	orgID, emplID, conversationID uuid.UUID,
	beforeSeq int64,
	limit int,
) (cablesEngineDtos.MessagePageResponseDto, error) {

	if _, _, err := c.authorizeConversation(ctx, orgID, emplID, conversationID); err != nil {
		return cablesEngineDtos.MessagePageResponseDto{}, err
	}

	switch {
	case limit <= 0:
		limit = cablesEngineConstants.DefaultMessagePageSize
	case limit > cablesEngineConstants.MaxMessagePageSize:
		limit = cablesEngineConstants.MaxMessagePageSize
	}

	// One extra row is the has-more probe: cheaper and more honest than a
	// separate COUNT, and it never disagrees with the page it describes.
	messages, err := c.cablesRepo.ListMessages(ctx, conversationID, beforeSeq, limit+1)
	if err != nil {
		return cablesEngineDtos.MessagePageResponseDto{}, err
	}

	hasMore := len(messages) > limit
	if hasMore {
		messages = messages[:limit]
	}

	views := make([]cablesEngineDtos.MessageResponseDto, 0, len(messages))
	for _, m := range messages {
		views = append(views, toMessageView(m))
	}

	page := cablesEngineDtos.MessagePageResponseDto{
		Messages: views,
		HasMore:  hasMore,
	}

	if hasMore && len(messages) > 0 {
		cursor := messages[len(messages)-1].Seq
		page.NextCursor = &cursor
	}

	return page, nil
}

/* ---------------------------------- <***> --------------------------------- */

// SendMessage implements [CablesEngineServiceInterface].
//
// The write goes over HTTP, not over the websocket: this is where membership,
// validation and idempotency are enforced, and none of that belongs on a socket
// frame. The socket only carries the fan-out at the end.
//
// Everything that must agree — the sequence number, the row, and the sidebar
// preview — happens in one transaction. The publish happens strictly after it
// commits: announcing a message a reader cannot yet fetch is a race that shows
// up as a phantom row in the client.
func (c *cablesService) SendMessage(
	ctx context.Context,
	orgID, emplID, conversationID uuid.UUID,
	request cablesEngineDtos.SendMessageRequestDto,
) (cablesEngineDtos.MessageResponseDto, error) {

	me, conversation, err := c.authorizeConversation(ctx, orgID, emplID, conversationID)
	if err != nil {
		return cablesEngineDtos.MessageResponseDto{}, err
	}

	body := strings.TrimSpace(request.Body)
	if body == "" {
		return cablesEngineDtos.MessageResponseDto{}, cablesEngineErrs.ErrMessageEmpty
	}
	if len(body) > cablesEngineConstants.MaxMessageBodyLength {
		return cablesEngineDtos.MessageResponseDto{}, cablesEngineErrs.ErrMessageTooLong
	}

	clientMessageID := strings.TrimSpace(request.ClientMessageID)

	// A retry after a dropped connection must not double-post. Checked here for
	// a clean answer; the partial unique index is the race backstop.
	if clientMessageID != "" {
		existing, err := c.cablesRepo.GetMessageByClientID(ctx, conversationID, clientMessageID)
		if err != nil && !errors.Is(err, cablesEngineErrs.ErrMessageNotFound) {
			return cablesEngineDtos.MessageResponseDto{}, err
		}
		if existing != nil {
			return toMessageView(*existing), nil
		}
	}

	now := time.Now()
	message := cables_engine_models.CablesMessage{
		ID:             uuid.New(),
		OrganizationID: orgID,
		ConversationID: conversationID,
		SenderMemberID: &me.ID,
		Type:           cablesEngineConstants.MessageTypeText,
		Body:           body,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if clientMessageID != "" {
		message.ClientMessageID = &clientMessageID
	}

	err = c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Allocated inside the transaction: a rolled-back send must not burn a
		// sequence number and leave a hole in the client's cursor.
		seq, err := c.cablesRepo.NextConversationSeq(txCtx, conversationID)
		if err != nil {
			return err
		}
		message.Seq = seq

		if err := c.cablesRepo.CreateMessage(txCtx, &message); err != nil {
			return err
		}

		if err := c.cablesRepo.TouchConversationAfterSend(txCtx, conversationID, message.ID, buildMessageBrief(body), now); err != nil {
			return err
		}

		// First message on a DM is what puts it in the other person's sidebar.
		return c.cablesRepo.RevealConversation(txCtx, conversationID)
	})
	if err != nil {
		return cablesEngineDtos.MessageResponseDto{}, err
	}

	// Committed — safe to announce.
	view := toMessageView(message)
	c.publishToConversation(conversationID, cablesEngineConstants.EventMessageCreated, view)

	// A DM also goes to each participant's own topic.
	//
	// Without this, the first message of a brand-new DM reaches nobody live:
	// the recipient has no reason to be subscribed to a conversation they have
	// never seen, so they would only discover it on their next refetch. Their
	// member topic is the one thing they are always listening on, which is what
	// makes a DM able to arrive anywhere in the suite — including from an app
	// that is not cables.
	//
	// Restricted to DMs on purpose: doing this for a channel would turn one
	// message into a publish per member.
	if conversation.Type == cablesEngineConstants.ConversationTypeDirect {
		c.notifyDirectParticipants(ctx, conversationID, me.ID, view)
	}

	// The sender's own read cursor follows their own message, so their sidebar
	// does not light up with a badge for something they just wrote.
	if err := c.cablesRepo.MarkConversationRead(ctx, conversationID, me.ID, message.Seq); err != nil {
		// Not worth failing a delivered message over: the badge self-corrects
		// the moment they open the conversation.
		c.logRealtimeFailure("mark own message read", err)
	}

	return view, nil
}

/* ---------------------------------- <***> --------------------------------- */

// MarkRead implements [CablesEngineServiceInterface]: move this member's read
// cursor, which is what clears the sidebar badge.
func (c *cablesService) MarkRead(ctx context.Context, orgID, emplID, conversationID uuid.UUID, seq int64) error {
	me, _, err := c.authorizeConversation(ctx, orgID, emplID, conversationID)
	if err != nil {
		return err
	}

	return c.cablesRepo.MarkConversationRead(ctx, conversationID, me.ID, seq)
}

/* ---------------------------------- <***> --------------------------------- */

// buildMessageBrief is the sidebar preview stored on the conversation row.
// Truncated on a rune boundary so a multi-byte character is never cut in half.
func buildMessageBrief(body string) string {
	brief := strings.Join(strings.Fields(body), " ")

	if len(brief) <= cablesEngineConstants.MaxMessageBriefLength {
		return brief
	}

	runes := []rune(brief)
	if len(runes) <= cablesEngineConstants.MaxMessageBriefLength {
		return brief
	}

	return string(runes[:cablesEngineConstants.MaxMessageBriefLength]) + "…"
}
