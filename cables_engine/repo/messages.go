package cablesEngineRepo

import (
	"context"
	"errors"
	"time"

	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NextConversationSeq implements [CablesRepoInterface]: allocate this
// conversation's next message sequence.
//
// One atomic UPDATE ... RETURNING, so two concurrent sends cannot take the same
// number — the row lock is held for the duration of the statement and the
// second send sees the first one's value. Doing this as SELECT-then-UPDATE
// would hand both senders the same seq and one of them would hit the unique
// index instead.
//
// Must be called inside the send transaction: a rolled-back send would
// otherwise leave a burned sequence number and a gap in the client's cursor.
func (c *cablesRepo) NextConversationSeq(ctx context.Context, conversationID uuid.UUID) (int64, error) {
	var seq int64

	err := c.conn(ctx).
		Raw(`UPDATE cables_engine.cables_conversations
		     SET last_seq = last_seq + 1, updated_at = now()
		     WHERE id = ?
		     RETURNING last_seq`, conversationID).
		Scan(&seq).
		Error

	if err != nil {
		return 0, err
	}
	if seq == 0 {
		return 0, cablesEngineErrs.ErrConversationNotFound
	}

	return seq, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CreateMessage implements [CablesRepoInterface].
func (c *cablesRepo) CreateMessage(ctx context.Context, message *cables_engine_models.CablesMessage) error {
	return c.conn(ctx).Create(message).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetMessageByClientID implements [CablesRepoInterface]: the idempotency probe.
//
// A client that retries a send after a dropped connection resolves to the row
// it already wrote instead of posting a duplicate.
func (c *cablesRepo) GetMessageByClientID(ctx context.Context, conversationID uuid.UUID, clientMessageID string) (*cables_engine_models.CablesMessage, error) {
	if clientMessageID == "" {
		return nil, cablesEngineErrs.ErrMessageNotFound
	}

	var message cables_engine_models.CablesMessage

	err := c.conn(ctx).
		Where("conversation_id = ? AND client_message_id = ?", conversationID, clientMessageID).
		First(&message).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrMessageNotFound
		}
		return nil, err
	}

	return &message, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListMessages implements [CablesRepoInterface]: one page of history, newest
// first, hitting idx_cables_messages_conversation_seq_desc.
//
// beforeSeq is the cursor rather than an offset or a page number: history keeps
// growing at the end being read, so OFFSET would skip or repeat rows as people
// post mid-scroll. Zero means "from the newest".
func (c *cablesRepo) ListMessages(ctx context.Context, conversationID uuid.UUID, beforeSeq int64, limit int) ([]cables_engine_models.CablesMessage, error) {
	var messages []cables_engine_models.CablesMessage

	query := c.conn(ctx).
		Where("conversation_id = ?", conversationID)

	if beforeSeq > 0 {
		query = query.Where("seq < ?", beforeSeq)
	}

	err := query.
		Order("seq DESC").
		Limit(limit).
		Find(&messages).
		Error

	return messages, err
}

/* ---------------------------------- <***> --------------------------------- */

// TouchConversationAfterSend implements [CablesRepoInterface]: refresh the
// denormalised columns the sidebar renders from.
//
// Runs in the send transaction so the preview can never describe a message that
// was rolled back.
func (c *cablesRepo) TouchConversationAfterSend(ctx context.Context, conversationID, messageID uuid.UUID, brief string, sentAt time.Time) error {
	return c.conn(ctx).
		Model(&cables_engine_models.CablesConversation{}).
		Where("id = ?", conversationID).
		Updates(map[string]any{
			"last_message_id":    messageID,
			"last_message_at":    sentAt,
			"last_message_brief": brief,
			"message_count":      gorm.Expr("message_count + 1"),
			"updated_at":         sentAt,
		}).
		Error
}
