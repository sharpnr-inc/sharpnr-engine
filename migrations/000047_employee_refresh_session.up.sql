-- --------------------------------------
-- employee.employee_refresh_session
-- One row per mobile login — a refresh-token "family".
--
-- Rotation overwrites current_jti on every refresh, so only the most recently
-- issued refresh token is accepted. Presenting a superseded jti for a session
-- that is still live means the token was replayed (it was already exchanged),
-- which is the signal that it leaked — the whole family is then revoked.
--
-- role is denormalised here so the refresh endpoint can mint an access token
-- with the correct role without a join back through the sentinel tables.
-- --------------------------------------
CREATE TABLE employee.employee_refresh_session (
    id              UUID PRIMARY KEY,
    employee_id     UUID NOT NULL,
    organization_id UUID NOT NULL,

    current_jti     UUID NOT NULL,
    role            VARCHAR(100) NOT NULL DEFAULT '',

    expires_at      TIMESTAMPTZ NOT NULL,
    revoked_at      TIMESTAMPTZ,
    revoked_reason  VARCHAR(40),

    created_at      TIMESTAMPTZ DEFAULT now(),
    rotated_at      TIMESTAMPTZ
);

-- "revoke every session for this employee" (logout-all, reuse fallout,
-- deactivation) and "list my active devices".
CREATE INDEX idx_employee_refresh_session_employee_id
    ON employee.employee_refresh_session(employee_id);

-- Drives the periodic purge of expired rows. Unlike the Redis-backed keys
-- elsewhere in the system these do not self-expire, so the sweep needs an
-- index to stay cheap as the table grows.
CREATE INDEX idx_employee_refresh_session_expires_at
    ON employee.employee_refresh_session(expires_at);
