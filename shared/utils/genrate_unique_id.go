// utils/id.go
package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	"strings"
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

func GenerateUniqueSharpnrOrgID(ctx context.Context, repo org_signup_repositories.MainOrgSignUpRepositoryInterface) (string, error) {
	for {
		randomPart, err := generateRandomString(13)
		if err != nil {
			return "", err
		}
		id := fmt.Sprintf("SHR%s", randomPart)

		// Check if it already exists
		_, exists, err := repo.FindOrgBySharpnrID(ctx, id)
		if err != nil {
			return "", err
		}
		if !exists {
			return id, nil
		}
	}
}
