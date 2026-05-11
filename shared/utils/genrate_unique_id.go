// utils/id.go
package utils

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	organization_repo "sharpnr-engine/core/organization/repo"
	"strings"

	"gorm.io/gorm"
)

// CHARSET includes alphanumeric uppercase (excluding confusing letters like I, O, 0)
const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"

func generateRandomString(length int) (string, error) {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		sb.WriteByte(charset[num.Int64()])
	}
	return sb.String(), nil
}

func GenerateUniqueSharpnrOrgID(ctx context.Context, repo organization_repo.OrganizationRepositoryInterface) (string, error) {
	for {
		// 1. Generate the random string (e.g., SHR + 13 chars)
		randomPart, err := generateRandomString(13)
		if err != nil {
			return "", fmt.Errorf("failed to generate random string: %w", err)
		}

		// Ensure the ID is All Caps as per your requirement
		id := strings.ToUpper(fmt.Sprintf("SHR%s", randomPart))

		// 2. Check existence using the Get method
		_, err = repo.GetOrganizationBySharpID(ctx, id)

		if err != nil {
			// If the error is 'Record Not Found', it means the ID is unique!
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return id, nil
			}
			// If it's a different error (like DB connection loss), return it
			return "", err
		}
	}
}
