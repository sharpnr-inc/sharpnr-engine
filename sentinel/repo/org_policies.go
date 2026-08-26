package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// these are the methods for the organization Policy Table
// so, these are not just empty, we will circle back.

func (r *sentinelRepo) CreateOrgWidePolicy(ctx context.Context, orgPolicy sentinelengine_models.OrgSentinelPolicy) error {
	panic("")
}

/* ---------------------------------- <***> --------------------------------- */

// this will returns the array or slice on in which we can find the SentinelPolicyID
// which we can use the fetch the Policy from SentielPolicy Table.
func (r *sentinelRepo) GetAllOrgWidePolicesEntry(ctx context.Context, orgID uuid.UUID) ([]sentinelengine_models.OrgSentinelPolicy, error) {
	panic("")
}

/* ---------------------------------- <***> --------------------------------- */
func (r *sentinelRepo) DeleteOrgPolicy(ctx context.Context, orgID uuid.UUID, orgPolicyID uuid.UUID) error {
	panic("")
}
