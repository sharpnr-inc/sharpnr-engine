CREATE SCHEMA IF NOT EXISTS cables_engine;

-- --------------------------------------
-- cables_engine.org_cables_accounts
-- One organization gets exactly one cables account — it is the root every
-- other cables row hangs off. The table exists (rather than scoping everything
-- straight to organization_id) so an individual, non-org cables account stays
-- possible later without reshaping the whole schema.
--
-- is_active is the kill switch that turns the cables service off for the org
-- without deleting any history. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE cables_engine.org_cables_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial.
CREATE UNIQUE INDEX uq_org_cables_accounts_organization_id
    ON cables_engine.org_cables_accounts(organization_id);

-- --------------------------------------
-- cables_engine.cables_channels
-- The named, discoverable face of a conversation. The channel row holds the
-- identity of the group (name, slug, topic, who can see it); the messages live
-- on the conversation it owns, via conversation_id.
--
-- conversation_id is written in the same transaction that creates the channel
-- and never changes. It has no FK: cables_conversations is not created yet, and
-- the house style keeps referential integrity in the service layer.
--
-- type and visibility are denormalised onto the channel on purpose — the
-- sidebar and channel browser filter on both, and neither should need a join
-- to cables_channel_settings to render a list.
--
-- is_archived is not deletion: an archived channel goes read-only and its
-- history stays searchable. active is the separate "temporarily disabled"
-- switch, and deleted_at is the real soft delete.
--
-- member_count is a denormalised counter kept in sync by the participant
-- service, so a channel list never runs a COUNT(*) per row.
-- --------------------------------------
CREATE TABLE cables_engine.cables_channels (
    id                    UUID PRIMARY KEY,
    organization_id       UUID NOT NULL,
    org_cables_account_id UUID NOT NULL,

    conversation_id       UUID NOT NULL,

    name                  VARCHAR(150) NOT NULL,
    slug                  VARCHAR(150) NOT NULL,
    topic                 TEXT,
    description           TEXT,

    type                  VARCHAR(20) NOT NULL,
    visibility            VARCHAR(20) NOT NULL DEFAULT 'PUBLIC',

    channel_category_id   UUID,
    sort_order            INTEGER NOT NULL DEFAULT 0,

    is_default            BOOLEAN NOT NULL DEFAULT false,

    is_archived           BOOLEAN NOT NULL DEFAULT false,
    archived_at           TIMESTAMPTZ,
    archived_by           UUID,

    member_count          INTEGER NOT NULL DEFAULT 0,

    active                BOOLEAN NOT NULL DEFAULT true,

    created_by            UUID,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now(),
    deleted_at            TIMESTAMPTZ
);

-- The tenant predicate every read carries.
CREATE INDEX idx_cables_channels_organization_id      ON cables_engine.cables_channels(organization_id);

CREATE INDEX idx_cables_channels_account_id           ON cables_engine.cables_channels(org_cables_account_id);
CREATE INDEX idx_cables_channels_category_id          ON cables_engine.cables_channels(channel_category_id);
CREATE INDEX idx_cables_channels_deleted_at           ON cables_engine.cables_channels(deleted_at);

-- The channel browser read: "public, non-archived channels in this org".
CREATE INDEX idx_cables_channels_org_visibility_archived
    ON cables_engine.cables_channels(organization_id, visibility, is_archived)
    WHERE deleted_at IS NULL;

-- @mentions and URLs resolve on slug, so it is unique per cables account among
-- live rows; a soft-deleted channel must not hold its slug hostage.
CREATE UNIQUE INDEX uq_cables_channels_account_slug
    ON cables_engine.cables_channels(org_cables_account_id, slug)
    WHERE deleted_at IS NULL;

-- A channel owns exactly one conversation, and a conversation is owned by at
-- most one channel — the reciprocal of the unique (channel_id) index planned on
-- cables_conversations.
CREATE UNIQUE INDEX uq_cables_channels_conversation_id
    ON cables_engine.cables_channels(conversation_id)
    WHERE deleted_at IS NULL;
