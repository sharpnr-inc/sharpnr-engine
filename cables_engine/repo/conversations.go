package cablesEngineRepo

import (
	"context"
	"errors"
	"time"

	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ConversationWithMembership is the sidebar projection: one conversation joined
// to the caller's own membership row.
//
// It exists because the sidebar needs both halves at once — the conversation's
// last_seq and preview, and this member's last_read_seq and mute flag — and
// returning them separately would mean a second query per row.
type ConversationWithMembership struct {
	cables_engine_models.CablesConversation

	// From the caller's membership row.
	Role               string
	LastReadSeq        int64
	UnreadMentionCount int
	IsMuted            bool
	IsPinned           bool
	IsHidden           bool

	// Channel identity, null for DMs.
	ChannelName *string
	ChannelSlug *string
	ChannelType *string
}

/* ---------------------------------- <***> --------------------------------- */

// CreateConversation implements [CablesRepoInterface].
func (c *cablesRepo) CreateConversation(ctx context.Context, conversation *cables_engine_models.CablesConversation) error {
	return c.conn(ctx).Create(conversation).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetConversation implements [CablesRepoInterface]. The org predicate is on the
// query, so a conversation id from another tenant reads as not found.
func (c *cablesRepo) GetConversation(ctx context.Context, orgID, conversationID uuid.UUID) (*cables_engine_models.CablesConversation, error) {
	var conversation cables_engine_models.CablesConversation

	err := c.conn(ctx).
		Where("organization_id = ? AND id = ?", orgID, conversationID).
		First(&conversation).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrConversationNotFound
		}
		return nil, err
	}

	return &conversation, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetDirectConversation implements [CablesRepoInterface]: the existing DM for a
// pair, or ErrConversationNotFound.
func (c *cablesRepo) GetDirectConversation(ctx context.Context, orgCablesAccountID uuid.UUID, directKey string) (*cables_engine_models.CablesConversation, error) {
	var conversation cables_engine_models.CablesConversation

	err := c.conn(ctx).
		Where("org_cables_account_id = ? AND direct_key = ?", orgCablesAccountID, directKey).
		First(&conversation).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrConversationNotFound
		}
		return nil, err
	}

	return &conversation, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListConversationsForMember implements [CablesRepoInterface]: the sidebar.
//
// One query for the whole left rail. It drives the ordering off the
// conversation's denormalised last_message_at rather than looking at messages,
// which is the entire reason that column exists.
//
// A conversation with no membership row is simply absent — that is how a
// private channel stays invisible without a visibility check here.
func (c *cablesRepo) ListConversationsForMember(ctx context.Context, orgID, cablesMemberID uuid.UUID) ([]ConversationWithMembership, error) {
	var rows []ConversationWithMembership

	err := c.conn(ctx).
		Table("cables_engine.cables_conversations AS conv").
		Select(`conv.*,
			mem.role                 AS role,
			mem.last_read_seq        AS last_read_seq,
			mem.unread_mention_count AS unread_mention_count,
			mem.is_muted             AS is_muted,
			mem.is_pinned            AS is_pinned,
			mem.is_hidden            AS is_hidden,
			ch.name                  AS channel_name,
			ch.slug                  AS channel_slug,
			ch.type                  AS channel_type`).
		Joins(`JOIN cables_engine.cables_conversation_members AS mem
			ON mem.conversation_id = conv.id
			AND mem.cables_member_id = ?
			AND mem.removed_at IS NULL
			AND mem.deleted_at IS NULL`, cablesMemberID).
		Joins(`LEFT JOIN cables_engine.cables_channels AS ch
			ON ch.id = conv.channel_id AND ch.deleted_at IS NULL`).
		Where("conv.organization_id = ? AND conv.deleted_at IS NULL", orgID).
		Where("mem.is_hidden = false").
		Order("mem.is_pinned DESC, conv.last_message_at DESC NULLS LAST, conv.created_at DESC").
		Scan(&rows).
		Error

	return rows, err
}

/* -------------------------------------------------------------------------- */
/*                            Conversation members                            */
/* -------------------------------------------------------------------------- */

// AddConversationMember implements [CablesRepoInterface].
//
// Upsert rather than insert: the unique index is unconditional, so somebody
// rejoining a channel they left clears removed_at on their original row instead
// of creating a second one. That keeps their read state — a rejoin should not
// mark five thousand messages unread.
func (c *cablesRepo) AddConversationMember(ctx context.Context, member *cables_engine_models.CablesConversationMember) error {
	return c.conn(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "conversation_id"}, {Name: "cables_member_id"}},
			// Repeats the partial index's predicate so Postgres can infer it.
			TargetWhere: clause.Where{Exprs: []clause.Expression{
				clause.Expr{SQL: "deleted_at IS NULL"},
			}},
			DoUpdates: clause.Assignments(map[string]any{
				"removed_at": nil,
				"removed_by": nil,
				"is_hidden":  false,
				"updated_at": time.Now(),
			}),
		}).
		Create(member).
		Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetConversationMember implements [CablesRepoInterface]. This is the
// membership check every read, send and websocket subscribe runs.
func (c *cablesRepo) GetConversationMember(ctx context.Context, conversationID, cablesMemberID uuid.UUID) (*cables_engine_models.CablesConversationMember, error) {
	var member cables_engine_models.CablesConversationMember

	err := c.conn(ctx).
		Where("conversation_id = ? AND cables_member_id = ? AND removed_at IS NULL", conversationID, cablesMemberID).
		First(&member).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cablesEngineErrs.ErrNotAConversationMember
		}
		return nil, err
	}

	return &member, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListConversationMembers implements [CablesRepoInterface]: everyone still in
// the conversation, for the member list and for notification fan-out.
func (c *cablesRepo) ListConversationMembers(ctx context.Context, conversationID uuid.UUID) ([]cables_engine_models.CablesConversationMember, error) {
	var members []cables_engine_models.CablesConversationMember

	err := c.conn(ctx).
		Where("conversation_id = ? AND removed_at IS NULL", conversationID).
		Find(&members).
		Error

	return members, err
}

/* ---------------------------------- <***> --------------------------------- */

// MarkConversationRead implements [CablesRepoInterface].
//
// GREATEST() rather than a plain assignment: read receipts arrive out of order
// from two tabs, and the badge must never travel backwards.
func (c *cablesRepo) MarkConversationRead(ctx context.Context, conversationID, cablesMemberID uuid.UUID, seq int64) error {
	now := time.Now()

	return c.conn(ctx).
		Model(&cables_engine_models.CablesConversationMember{}).
		Where("conversation_id = ? AND cables_member_id = ?", conversationID, cablesMemberID).
		Updates(map[string]any{
			"last_read_seq":        gorm.Expr("GREATEST(last_read_seq, ?)", seq),
			"last_read_at":         now,
			"unread_mention_count": 0,
			"updated_at":           now,
		}).
		Error
}

/* ---------------------------------- <***> --------------------------------- */

// RevealConversation implements [CablesRepoInterface]: un-hide a conversation
// for everyone in it.
//
// This is what makes an opened-but-silent DM appear in the other person's
// sidebar only once somebody actually says something.
func (c *cablesRepo) RevealConversation(ctx context.Context, conversationID uuid.UUID) error {
	return c.conn(ctx).
		Model(&cables_engine_models.CablesConversationMember{}).
		Where("conversation_id = ? AND is_hidden = true", conversationID).
		Updates(map[string]any{"is_hidden": false, "updated_at": time.Now()}).
		Error
}
