package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}
