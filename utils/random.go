package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func GenerateRandomBytes(options ...int) ([]byte, error) {
	length := 32
	if len(options) > 0 {
		length = options[0]
	}

	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// length is optional, defaults to 32 bytes
func GenerateToken(prefix ...string) (string, error) {
	length := 32

	prefixString := ""
	if len(prefix) > 0 {
		prefixString = prefix[0] + "_"
	}

	key, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}

	return prefixString + base64.RawURLEncoding.EncodeToString(key), nil
}

// length is optional, defaults to 32 bytes
func GenerateRandom(options ...int) (string, error) {
	key, err := GenerateRandomBytes(options...)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}
