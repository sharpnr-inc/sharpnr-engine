package cablesEngineService

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	auditEngineModels "sharpnr-engine/audit_log_engine/models"
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// activeAccount is the guard clause in front of every channel flow: the org has
// a cables account and it is switched on.
//
// The two failures are kept apart on purpose. "Never started" tells the client
// to run the start flow; "disabled" tells it the org switched cables off, which
// is an answer support can act on rather than a dead end.
func (c *cablesService) activeAccount(ctx context.Context, orgID uuid.UUID) (*cables_engine_models.OrgCablesAccount, error) {
	account, err := c.cablesRepo.GetCablesAccountDetails(ctx, orgID)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, cablesEngineErrs.ErrCablesAccountIsDisabled
	}

	return account, nil
}

/* ---------------------------------- <***> --------------------------------- */

// deriveChannelSlug builds the url/mention-safe key from a channel name
// ("Design Team" -> "design-team") for callers, which is all of them: the slug
// is never accepted off the wire.
//
// Uniqueness is not this function's job: the caller checks it against the
// cables account, which is where a collision can actually be seen. A name with
// no slug-able character yields "", which the caller rejects.
func deriveChannelSlug(name string) string {
	var b strings.Builder

	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-' || r == '_' || r == '.':
			// Collapse runs of separators into a single hyphen.
			if slug := b.String(); slug != "" && !strings.HasSuffix(slug, "-") {
				b.WriteRune('-')
			}
		}
	}

	slug := strings.Trim(b.String(), "-")
	if len(slug) > cablesEngineConstants.MaxChannelSlugLength {
		slug = strings.Trim(slug[:cablesEngineConstants.MaxChannelSlugLength], "-")
	}
	return slug
}

/* ---------------------------------- <***> --------------------------------- */

// resolveChannelType and resolveChannelVisibility turn the request's strings
// into the module's enums, defaulting an empty value.
//
// They validate again behind the binding tags rather than trusting them: the
// service is also reachable from the orchestrator, and an enum reaching the
// database unchecked is how a channel nobody can render gets stored.
func resolveChannelType(raw string) (cablesEngineConstants.ChannelType, error) {
	switch cablesEngineConstants.ChannelType(strings.ToUpper(strings.TrimSpace(raw))) {
	case "":
		return cablesEngineConstants.DefaultChannelType, nil
	case cablesEngineConstants.ChannelTypeText:
		return cablesEngineConstants.ChannelTypeText, nil
	case cablesEngineConstants.ChannelTypeAudio:
		return cablesEngineConstants.ChannelTypeAudio, nil
	case cablesEngineConstants.ChannelTypeVideo:
		return cablesEngineConstants.ChannelTypeVideo, nil
	default:
		return "", cablesEngineErrs.ErrInvalidChannelType
	}
}

func resolveChannelVisibility(raw string) (cablesEngineConstants.ChannelVisibility, error) {
	switch cablesEngineConstants.ChannelVisibility(strings.ToUpper(strings.TrimSpace(raw))) {
	case "":
		return cablesEngineConstants.DefaultChannelVisibility, nil
	case cablesEngineConstants.ChannelVisibilityPublic:
		return cablesEngineConstants.ChannelVisibilityPublic, nil
	case cablesEngineConstants.ChannelVisibilityPrivate:
		return cablesEngineConstants.ChannelVisibilityPrivate, nil
	default:
		return "", cablesEngineErrs.ErrInvalidChannelVisibility
	}
}

/* ---------------------------------- <***> --------------------------------- */

// toChannelView maps a stored channel onto the wire shape the frontend reads.
func toChannelView(channel cables_engine_models.CablesChannel) cablesEngineDtos.ChannelResponseDto {
	return cablesEngineDtos.ChannelResponseDto{
		ID:             channel.ID,
		ConversationID: channel.ConversationID,

		Name:        channel.Name,
		Slug:        channel.Slug,
		Topic:       channel.Topic,
		Description: channel.Description,

		Type:       string(channel.Type),
		Visibility: string(channel.Visibility),

		IsDefault:   channel.IsDefault,
		IsArchived:  channel.IsArchived,
		MemberCount: channel.MemberCount,

		CreatedAt: channel.CreatedAt,
		UpdatedAt: channel.UpdatedAt,
	}
}

/* ---------------------------------- <***> --------------------------------- */

// recordAudit stamps one cables event onto the org's history. It is called
// inside the same transaction as the write it describes, so the two commit or
// roll back together.
//
// It carries the resource type because cables writes two kinds of resource —
// the account and the channels under it — and the history reads by type.
func (c *cablesService) recordAudit(
	ctx context.Context,
	orgID, emplID, resourceID uuid.UUID,
	resourceType, action, description string,
	metadata map[string]any,
) error {
	if c.auditLogEngine == nil {
		return nil
	}

	// Protected routes have the client IP stamped onto ctx by the auth
	// middleware; when nothing injected it the field stays empty rather than
	// failing the write.
	var ipAddress string
	if authz, ok := sentinelContexts.Extract(ctx); ok {
		ipAddress = authz.Request.IPAddress
	}

	return c.auditLogEngine.Record(ctx, auditEngineDtos.RecordEventInput{
		OrganizationID: orgID,
		ActorID:        &emplID,
		ActorType:      auditEngineModels.ActorTypeEmployee,
		ActorIPAddress: ipAddress,
		Module:         cablesEngineConstants.AuditModule,
		Action:         action,
		ResourceType:   resourceType,
		ResourceID:     &resourceID,
		Description:    description,
		Metadata:       metadata,
	})
}

/* ---------------------------------- <***> --------------------------------- */

func describeChannel(verb string, channel cables_engine_models.CablesChannel) string {
	return fmt.Sprintf("%s channel #%s (%s) was %s",
		channel.Visibility, channel.Slug, channel.Name, verb)
}

/* -------------------------------------------------------------------------- */
/*                          Members & authorization                           */
/* -------------------------------------------------------------------------- */

// resolveMember maps the employee behind a request to their cables identity.
//
// Everything downstream addresses members, not employees, so this is the first
// step of nearly every flow. A miss means the employee predates the last roster
// sync — Bootstrap is what fixes that, which is why the client calls it first.
func (c *cablesService) resolveMember(ctx context.Context, orgID, emplID uuid.UUID) (*cables_engine_models.CablesMember, error) {
	account, err := c.activeAccount(ctx, orgID)
	if err != nil {
		return nil, err
	}

	return c.cablesRepo.GetMemberByEmployee(ctx, account.ID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

// authorizeConversation is the guard in front of every read and write on a
// conversation, and behind the websocket subscribe as well.
//
// One rule, one place: you may touch a conversation if you hold a live
// membership row in it. Nothing branches on channel visibility here — a private
// channel is invisible because no membership row exists, not because a separate
// check says so.
func (c *cablesService) authorizeConversation(
	ctx context.Context,
	orgID, emplID, conversationID uuid.UUID,
) (*cables_engine_models.CablesMember, *cables_engine_models.CablesConversation, error) {

	me, err := c.resolveMember(ctx, orgID, emplID)
	if err != nil {
		return nil, nil, err
	}

	// Scoped to the org first, so an id from another tenant reads as not found
	// rather than as a membership question.
	conversation, err := c.cablesRepo.GetConversation(ctx, orgID, conversationID)
	if err != nil {
		return nil, nil, err
	}

	if _, err := c.cablesRepo.GetConversationMember(ctx, conversationID, me.ID); err != nil {
		return nil, nil, err
	}

	return me, conversation, nil
}

/* ---------------------------------- <***> --------------------------------- */

// addMember puts one member in a conversation. hidden keeps the row out of
// their sidebar until the first message — see RevealConversation.
func (c *cablesService) addMember(
	ctx context.Context,
	orgID, conversationID, cablesMemberID uuid.UUID,
	role cablesEngineConstants.ParticipantRole,
	addedBy *uuid.UUID,
	hidden bool,
) error {
	now := time.Now()

	return c.cablesRepo.AddConversationMember(ctx, &cables_engine_models.CablesConversationMember{
		ID:                uuid.New(),
		OrganizationID:    orgID,
		ConversationID:    conversationID,
		CablesMemberID:    cablesMemberID,
		Role:              role,
		NotificationLevel: cablesEngineConstants.NotificationLevelAll,
		IsHidden:          hidden,
		AddedAt:           now,
		AddedBy:           addedBy,
		CreatedAt:         now,
		UpdatedAt:         now,
	})
}

/* ---------------------------------- <***> --------------------------------- */

// otherDirectMember resolves the person on the other side of a DM, which is
// what the client renders as its title and avatar.
func (c *cablesService) otherDirectMember(ctx context.Context, conversationID, meID uuid.UUID) (*uuid.UUID, error) {
	members, err := c.cablesRepo.ListConversationMembers(ctx, conversationID)
	if err != nil {
		return nil, err
	}

	for _, m := range members {
		if m.CablesMemberID != meID {
			otherID := m.CablesMemberID
			return &otherID, nil
		}
	}

	// A DM whose counterpart left: the conversation still renders, just without
	// a name to hang on it.
	return nil, nil
}

// summarizeDirect builds the sidebar row for a DM that was just opened. It is
// a fresh conversation, so the counters are all zero rather than read back.
func (c *cablesService) summarizeDirect(
	_ context.Context,
	conversation cables_engine_models.CablesConversation,
	_ uuid.UUID,
	otherMemberID uuid.UUID,
) cablesEngineDtos.ConversationSummaryDto {
	other := otherMemberID

	return cablesEngineDtos.ConversationSummaryDto{
		ID:               conversation.ID,
		Type:             string(conversation.Type),
		OtherMemberID:    &other,
		LastSeq:          conversation.LastSeq,
		LastMessageAt:    conversation.LastMessageAt,
		LastMessageBrief: conversation.LastMessageBrief,
		MemberCount:      conversation.ParticipantCount,
	}
}

/* -------------------------------------------------------------------------- */
/*                                 Wire shapes                                */
/* -------------------------------------------------------------------------- */

func toMemberView(member cables_engine_models.CablesMember) cablesEngineDtos.MemberResponseDto {
	return cablesEngineDtos.MemberResponseDto{
		ID:          member.ID,
		EmployeeID:  member.EmployeeID,
		Type:        string(member.Type),
		DisplayName: member.DisplayName,
		AvatarURL:   member.AvatarURL,
		Presence:    string(member.Presence),
	}
}

// toMessageView maps a stored message onto the wire.
//
// A deleted message keeps its row so seq stays stable and threads do not break,
// so it goes out as a tombstone — the flag set and the body blanked — rather
// than being omitted, which would leave a gap the client reads as missing
// history.
func toMessageView(message cables_engine_models.CablesMessage) cablesEngineDtos.MessageResponseDto {
	view := cablesEngineDtos.MessageResponseDto{
		ID:              message.ID,
		ConversationID:  message.ConversationID,
		Seq:             message.Seq,
		SenderMemberID:  message.SenderMemberID,
		Type:            string(message.Type),
		Body:            message.Body,
		ClientMessageID: message.ClientMessageID,
		IsEdited:        message.IsEdited,
		EditedAt:        message.EditedAt,
		CreatedAt:       message.CreatedAt,
	}

	if message.SubType != nil {
		subType := string(*message.SubType)
		view.SubType = &subType
	}

	if message.DeletedAt.Valid {
		view.IsDeleted = true
		view.Body = ""
	}

	return view
}

/* -------------------------------------------------------------------------- */
/*                                  Realtime                                  */
/* -------------------------------------------------------------------------- */

// cablesEngineErrsNotYourTopic is the deny for a member topic that belongs to
// somebody else. Local to the service because it never leaves the socket path.
var cablesEngineErrsNotYourTopic = errors.New("this topic belongs to another member")

// parseTopicUUID pulls the id out of a "<namespace>:<uuid>" topic. The engine
// has already established the namespace is ours; this validates the rest before
// it reaches a query.
func parseTopicUUID(topic string) (uuid.UUID, error) {
	_, rest, found := strings.Cut(topic, ":")
	if !found {
		return uuid.Nil, cablesEngineErrs.ErrConversationNotFound
	}

	id, err := uuid.Parse(rest)
	if err != nil {
		return uuid.Nil, cablesEngineErrs.ErrConversationNotFound
	}

	return id, nil
}
