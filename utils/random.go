package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// length is optional, defaults to 32 bytes
func GenerateRandom(options ...int) (string, error) {
	length := 32
	if len(options) > 0 {
		length = options[0]
	}

	// Create a byte slice to hold the random key.
	key := make([]byte, length)

	// Fill the byte slice with random bytes.
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a hexadecimal string.
	return hex.EncodeToString(key), nil
}
