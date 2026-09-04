-- --------------------------------------
-- cables_engine.cables_members
-- An identity inside cables. Not the employee row: a bot and a meeting guest
-- both need to author a message and appear in a roster, and neither has an
-- employee id, so sender_member_id downstream points here and never needs a
-- nullable-employee branch.
--
-- employee_id is unique per account among live rows, which is what makes the
-- roster sync an upsert rather than a scan. display_name is denormalised from
-- the employee at sync time so rendering a page of messages joins nothing.
--
-- presence is written by the websocket layer, not by a request handler.
-- --------------------------------------
CREATE TABLE cables_engine.cables_members (
    id                    UUID PRIMARY KEY,
    organization_id       UUID NOT NULL,
    org_cables_account_id UUID NOT NULL,

    employee_id           UUID,

    type                  VARCHAR(20) NOT NULL DEFAULT 'EMPLOYEE',

    display_name          VARCHAR(150) NOT NULL,
    avatar_url            TEXT,

    presence              VARCHAR(20) NOT NULL DEFAULT 'OFFLINE',
    last_seen_at          TIMESTAMPTZ,
    custom_status         TEXT,

    is_ephemeral          BOOLEAN NOT NULL DEFAULT false,

    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now(),
    deleted_at            TIMESTAMPTZ
);

CREATE INDEX idx_cables_members_account         ON cables_engine.cables_members(org_cables_account_id);
CREATE INDEX idx_cables_members_organization_id ON cables_engine.cables_members(organization_id);
CREATE INDEX idx_cables_members_deleted_at      ON cables_engine.cables_members(deleted_at);

-- One member row per employee per account. Partial, so a removed member does
-- not block the employee from ever being re-added.
CREATE UNIQUE INDEX uq_cables_members_account_employee
    ON cables_engine.cables_members(org_cables_account_id, employee_id)
    WHERE employee_id IS NOT NULL AND deleted_at IS NULL;

-- --------------------------------------
-- cables_engine.cables_conversations
-- THE message container. A channel is the named, discoverable face of a
-- conversation; a DM is a conversation with no channel row. Everything
-- downstream — messages, membership, unread, calls — keys on conversation_id
-- only, which is what keeps channel chat and DM chat on one code path.
--
-- last_seq is the per-conversation sequence allocator: a send takes
-- last_seq + 1 under an atomic UPDATE ... RETURNING, and unread is then
-- (conversation.last_seq - member.last_read_seq), a subtraction rather than a
-- COUNT(*).
--
-- direct_key is what stops duplicate DMs: the two member ids sorted and joined,
-- unique per account, so "open a DM with X" is an upsert instead of a scan.
-- GROUP conversations deliberately leave it NULL — a hash of the member set
-- goes stale the moment somebody is added, and a stale value under a unique
-- index is worse than no dedupe.
--
-- last_message_brief/at are denormalised so the sidebar renders without joining
-- to messages.
-- --------------------------------------
CREATE TABLE cables_engine.cables_conversations (
    id                    UUID PRIMARY KEY,
    organization_id       UUID NOT NULL,
    org_cables_account_id UUID NOT NULL,

    type                  VARCHAR(20) NOT NULL,
    channel_id            UUID,
    meeting_id            UUID,

    direct_key            TEXT,
    title                 VARCHAR(150),

    last_seq              BIGINT NOT NULL DEFAULT 0,
    last_message_id       UUID,
    last_message_at       TIMESTAMPTZ,
    last_message_brief    TEXT,
    message_count         BIGINT NOT NULL DEFAULT 0,
    participant_count     INTEGER NOT NULL DEFAULT 0,

    active_call_id        UUID,

    created_by            UUID,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now(),
    deleted_at            TIMESTAMPTZ
);

CREATE INDEX idx_cables_conversations_org_type ON cables_engine.cables_conversations(organization_id, type);
CREATE INDEX idx_cables_conversations_account  ON cables_engine.cables_conversations(org_cables_account_id);
CREATE INDEX idx_cables_conversations_deleted  ON cables_engine.cables_conversations(deleted_at);

-- The reciprocal of uq_cables_channels_conversation_id: one conversation per
-- channel, one channel per conversation.
CREATE UNIQUE INDEX uq_cables_conversations_channel_id
    ON cables_engine.cables_conversations(channel_id)
    WHERE channel_id IS NOT NULL AND deleted_at IS NULL;

CREATE UNIQUE INDEX uq_cables_conversations_direct_key
    ON cables_engine.cables_conversations(org_cables_account_id, direct_key)
    WHERE direct_key IS NOT NULL AND deleted_at IS NULL;

-- --------------------------------------
-- cables_engine.cables_conversation_members
-- Membership AND per-member read state in one row, so the sidebar query — "my
-- conversations with their unread badges" — is a single indexed read for one
-- member.
--
-- Removal is soft (removed_at) so old posts and system messages still resolve
-- the member, and the unique index is unconditional: rejoining clears
-- removed_at on the existing row rather than inserting a second one.
-- --------------------------------------
CREATE TABLE cables_engine.cables_conversation_members (
    id                   UUID PRIMARY KEY,
    organization_id      UUID NOT NULL,
    conversation_id      UUID NOT NULL,
    cables_member_id     UUID NOT NULL,

    role                 VARCHAR(20) NOT NULL DEFAULT 'MEMBER',

    last_read_seq        BIGINT NOT NULL DEFAULT 0,
    last_read_message_id UUID,
    last_read_at         TIMESTAMPTZ,
    unread_mention_count INTEGER NOT NULL DEFAULT 0,

    notification_level   VARCHAR(20) NOT NULL DEFAULT 'ALL',
    is_muted             BOOLEAN NOT NULL DEFAULT false,
    muted_until          TIMESTAMPTZ,
    is_pinned            BOOLEAN NOT NULL DEFAULT false,
    is_hidden            BOOLEAN NOT NULL DEFAULT false,

    added_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    added_by             UUID,
    removed_at           TIMESTAMPTZ,
    removed_by           UUID,

    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now(),
    deleted_at           TIMESTAMPTZ
);

CREATE UNIQUE INDEX uq_cables_conversation_members
    ON cables_engine.cables_conversation_members(conversation_id, cables_member_id)
    WHERE deleted_at IS NULL;

-- The sidebar read: everything this member is still in.
CREATE INDEX idx_cables_conversation_members_member
    ON cables_engine.cables_conversation_members(cables_member_id, is_hidden)
    WHERE removed_at IS NULL AND deleted_at IS NULL;

-- The fan-out read: who is in this conversation right now.
CREATE INDEX idx_cables_conversation_members_conversation
    ON cables_engine.cables_conversation_members(conversation_id)
    WHERE removed_at IS NULL AND deleted_at IS NULL;

-- --------------------------------------
-- cables_engine.cables_messages
-- The hot table, shaped around one query:
--
--   SELECT ... WHERE conversation_id = ? AND seq < ? ORDER BY seq DESC LIMIT 50
--
-- seq is per-conversation, allocated from conversations.last_seq inside the
-- send transaction. It is the pagination cursor, the basis of unread counts,
-- and the tie-breaker created_at cannot be when two sends share a millisecond.
--
-- client_message_id is generated by the sender before the request, so a retry
-- over a dropped socket cannot double-post and the optimistic bubble can be
-- reconciled to the real row.
--
-- Deletion is soft and keeps the row so seq stays stable and threads do not
-- break; the service blanks body/content instead.
-- --------------------------------------
CREATE TABLE cables_engine.cables_messages (
    id                  UUID PRIMARY KEY,
    organization_id     UUID NOT NULL,
    conversation_id     UUID NOT NULL,

    seq                 BIGINT NOT NULL,

    sender_member_id    UUID,

    type                VARCHAR(20) NOT NULL DEFAULT 'TEXT',
    sub_type            VARCHAR(40),

    body                TEXT NOT NULL DEFAULT '',
    content             JSONB,

    parent_message_id   UUID,
    reply_to_message_id UUID,
    reply_count         INTEGER NOT NULL DEFAULT 0,
    last_reply_at       TIMESTAMPTZ,
    thread_broadcast    BOOLEAN NOT NULL DEFAULT false,

    client_message_id   TEXT,

    has_attachments     BOOLEAN NOT NULL DEFAULT false,
    attachment_count    INTEGER NOT NULL DEFAULT 0,
    reaction_count      INTEGER NOT NULL DEFAULT 0,
    is_pinned           BOOLEAN NOT NULL DEFAULT false,

    is_edited           BOOLEAN NOT NULL DEFAULT false,
    edited_at           TIMESTAMPTZ,
    deleted_by          UUID,

    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now(),
    deleted_at          TIMESTAMPTZ
);

-- seq is unique within a conversation; this is also the backstop if two sends
-- ever race the allocator.
CREATE UNIQUE INDEX uq_cables_messages_conversation_seq
    ON cables_engine.cables_messages(conversation_id, seq);

-- The message-list read, newest first.
CREATE INDEX idx_cables_messages_conversation_seq_desc
    ON cables_engine.cables_messages(conversation_id, seq DESC)
    WHERE deleted_at IS NULL;

-- Thread view.
CREATE INDEX idx_cables_messages_parent
    ON cables_engine.cables_messages(parent_message_id, seq)
    WHERE parent_message_id IS NOT NULL;

-- Send idempotency: a retried POST resolves to the row already written.
CREATE UNIQUE INDEX uq_cables_messages_client_id
    ON cables_engine.cables_messages(conversation_id, client_message_id)
    WHERE client_message_id IS NOT NULL;

CREATE INDEX idx_cables_messages_organization_id ON cables_engine.cables_messages(organization_id);
