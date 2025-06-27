package utils

import (
	"crypto/rand"
	"math/big"
)

const idCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomSharpnrOrgID() (string, error) {
	b := make([]byte, 13)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(idCharset))))
		if err != nil {
			return "", err
		}
		b[i] = idCharset[n.Int64()]
	}
	return "SHR" + string(b), nil
}
