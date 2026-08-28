package sentinel_repo

import (
	"context"
	"fmt"
	"time"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// The attachment registers: which policies reach which principals.
//
// Policies live in ONE table (policy.go). These four tables decide who each one
// applies to, and they are the ONLY thing that decides it — there is no scope
// column and no special case. The invariant is:
//
//	a policy applies to a principal iff a live attachment row reaches them.
//
// All four are structurally identical — (organization_id, <principal>,
// sentinel_policy_id) with soft delete — so they share one implementation here
// and differ only by the row in attachmentTables. Adding a fifth principal kind
// (department, cost centre, ...) is a new table plus one map entry; nothing
// above the repo learns a table name.

// PrincipalKind is what a policy can be attached to.
type PrincipalKind string

const (
	PrincipalOrg   PrincipalKind = "organization" // everyone in the org
	PrincipalRole  PrincipalKind = "role"
	PrincipalGroup PrincipalKind = "group"
	PrincipalUser  PrincipalKind = "user" // a single employee
)

// attachmentTable is the physical register behind one principal kind.
type attachmentTable struct {
	name         string // fully-qualified table
	principalCol string // column holding the principal id
}

var attachmentTables = map[PrincipalKind]attachmentTable{
	// An org edge points back at its own org, so for this one the principal
	// column IS organization_id. The helpers below notice that and avoid
	// emitting the same predicate twice.
	PrincipalOrg:   {"sentinel_engine.org_sentinel_policies", "organization_id"},
	PrincipalRole:  {"sentinel_engine.role_policies", "role_id"},
	PrincipalGroup: {"sentinel_engine.groups_sentinel_policies", "group_id"},
	PrincipalUser:  {"sentinel_engine.user_policies", "employee_id"},
}

// AllPrincipalKinds is the iteration order used wherever every register must be
// consulted (see ListPolicyAttachments, and the engine's union in
// engine_store.go). Broadest reach first.
var AllPrincipalKinds = []PrincipalKind{PrincipalOrg, PrincipalRole, PrincipalGroup, PrincipalUser}

func (k PrincipalKind) table() (attachmentTable, error) {
	t, ok := attachmentTables[k]
	if !ok {
		return attachmentTable{}, fmt.Errorf("sentinel: unknown principal kind %q", k)
	}
	return t, nil
}

/* -------------------------------------------------------------------------- */
/*                            Shared implementation                           */
/* -------------------------------------------------------------------------- */

// attach registers a policy against a principal. Idempotent by the register's
// partial unique index (organization_id, <principal>, sentinel_policy_id) WHERE
// deleted_at IS NULL: re-attaching a live pair does nothing, while re-attaching
// a previously detached pair inserts a fresh row — the soft-deleted one does not
// block it. ON CONFLICT is used rather than a read-then-write so two concurrent
// attaches cannot both pass the check and collide.
func (r *sentinelRepo) attach(ctx context.Context, kind PrincipalKind, orgID, principalID, policyID uuid.UUID) error {
	t, err := kind.table()
	if err != nil {
		return err
	}
	row := map[string]any{
		"id":                 uuid.New(), // app-generated, never a DB default
		"organization_id":    orgID,
		"sentinel_policy_id": policyID,
		"created_at":         time.Now(),
	}
	// For PrincipalOrg this rewrites organization_id with the same value.
	row[t.principalCol] = principalID

	return r.conn(ctx).Table(t.name).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(row).Error
}

/* ---------------------------------- <***> --------------------------------- */

// detach soft-deletes the live registration, leaving the policy row untouched.
// Note that .Table() bypasses GORM's soft-delete scoping, so deleted_at is
// handled explicitly here and in every read below.
func (r *sentinelRepo) detach(ctx context.Context, kind PrincipalKind, orgID, principalID, policyID uuid.UUID) error {
	t, err := kind.table()
	if err != nil {
		return err
	}
	q := r.conn(ctx).Table(t.name).
		Where("organization_id = ?", orgID).
		Where("sentinel_policy_id = ?", policyID).
		Where("deleted_at IS NULL")
	if t.principalCol != "organization_id" {
		q = q.Where(t.principalCol+" = ?", principalID)
	}
	return q.Update("deleted_at", time.Now()).Error
}

/* ---------------------------------- <***> --------------------------------- */

// isAttached reports whether a live registration exists.
func (r *sentinelRepo) isAttached(ctx context.Context, kind PrincipalKind, orgID, principalID, policyID uuid.UUID) (bool, error) {
	t, err := kind.table()
	if err != nil {
		return false, err
	}
	q := r.conn(ctx).Table(t.name).
		Where("organization_id = ?", orgID).
		Where("sentinel_policy_id = ?", policyID).
		Where("deleted_at IS NULL")
	if t.principalCol != "organization_id" {
		q = q.Where(t.principalCol+" = ?", principalID)
	}
	var count int64
	err = q.Count(&count).Error
	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */

// policiesAttachedTo returns the live, non-deleted policies registered against
// one principal. Enabled is NOT filtered here — that is the engine's concern
// (engine_store.go); an admin screen still wants to see disabled policies.
func (r *sentinelRepo) policiesAttachedTo(ctx context.Context, kind PrincipalKind, orgID, principalID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	t, err := kind.table()
	if err != nil {
		return nil, err
	}
	q := r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Select("p.*").
		Joins(fmt.Sprintf("JOIN %s AS att ON att.sentinel_policy_id = p.id AND att.deleted_at IS NULL", t.name)).
		Where("p.organization_id = ? AND p.deleted_at IS NULL", orgID).
		Where("att.organization_id = ?", orgID)
	if t.principalCol != "organization_id" {
		q = q.Where("att."+t.principalCol+" = ?", principalID)
	}
	var policies []*sentinelengine_models.SentinelPolicy
	err = q.Order("p.created_at DESC").Find(&policies).Error
	return policies, err
}

/* ---------------------------------- <***> --------------------------------- */

// principalsAttachedToPolicy is the reverse lookup: which principals of one
// kind a policy currently reaches.
func (r *sentinelRepo) principalsAttachedToPolicy(ctx context.Context, kind PrincipalKind, orgID, policyID uuid.UUID) ([]uuid.UUID, error) {
	t, err := kind.table()
	if err != nil {
		return nil, err
	}
	var ids []uuid.UUID
	err = r.conn(ctx).Table(t.name).
		Where("organization_id = ?", orgID).
		Where("sentinel_policy_id = ?", policyID).
		Where("deleted_at IS NULL").
		Pluck(t.principalCol, &ids).Error
	return ids, err
}

/* -------------------------------------------------------------------------- */
/*                                Organization                                */
/* -------------------------------------------------------------------------- */

// AttachPolicyToOrg makes a policy apply to every employee in the org. This is
// what `scope = 'organization'` used to mean, now expressed as a registration
// that can be detached and re-attached with a history.
//
// There is no principal id: the org IS the principal.
func (r *sentinelRepo) AttachPolicyToOrg(ctx context.Context, orgID, policyID uuid.UUID) error {
	return r.attach(ctx, PrincipalOrg, orgID, orgID, policyID)
}

func (r *sentinelRepo) DetachPolicyFromOrg(ctx context.Context, orgID, policyID uuid.UUID) error {
	return r.detach(ctx, PrincipalOrg, orgID, orgID, policyID)
}

func (r *sentinelRepo) IsPolicyAttachedToOrg(ctx context.Context, orgID, policyID uuid.UUID) (bool, error) {
	return r.isAttached(ctx, PrincipalOrg, orgID, orgID, policyID)
}

// ListPoliciesOfOrg returns the org-wide policies — those that reach everyone.
func (r *sentinelRepo) ListPoliciesOfOrg(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	return r.policiesAttachedTo(ctx, PrincipalOrg, orgID, orgID)
}

/* -------------------------------------------------------------------------- */
/*                                    Role                                    */
/* -------------------------------------------------------------------------- */

func (r *sentinelRepo) AttachPolicyToRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	return r.attach(ctx, PrincipalRole, orgID, roleID, policyID)
}

func (r *sentinelRepo) DetachPolicyFromRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	return r.detach(ctx, PrincipalRole, orgID, roleID, policyID)
}

func (r *sentinelRepo) IsPolicyAttachedToRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) (bool, error) {
	return r.isAttached(ctx, PrincipalRole, orgID, roleID, policyID)
}

func (r *sentinelRepo) ListPoliciesOfRole(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	return r.policiesAttachedTo(ctx, PrincipalRole, orgID, roleID)
}

/* -------------------------------------------------------------------------- */
/*                                    Group                                   */
/* -------------------------------------------------------------------------- */

func (r *sentinelRepo) AttachPolicyToGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) error {
	return r.attach(ctx, PrincipalGroup, orgID, groupID, policyID)
}

func (r *sentinelRepo) DetachPolicyFromGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) error {
	return r.detach(ctx, PrincipalGroup, orgID, groupID, policyID)
}

func (r *sentinelRepo) IsPolicyAttachedToGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) (bool, error) {
	return r.isAttached(ctx, PrincipalGroup, orgID, groupID, policyID)
}

func (r *sentinelRepo) ListPoliciesOfGroup(ctx context.Context, orgID, groupID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	return r.policiesAttachedTo(ctx, PrincipalGroup, orgID, groupID)
}

/* -------------------------------------------------------------------------- */
/*                                    User                                    */
/* -------------------------------------------------------------------------- */

// A direct grant: the org wants one employee to have a policy without giving
// them a role or putting them in a group.

func (r *sentinelRepo) AttachPolicyToUser(ctx context.Context, orgID, userID, policyID uuid.UUID) error {
	return r.attach(ctx, PrincipalUser, orgID, userID, policyID)
}

func (r *sentinelRepo) DetachPolicyFromUser(ctx context.Context, orgID, userID, policyID uuid.UUID) error {
	return r.detach(ctx, PrincipalUser, orgID, userID, policyID)
}

func (r *sentinelRepo) IsPolicyAttachedToUser(ctx context.Context, orgID, userID, policyID uuid.UUID) (bool, error) {
	return r.isAttached(ctx, PrincipalUser, orgID, userID, policyID)
}

func (r *sentinelRepo) ListPoliciesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	return r.policiesAttachedTo(ctx, PrincipalUser, orgID, userID)
}

/* -------------------------------------------------------------------------- */
/*                              Reverse lookup                                */
/* -------------------------------------------------------------------------- */

// ListPolicyAttachments reports every principal a policy currently reaches,
// grouped by kind. This is what replaces the old `scope` label on PolicyView:
// instead of one bucket that could disagree with reality, the UI shows the
// actual registrations — and a policy attached to nothing is visibly
// unattached rather than silently mislabelled.
func (r *sentinelRepo) ListPolicyAttachments(ctx context.Context, orgID, policyID uuid.UUID) (map[PrincipalKind][]uuid.UUID, error) {
	out := make(map[PrincipalKind][]uuid.UUID, len(AllPrincipalKinds))
	for _, kind := range AllPrincipalKinds {
		ids, err := r.principalsAttachedToPolicy(ctx, kind, orgID, policyID)
		if err != nil {
			return nil, fmt.Errorf("sentinel: listing %s attachments of policy %s: %w", kind, policyID, err)
		}
		if len(ids) > 0 {
			out[kind] = ids
		}
	}
	return out, nil
}
