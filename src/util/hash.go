package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// HmacSHA256 hash to hexits
func HmacSHA256(secret, message []byte) string {
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	// to lowercase hexits
	return hex.EncodeToString(hash.Sum(nil))
}
