package organization_services

import (
	"context"

	"github.com/google/uuid"
)

func (r *organizationService) GetSignupSlug(
	ctx context.Context, tempID uuid.UUID,
) (string, error) {

	signupID := tempID.String()
	// 2. Fetch from Redis
	staging, err := r.orgRedisRepo.Get(ctx, signupID)
	if err != nil {
		return "", err
	}

	// Same two failure modes as VerifySlugCreateDefaultOrg, so they carry the
	// same sentinels: the caller has to tell "your signup expired, start over"
	// apart from "Redis is down", and a bare errors.New cannot express that.
	if staging == nil {
		return "", ErrSignupNotFound
	}

	// 3. Return ONLY email
	if staging.Email == "" {
		return "", ErrInvalidSignupData
	}

	return staging.Email, nil
}

// DiscardSignupStaging removes a consumed signup staging record (and its
// email/phone reverse-lookup keys) from Redis.
//
// This is separate from VerifySlugCreateDefaultOrg on purpose: Redis takes no
// part in the database transaction, so the record may only be dropped once that
// transaction has actually committed. Call it after the commit, never inside.
func (r *organizationService) DiscardSignupStaging(
	ctx context.Context, tempID uuid.UUID,
) error {
	return r.orgRedisRepo.Delete(ctx, tempID.String())
}
