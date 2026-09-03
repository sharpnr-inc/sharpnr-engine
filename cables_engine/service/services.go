package cablesEngineService

import (
	"context"
	"errors"
	"strings"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	cables_engine_models "sharpnr-engine/cables_engine/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartCablesService implements [CablesEngineServiceInterface].
//
// Provisioning is idempotent by intent rather than by accident: an org that
// already has an account never gets a second one (the unique index on
// organization_id would reject it anyway) — a switched-off account is turned
// back on, and a live one is told it is already live.
func (c *cablesService) StartCablesService(ctx context.Context, orgID, empID uuid.UUID, request cablesEngineDtos.StartCablesServiceRequestsDto) error {
	if !request.AgreeToTermsAndCondtions || !request.StartService {
		return cablesEngineErrs.ErrTermsNotAccepted
	}

	existing, err := c.cablesRepo.GetCablesAccountDetails(ctx, orgID)
	if err != nil && !errors.Is(err, cablesEngineErrs.ErrCablesServiceNotInitialized) {
		return err
	}

	if existing != nil {
		if existing.IsActive {
			return cablesEngineErrs.ErrCablesServiceAlreadyStarted
		}

		return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)

			if err := c.cablesRepo.CablesAccountSwitch(txCtx, orgID, existing.ID, true); err != nil {
				return err
			}

			return c.recordAudit(txCtx, orgID, empID, existing.ID,
				cablesEngineConstants.AuditResourceTypeCablesAccount,
				cablesEngineConstants.AuditActionCablesServiceResumed,
				"Cables service was switched back on for the organization",
				nil,
			)
		})
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		account, err := c.cablesRepo.IntializeOrgCablesAccount(txCtx, orgID)
		if err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, empID, account.ID,
			cablesEngineConstants.AuditResourceTypeCablesAccount,
			cablesEngineConstants.AuditActionCablesServiceStarted,
			"Cables service was started for the organization",
			nil,
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetCablesServiceDetails implements [CablesEngineServiceInterface].
//
// "Never started" is a false/false answer rather than a 404: this is the read
// the client makes before it knows whether the org has cables at all, so a
// missing account is information, not a failure.
func (c *cablesService) GetCablesServiceDetails(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.CablesServiceResponseDto, error) {
	account, err := c.cablesRepo.GetCablesAccountDetails(ctx, orgID)
	if err != nil {
		if errors.Is(err, cablesEngineErrs.ErrCablesServiceNotInitialized) {
			return cablesEngineDtos.CablesServiceResponseDto{}, nil
		}
		return cablesEngineDtos.CablesServiceResponseDto{}, err
	}

	return cablesEngineDtos.CablesServiceResponseDto{
		IsStarted: true,
		IsActive:  account.IsActive,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

// AddNewChannel implements [CablesEngineServiceInterface].
//
// The slug is derived from the name and checked against the account, which
// gives the caller a precise error instead of the partial unique index's
// constraint violation — the index still enforces it underneath as the race
// backstop.
func (c *cablesService) AddNewChannel(ctx context.Context, orgID, emplID uuid.UUID, request cablesEngineDtos.CreateChannelRequestDto) error {
	account, err := c.activeAccount(ctx, orgID)
	if err != nil {
		return err
	}

	name := strings.TrimSpace(request.Name)
	if name == "" {
		return cablesEngineErrs.ErrChannelNameRequired
	}

	// A name made entirely of punctuation ("###") leaves nothing to mention or
	// link, so it is rejected here rather than stored with an empty slug.
	slug := deriveChannelSlug(name)
	if slug == "" {
		return cablesEngineErrs.ErrChannelNameRequired
	}

	channelType, err := resolveChannelType(request.Type)
	if err != nil {
		return err
	}

	visibility, err := resolveChannelVisibility(request.Visibility)
	if err != nil {
		return err
	}

	exists, err := c.cablesRepo.ChannelSlugExists(ctx, account.ID, slug)
	if err != nil {
		return err
	}
	if exists {
		return cablesEngineErrs.ErrChannelAlreadyExists
	}

	channel := cables_engine_models.CablesChannel{
		ID:                 uuid.New(),
		OrganizationID:     orgID,
		OrgCablesAccountID: account.ID,

		// The conversation is the message container this channel owns. Its row
		// lands when cables_conversations exists; the id is minted here so the
		// column — and the unique index over it — is honest in the meantime,
		// and so the client can hold a stable handle from day one.
		ConversationID: uuid.New(),

		Name:       name,
		Slug:       slug,
		Type:       channelType,
		Visibility: visibility,

		Active: true,

		// CreatedBy stays nil: it points at cables_members.id, and members do
		// not exist yet. Writing the employee id into it would put the wrong
		// kind of id in the column; the audit entry below carries who did it
		// until the members table lands.
	}

	if topic := strings.TrimSpace(request.Topic); topic != "" {
		channel.Topic = &topic
	}
	if description := strings.TrimSpace(request.Description); description != "" {
		channel.Description = &description
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.cablesRepo.CreateChannel(txCtx, channel); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, channel.ID,
			cablesEngineConstants.AuditResourceTypeChannel,
			cablesEngineConstants.AuditActionChannelCreated,
			describeChannel("created", channel),
			map[string]any{
				"name":       channel.Name,
				"slug":       channel.Slug,
				"type":       string(channel.Type),
				"visibility": string(channel.Visibility),
			},
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// ListChannels implements [CablesEngineServiceInterface].
//
// Every channel of the account is returned, private ones included: membership
// filtering belongs to the participant table, which does not exist yet. This is
// the admin-side list — see the ⚠ note in the orchestrator.
func (c *cablesService) ListChannels(ctx context.Context, orgID, emplID uuid.UUID) ([]cablesEngineDtos.ChannelResponseDto, error) {
	account, err := c.activeAccount(ctx, orgID)
	if err != nil {
		return nil, err
	}

	channels, err := c.cablesRepo.ListChannels(ctx, orgID, account.ID)
	if err != nil {
		return nil, err
	}

	views := make([]cablesEngineDtos.ChannelResponseDto, 0, len(channels))
	for _, channel := range channels {
		views = append(views, toChannelView(channel))
	}

	return views, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetChannelDetails implements [CablesEngineServiceInterface].
func (c *cablesService) GetChannelDetails(ctx context.Context, orgID, emplID, channelID uuid.UUID) (cablesEngineDtos.ChannelResponseDto, error) {
	if _, err := c.activeAccount(ctx, orgID); err != nil {
		return cablesEngineDtos.ChannelResponseDto{}, err
	}

	channel, err := c.cablesRepo.GetChannelByID(ctx, orgID, channelID)
	if err != nil {
		return cablesEngineDtos.ChannelResponseDto{}, err
	}

	return toChannelView(*channel), nil
}
