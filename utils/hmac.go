package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type HmacUtil struct{}

var Hmac = &HmacUtil{}

const stateValidityDuration = 300 // 5Minutes in seconds

func (u *HmacUtil) GenerateHMACState(secret string) (string, error) {
	hmacSecretKey := []byte(secret)

	// generate a random nonce (16 bytes)
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// get the current timestamp as a string
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// create the message to sign, combining the nonce and timestamp
	message := fmt.Sprintf("%s:%s", base64.URLEncoding.EncodeToString(nonce), timestamp)

	// create an HMAC-SHA256 signature using the HMAC secret key and the message
	h := hmac.New(sha256.New, hmacSecretKey)
	h.Write([]byte(message))
	// convert the signature to a hexadecimal string
	signature := hex.EncodeToString(h.Sum(nil))

	// combine the message and signature, and encode them in base64 for safe URL transmission
	encodedState := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", message, signature)))
	return encodedState, nil
}

func (u *HmacUtil) VerifySecureHMACState(secret, state string) bool {
	hmacSecretKey := []byte(secret)

	// decode the base64-encoded state string
	stateBytes, err := base64.URLEncoding.DecodeString(state)

	if err != nil {
		return false
	}

	// split the decoded state into its constituent parts: message and signature
	parts := strings.SplitN(string(stateBytes), ":", 3)
	if len(parts) != 3 {
		return false
	}

	// extract the message, timestamp and the received HMAC signature
	nonce := parts[0]
	timestamp := parts[1]
	receivedSignature := parts[2]

	// reconstruct the message
	message := fmt.Sprintf("%s:%s", nonce, timestamp)

	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false
	}

	// verify that the state is within the allowed validity duration (e.g., 5 minutes)
	if time.Since(time.Unix(timestampInt, 0)) > time.Duration(stateValidityDuration)*time.Second {
		return false
	}

	// recompute the HMAC using the extracted message and the shared secret key
	h := hmac.New(sha256.New, hmacSecretKey)
	h.Write([]byte(message))
	// convert the computed HMAC to a hexadecimal string
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// compare the recomputed HMAC with the received signature
	// hmac.Equal is used to prevent timing attacks
	return hmac.Equal([]byte(receivedSignature), []byte(expectedSignature))
}
