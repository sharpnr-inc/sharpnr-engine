package cablesEngineService

import (
	"context"
	"fmt"
	"strings"

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
