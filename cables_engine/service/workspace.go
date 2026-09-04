package cablesEngineService

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"
	coreEmployeeConstants "sharpnr-engine/core/employees/constants"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Bootstrap implements [CablesEngineServiceInterface]: everything the shell
// needs to render, in one round trip.
//
// It also reconciles the roster, which is why this is the endpoint the client
// calls first: an employee hired since the last visit gets their cables member
// row here rather than through a hook on employee creation. Coupling the
// employee module to cables would mean HR writes failing when cables is off.
func (c *cablesService) Bootstrap(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.BootstrapResponseDto, error) {
	account, err := c.activeAccount(ctx, orgID)
	if err != nil {
		return cablesEngineDtos.BootstrapResponseDto{}, err
	}

	if err := c.syncRoster(ctx, orgID, account.ID); err != nil {
		return cablesEngineDtos.BootstrapResponseDto{}, err
	}

	me, err := c.cablesRepo.GetMemberByEmployee(ctx, account.ID, emplID)
	if err != nil {
		return cablesEngineDtos.BootstrapResponseDto{}, err
	}

	members, err := c.cablesRepo.ListMembers(ctx, account.ID)
	if err != nil {
		return cablesEngineDtos.BootstrapResponseDto{}, err
	}

	conversations, err := c.listConversations(ctx, orgID, me.ID)
	if err != nil {
		return cablesEngineDtos.BootstrapResponseDto{}, err
	}

	memberViews := make([]cablesEngineDtos.MemberResponseDto, 0, len(members))
	for _, m := range members {
		memberViews = append(memberViews, toMemberView(m))
	}

	return cablesEngineDtos.BootstrapResponseDto{
		Me:            toMemberView(*me),
		Account:       cablesEngineDtos.CablesServiceResponseDto{IsStarted: true, IsActive: account.IsActive},
		Conversations: conversations,
		Members:       memberViews,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// syncRoster mirrors the org's employees into cables_members.
//
// The one cross-module read in the whole module, and it happens here rather
// than on any hot path: after this, rendering an author or a member picker is a
// single-table read of cables' own data. Existing rows keep their id and
// presence — see UpsertEmployeeMembers — so a sync never disturbs a live
// session or orphans the ids messages point at.
func (c *cablesService) syncRoster(ctx context.Context, orgID, orgCablesAccountID uuid.UUID) error {
	if c.employeeIamService == nil {
		return nil
	}

	employees, err := c.employeeIamService.ListOrgEmployees(ctx, orgID)
	if err != nil {
		return fmt.Errorf("failed to read the organization roster: %w", err)
	}

	members := make([]cables_engine_models.CablesMember, 0, len(employees))
	now := time.Now()

	for _, e := range employees {
		// An exited employee keeps their member row — their old messages still
		// have to render an author — but is not re-created once removed.
		if e.Status == string(coreEmployeeConstants.EmployeeStatusExited) {
			continue
		}

		employeeID := e.ID
		displayName := strings.TrimSpace(e.DisplayName)
		if displayName == "" {
			displayName = e.Email
		}

		members = append(members, cables_engine_models.CablesMember{
			ID:                 uuid.New(),
			OrganizationID:     orgID,
			OrgCablesAccountID: orgCablesAccountID,
			EmployeeID:         &employeeID,
			Type:               cablesEngineConstants.MemberTypeEmployee,
			DisplayName:        displayName,
			Presence:           cablesEngineConstants.PresenceStatusOffline,
			CreatedAt:          now,
			UpdatedAt:          now,
		})
	}

	return c.cablesRepo.UpsertEmployeeMembers(ctx, members)
}

/* ---------------------------------- <***> --------------------------------- */

// ListConversations implements [CablesEngineServiceInterface]: the sidebar on
// its own, for a refresh that does not need the whole bootstrap payload.
func (c *cablesService) ListConversations(ctx context.Context, orgID, emplID uuid.UUID) ([]cablesEngineDtos.ConversationSummaryDto, error) {
	me, err := c.resolveMember(ctx, orgID, emplID)
	if err != nil {
		return nil, err
	}

	return c.listConversations(ctx, orgID, me.ID)
}

func (c *cablesService) listConversations(ctx context.Context, orgID, cablesMemberID uuid.UUID) ([]cablesEngineDtos.ConversationSummaryDto, error) {
	rows, err := c.cablesRepo.ListConversationsForMember(ctx, orgID, cablesMemberID)
	if err != nil {
		return nil, err
	}

	result := make([]cablesEngineDtos.ConversationSummaryDto, 0, len(rows))

	for _, row := range rows {
		summary := cablesEngineDtos.ConversationSummaryDto{
			ID:   row.ID,
			Type: string(row.Type),

			ChannelID:   row.ChannelID,
			Name:        row.ChannelName,
			Slug:        row.ChannelSlug,
			ChannelType: row.ChannelType,

			LastSeq:          row.LastSeq,
			LastMessageAt:    row.LastMessageAt,
			LastMessageBrief: row.LastMessageBrief,

			// The whole point of storing seq: the badge is a subtraction, not a
			// COUNT(*) over messages. Clamped because a read receipt can
			// legitimately arrive ahead of the caller's view of last_seq.
			UnreadCount:    max64(row.LastSeq-row.LastReadSeq, 0),
			UnreadMentions: row.UnreadMentionCount,
			IsMuted:        row.IsMuted,
			IsPinned:       row.IsPinned,

			MemberCount: row.ParticipantCount,
		}

		// A DM has no name of its own — the client renders the other person
		// from the roster, which is why this resolves an id rather than a
		// title.
		if row.Type == cablesEngineConstants.ConversationTypeDirect {
			otherID, err := c.otherDirectMember(ctx, row.ID, cablesMemberID)
			if err != nil {
				return nil, err
			}
			summary.OtherMemberID = otherID
		}

		result = append(result, summary)
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

// OpenDirectConversation implements [CablesEngineServiceInterface]: get or
// create the DM between the caller and one other member.
//
// Idempotent by construction. direct_key is the two member ids sorted and
// joined, unique per account, so a double-click or two devices racing resolve to
// the same conversation instead of creating two DMs with the same person.
func (c *cablesService) OpenDirectConversation(ctx context.Context, orgID, emplID uuid.UUID, otherMemberID uuid.UUID) (cablesEngineDtos.ConversationSummaryDto, error) {
	account, err := c.activeAccount(ctx, orgID)
	if err != nil {
		return cablesEngineDtos.ConversationSummaryDto{}, err
	}

	me, err := c.cablesRepo.GetMemberByEmployee(ctx, account.ID, emplID)
	if err != nil {
		return cablesEngineDtos.ConversationSummaryDto{}, err
	}

	if me.ID == otherMemberID {
		return cablesEngineDtos.ConversationSummaryDto{}, cablesEngineErrs.ErrCannotDirectMessageSelf
	}

	directKey := buildDirectKey(me.ID, otherMemberID)

	existing, err := c.cablesRepo.GetDirectConversation(ctx, account.ID, directKey)
	if err != nil && !errors.Is(err, cablesEngineErrs.ErrConversationNotFound) {
		return cablesEngineDtos.ConversationSummaryDto{}, err
	}

	if existing != nil {
		return c.summarizeDirect(ctx, *existing, me.ID, otherMemberID), nil
	}

	conversation := cables_engine_models.CablesConversation{
		ID:                 uuid.New(),
		OrganizationID:     orgID,
		OrgCablesAccountID: account.ID,
		Type:               cablesEngineConstants.ConversationTypeDirect,
		DirectKey:          &directKey,
		ParticipantCount:   2,
		CreatedBy:          &me.ID,
	}

	err = c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.cablesRepo.CreateConversation(txCtx, &conversation); err != nil {
			return err
		}

		// The opener sees the DM immediately; the other side stays hidden until
		// somebody actually says something, so opening a DM and typing nothing
		// does not clutter their sidebar.
		if err := c.addMember(txCtx, orgID, conversation.ID, me.ID, cablesEngineConstants.ParticipantRoleMember, nil, false); err != nil {
			return err
		}

		return c.addMember(txCtx, orgID, conversation.ID, otherMemberID, cablesEngineConstants.ParticipantRoleMember, &me.ID, true)
	})
	if err != nil {
		return cablesEngineDtos.ConversationSummaryDto{}, err
	}

	return c.summarizeDirect(ctx, conversation, me.ID, otherMemberID), nil
}

/* ---------------------------------- <***> --------------------------------- */

// buildDirectKey is the dedupe key for a 1:1 DM: both ids, sorted so the pair
// yields the same key whichever side opens it.
//
// Only DIRECT conversations get one. A GROUP deliberately does not — a hash of
// the member set goes stale the moment somebody is added, and a stale value
// under a unique index is worse than no dedupe at all.
func buildDirectKey(a, b uuid.UUID) string {
	ids := []string{a.String(), b.String()}
	sort.Strings(ids)
	return ids[0] + ":" + ids[1]
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
