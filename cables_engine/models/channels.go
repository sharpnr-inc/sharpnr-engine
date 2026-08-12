package cables_engine_models

import (
	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// so a channel account gonna have multiple channels
// a organization can as much channels as they want
// Channel is kinda Group

type CablesChannel struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	OrgCablesAccountID uuid.UUID

	Name string

	CreatedAt time.Time
	Active    bool
	UpdatedAt time.Time
	CreatedBy *uuid.UUID
	DeletedAt gorm.DeletedAt
}

func (t *CablesChannel) TableName() string {
	return "cables.cables_channels"
}

/* ---------------------------------- <***> --------------------------------- */

// Now we will create the Configuration Tables for Channels
// each channels gonna have their setting tables
// Note: For now this channel setting tables is kinda empty but i have added some features
// we add the stuff as we move along

type CablesChannelSettings struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	CablesChannelID uuid.UUID

	ChannelHeadline   *string
	ChannelThemeColor *string
	ChannelType       cablesEngineConstants.ChannelType // Voice, Message, live videos, kinda discord

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *CablesChannelSettings) TableName() string {
	return "cables.cables_channel_settings"
}

/* ---------------------------------- <***> --------------------------------- */
