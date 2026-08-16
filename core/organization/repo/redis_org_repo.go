package organization_repo

import (
	"context"
	"encoding/json"
	"errors"
	organizationErrs "sharpnr-engine/core/organization/errs"
	"sharpnr-engine/core/organization/models"

	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (rd *organizationRedisRepo) Create(
	ctx context.Context,
	tempOrg *models.OrganizationSignUpStaging,
) error {

	if tempOrg.ID == uuid.Nil {
		return errors.New("organization signup staging ID is required")
	}

	if tempOrg.Email == "" && tempOrg.Phone == "" {
		return errors.New("email or phone is required")
	}

	ttl := rd.stagingTTL

	key := "org:signup:staging:" + tempOrg.ID.String()

	payload, err := json.Marshal(tempOrg)
	if err != nil {
		return err
	}

	pipe := rd.redis.TxPipeline()

	// Primary object
	pipe.Set(ctx, key, payload, ttl)

	// Reverse lookups (TTL must match)
	if tempOrg.Email != "" {
		emailKey := "org:signup:staging:email:" + strings.ToLower(tempOrg.Email)
		pipe.Set(ctx, emailKey, tempOrg.ID.String(), ttl)
	}

	if tempOrg.Phone != "" {
		phoneKey := "org:signup:staging:phone:" + tempOrg.Phone
		pipe.Set(ctx, phoneKey, tempOrg.ID.String(), ttl)
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// Proceed with organization creation
func (rd *organizationRedisRepo) Get(
	ctx context.Context,
	signupID uuid.UUID,
) (*models.OrganizationSignUpStaging, error) {

	key := "org:signup:staging:" + signupID.String()

	val, err := rd.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, organizationErrs.SignupSessionExpiredErr
	}
	if err != nil {
		return nil, err
	}

	var tempOrg models.OrganizationSignUpStaging
	if err := json.Unmarshal([]byte(val), &tempOrg); err != nil {
		return nil, err
	}

	return &tempOrg, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (rd *organizationRedisRepo) Delete(
	ctx context.Context,
	signupID uuid.UUID,
) error {

	keys := []string{"org:signup:staging:" + signupID.String()}

	staging, err := rd.Get(ctx, signupID)
	if err != nil {
		if errors.Is(err, organizationErrs.SignupSessionExpiredErr) {
			// Already gone/expired. Nothing to delete.
			return nil
		}
		return err
	}

	if staging.Email != "" {
		keys = append(keys, "org:signup:staging:email:"+strings.ToLower(staging.Email))
	}
	if staging.Phone != "" {
		keys = append(keys, "org:signup:staging:phone:"+staging.Phone)
	}

	return rd.redis.Del(ctx, keys...).Err()
}

/* ---------------------------------- <***> --------------------------------- */

func (rd *organizationRedisRepo) GetByEmailOrPhone(
	ctx context.Context,
	email string,
	phone string,
) (bool, error) {

	if email == "" && phone == "" {
		return false, errors.New("email or phone is required")
	}

	keys := make([]string, 0, 2)

	if email != "" {
		keys = append(
			keys,
			"org:signup:staging:email:"+strings.ToLower(email),
		)
	}

	if phone != "" {
		keys = append(
			keys,
			"org:signup:staging:phone:"+phone,
		)
	}

	count, err := rd.redis.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
