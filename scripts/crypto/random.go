package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// generateHMACKey generates a cryptographically secure random key of the specified length.
func generateHMACKey(length int) (string, error) {
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

func main() {
	// Define the length of the HMAC key in bytes (32 bytes for 256-bit security).
	keyLength := 32

	// Generate a secure random key for HMAC.
	hmacSecretKey, err := generateHMACKey(keyLength)
	if err != nil {
		fmt.Println("Error generating HMAC key:", err)
		return
	}

	fmt.Printf("Generated HMAC key (hex): %s\n", hmacSecretKey)
}
