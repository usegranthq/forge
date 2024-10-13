package main

import (
	"fmt"

	"github.com/usegranthq/backend/utils"
)

func main() {
	// Define the length of the HMAC key in bytes (32 bytes for 256-bit security).
	keyLength := 32

	// Generate a secure random key for HMAC.
	hmacSecretKey, err := utils.GenerateRandom(keyLength)
	if err != nil {
		fmt.Println("Error generating HMAC key:", err)
		return
	}

	fmt.Printf("Generated HMAC key (hex): %s\n", hmacSecretKey)
}
