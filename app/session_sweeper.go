package app

import (
	"context"
	"log/slog"
	"time"

	authenticationEngineServices "sharpnr-engine/authentication/services"
)

// sessionSweepInterval is how often expired mobile refresh sessions are
// deleted. Deliberately a constant rather than an env var: nothing about the
// cadence is environment-specific, and expired rows are already inert — the
// sweep only reclaims space, so the exact period is not load-bearing.
const sessionSweepInterval = 1 * time.Hour

// StartExpiredSessionSweeper deletes expired mobile refresh sessions on a
// loop until ctx is cancelled.
//
// Unlike the Redis-backed keys elsewhere in the system, these rows have no TTL
// and are never removed as a side effect of normal traffic — a revoked or
// expired session is simply skipped by the rotation query. Without this sweep
// the table gains a permanent row per login.
//
// Runs one pass immediately so a restart reclaims anything that expired while
// the process was down.
func StartExpiredSessionSweeper(
	ctx context.Context,
	authSvc authenticationEngineServices.AuthenticationServiceInterface,
) {
	sweep := func() {
		deleted, err := authSvc.PurgeExpiredMobileSessions(ctx)
		if err != nil {
			// Logged, not fatal: a failed sweep costs disk, not correctness,
			// and the next tick retries.
			slog.Error("expired mobile session sweep failed", "error", err.Error())
			return
		}

		if deleted > 0 {
			slog.Info("expired mobile sessions purged", "count", deleted)
		}
	}

	sweep()

	ticker := time.NewTicker(sessionSweepInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			sweep()
		}
	}
}
