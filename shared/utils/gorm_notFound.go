package utils

import "strings"

// Utility to check for 404/Not Found errors
// Adjust this logic to match your specific DB or Error library (e.g., GORM, sql)
func IsNotFoundError(err error) bool {
	// Example: return errors.Is(err, gorm.ErrRecordNotFound)
	// Example: return strings.Contains(err.Error(), "not found")
	return err != nil && strings.Contains(strings.ToLower(err.Error()), "not found")
}
