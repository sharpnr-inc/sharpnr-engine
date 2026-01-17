package organization_repo

import (
	"context"
	"encoding/json"
	"errors"
	"sharpnr-engine/core/organization/models"

	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (rd *orgTempSignUpRepo) Create(
	ctx context.Context,
	tempOrg *models.OrganizationSignUpStaging,
) error {

	if tempOrg.ID == uuid.Nil {
		return errors.New("organization signup staging ID is required")
	}

	if tempOrg.Email == "" && tempOrg.Phone == "" {
		return errors.New("email or phone is required")
	}

	const signupOTPTTL = 40 * time.Minute

	ttl := signupOTPTTL

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

// Get fetches a temporary organization signup record from Redis.
//
// This method is used during the organization onboarding flow to retrieve
// signup data that was previously stored with a TTL (for example, during
// OTP verification or onboarding continuation).
//
// The data is stored in Redis as a JSON payload under the key:
//
//	org:signup:staging:{signupID}
//
// Behavior:
//   - If the key exists, the JSON payload is unmarshaled into
//     models.OrganizationSignUpStaging and returned.
//   - If the key does not exist (expired or never created), the method
//     returns (nil, nil).
//   - Any Redis or JSON unmarshalling error is returned as-is.
//
// Parameters:
//   - ctx: Context for request-scoped cancellation and deadlines.
//   - signupID: Unique identifier for the signup session.
//
// Returns:
//   - *models.OrganizationSignUpStaging: The staged signup data if found.
//   - nil, nil: If the signup session does not exist or has expired.
//   - error: For Redis access or unmarshalling failures.
//
// Example:
//
//	signupID := "7c9f6b8e-2c0b-4d5c-bdf1-9b1e3d41e92a"
//
//	staging, err := tempSignupRepo.Get(ctx, signupID)
//	if err != nil {
//	    return err
//	}
//
//	if staging == nil {
//	    // Signup expired or invalid
//	    return ErrSignupSessionExpired
//	}
//
//	if staging.Otp != inputOtp {
//	    return ErrInvalidOTP
//	}
//
//	// Proceed with organization creation
func (rd *orgTempSignUpRepo) Get(
	ctx context.Context,
	signupID string,
) (*models.OrganizationSignUpStaging, error) {
	if signupID == "" {
		return nil, errors.New("signupID is required")
	}

	key := "org:signup:staging:" + signupID

	val, err := rd.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		// Not found or expired
		return nil, nil
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

// Existence check (used before OTP resend, etc.)
func (rd *orgTempSignUpRepo) Exists(
	ctx context.Context,
	signupID string,
) (bool, models.OrganizationSignUpStaging, error) {

	var staging models.OrganizationSignUpStaging

	if signupID == "" {
		return false, staging, errors.New("signupID is required")
	}

	key := "org:signup:staging:" + signupID

	// 1️⃣ Check existence
	exists, err := rd.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, staging, err
	}

	if exists == 0 {
		return false, staging, nil
	}

	// 2️⃣ Fetch value
	val, err := rd.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// Key vanished between EXISTS and GET (rare but possible)
			return false, staging, nil
		}
		return false, staging, err
	}

	// 3️⃣ Deserialize
	if err := json.Unmarshal([]byte(val), &staging); err != nil {
		return true, staging, err
	}

	return true, staging, nil
}

func (rd *orgTempSignUpRepo) Delete(
	ctx context.Context,
	signupID string,
) error {

	if signupID == "" {
		return errors.New("signupID is required")
	}

	key := "org:signup:staging:" + signupID

	return rd.redis.Del(ctx, key).Err()
}

func (rd *orgTempSignUpRepo) GetByEmailOrPhone(
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
