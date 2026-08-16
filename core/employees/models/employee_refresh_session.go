package coreEmployeeModels

import (
	"time"

	"github.com/google/uuid"
)

/*
EmployeeRefreshSession is one mobile login — a refresh-token "family".

The refresh token handed to the device carries this row's ID as its "sid"
claim and a per-issue "jti". CurrentJTI records which jti is currently
redeemable; every rotation overwrites it, which retires the previous token
immediately.

That gives reuse detection for free. If a request arrives with a jti that is
well-formed and correctly signed but no longer matches CurrentJTI, that token
was already exchanged — the legitimate client has moved on. Two parties holding
the same refresh token means it leaked, so the whole family is revoked and both
the attacker and the real user are forced back through login.

Rows are not self-expiring the way the Redis-backed keys elsewhere are, so
expired rows are swept periodically — see PurgeExpiredRefreshSessions.
*/
type EmployeeRefreshSession struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID     uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`

	// CurrentJTI is the only refresh token jti this session will accept.
	CurrentJTI uuid.UUID `gorm:"type:uuid;not null"`

	// Role is denormalised from the sentinel role at login so a refresh can
	// mint a correctly-scoped access token without re-resolving it.
	Role string `gorm:"type:varchar(100);not null"`

	ExpiresAt time.Time `gorm:"not null"`

	// RevokedAt is set on logout, on reuse detection, or when an admin kills
	// the session. A non-nil value means the session is dead regardless of
	// ExpiresAt.
	RevokedAt     *time.Time
	RevokedReason string `gorm:"type:varchar(40)"`

	CreatedAt time.Time
	RotatedAt *time.Time
}

func (EmployeeRefreshSession) TableName() string {
	return "employee.employee_refresh_session"
}

// Revocation reasons recorded in RevokedReason. Kept short and stable — they
// are read back in audit logs and support tooling.
const (
	RevokeReasonLogout       = "logout"
	RevokeReasonTokenReuse   = "token_reuse"
	RevokeReasonLogoutAll    = "logout_all"
	RevokeReasonAdminRevoked = "admin_revoked"
)

// IsLive reports whether the session can still be used at the given instant.
func (s EmployeeRefreshSession) IsLive(now time.Time) bool {
	return s.RevokedAt == nil && s.ExpiresAt.After(now)
}
