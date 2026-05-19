package app

import "gorm.io/gorm"

func bootStrapSentinelEngine(db *gorm.DB) string {
	return "Nothing" // No dependencies to initialize for Sentinel Engine at this time.
}
