package CryptsSha512

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// Generate a salt string with 16 bytes of crypto/rand data.
func generateSalt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func HmacSha512(message, secretKey []byte) []byte {
	mac := hmac.New(sha512.New, secretKey)
	mac.Write(message)
	return mac.Sum(nil)
}
