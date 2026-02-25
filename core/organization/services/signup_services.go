package organization_services

import (
	"context"
	"errors"

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

	if staging == nil {
		return "", errors.New("email not found")
	}

	// 3. Return ONLY email
	if staging.Email == "" {
		return "", errors.New("no email")
	}

	return staging.Email, nil
}
