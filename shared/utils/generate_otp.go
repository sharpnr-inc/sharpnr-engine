package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateOtp() (bool, string, error) {
	// generate a random 6-digit number
	var digits int64
	b := make([]byte, 3) // 3 bytes ~ 24bits which can cover 0-16M
	if _, err := rand.Read(b); err != nil {
		return false, "", err
	}
	digits = int64(b[0])<<16 | int64(b[1])<<8 | int64(b[2])
	otp := fmt.Sprintf("%06d", digits%1000000) // keep it in range 000000-999999

	return true, otp, nil
}
