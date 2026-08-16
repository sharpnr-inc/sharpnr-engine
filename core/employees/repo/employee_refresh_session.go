package coreEmployeeRepo

import (
	"context"
	"time"

	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateRefreshSession(
	ctx context.Context,
	session *coreEmployeeModels.EmployeeRefreshSession,
) error {

	if session.ID == uuid.Nil {
		session.ID = uuid.New()
	}

	return r.conn(ctx).WithContext(ctx).Create(session).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetRefreshSessionByID(
	ctx context.Context,
	sessionID uuid.UUID,
) (*coreEmployeeModels.EmployeeRefreshSession, error) {

	var session coreEmployeeModels.EmployeeRefreshSession
	if err := r.conn(ctx).WithContext(ctx).
		Where("id = ?", sessionID).
		First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

/* ---------------------------------- <***> --------------------------------- */

// RotateRefreshSession swaps currentJTI for newJTI, but only if the session is
// still live and expectedJTI is the jti it currently holds. It reports whether
// that swap happened.
//
// The guard clauses live in the WHERE rather than in a preceding SELECT on
// purpose. Two refreshes racing with the same token would both pass a
// read-then-write check and both mint tokens; here the UPDATE is a single
// atomic statement, so exactly one matches a row and the loser gets
// rowsAffected == 0. The caller then reads the session back to find out *why*
// it lost — superseded jti, revoked, or expired — and only the superseded case
// is treated as reuse.
func (r *employeeRepo) RotateRefreshSession(
	ctx context.Context,
	sessionID uuid.UUID,
	expectedJTI uuid.UUID,
	newJTI uuid.UUID,
	now time.Time,
) (bool, error) {

	res := r.conn(ctx).WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeRefreshSession{}).
		Where(
			"id = ? AND current_jti = ? AND revoked_at IS NULL AND expires_at > ?",
			sessionID, expectedJTI, now,
		).
		Updates(map[string]any{
			"current_jti": newJTI,
			"rotated_at":  now,
		})

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected == 1, nil
}

/* ---------------------------------- <***> --------------------------------- */

// RevokeRefreshSession kills a single session. Already-revoked rows are left
// untouched so the original reason and timestamp survive — a logout that
// follows a reuse detection must not overwrite the security-relevant record.
func (r *employeeRepo) RevokeRefreshSession(
	ctx context.Context,
	sessionID uuid.UUID,
	reason string,
	at time.Time,
) error {

	return r.conn(ctx).WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeRefreshSession{}).
		Where("id = ? AND revoked_at IS NULL", sessionID).
		Updates(map[string]any{
			"revoked_at":     at,
			"revoked_reason": reason,
		}).Error
}

/* ---------------------------------- <***> --------------------------------- */

// RevokeAllRefreshSessionsForEmployee kills every live session an employee has.
// Used for "sign out everywhere" and as the blast radius of a reuse detection.
func (r *employeeRepo) RevokeAllRefreshSessionsForEmployee(
	ctx context.Context,
	employeeID uuid.UUID,
	reason string,
	at time.Time,
) (int64, error) {

	res := r.conn(ctx).WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeRefreshSession{}).
		Where("employee_id = ? AND revoked_at IS NULL", employeeID).
		Updates(map[string]any{
			"revoked_at":     at,
			"revoked_reason": reason,
		})

	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

/* ---------------------------------- <***> --------------------------------- */

// PurgeExpiredRefreshSessions deletes rows that expired before the cutoff.
//
// Revoked-but-unexpired rows are deliberately kept: they still need to reject
// any token that is presented against them until natural expiry. Once expired,
// the token would be rejected on its own claims anyway, so the row is dead
// weight.
func (r *employeeRepo) PurgeExpiredRefreshSessions(
	ctx context.Context,
	before time.Time,
) (int64, error) {

	res := r.conn(ctx).WithContext(ctx).
		Where("expires_at < ?", before).
		Delete(&coreEmployeeModels.EmployeeRefreshSession{})

	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
