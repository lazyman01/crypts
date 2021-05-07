package CryptsSha256

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HmacSha256
func HmacSha256(message, privateKey []byte) []byte {
	mac := hmac.New(sha256.New, privateKey)
	mac.Write(message)
	return mac.Sum(nil)
}