package external

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/usegranthq/backend/config"
)

type externalTurnstile struct{}

var Turnstile = &externalTurnstile{}

const turnstileURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

type turnstileResponse struct {
	// Success indicates if the challenge was passed
	Success bool `json:"success"`
	// ChallengeTs is the timestamp of the captcha
	ChallengeTs string `json:"challenge_ts"`
	// Hostname is the hostname of the passed captcha
	Hostname string `json:"hostname"`
	// ErrorCodes contains error codes returned by hCaptcha (optional)
	ErrorCodes []string `json:"error-codes"`
	// Action  is the customer widget identifier passed to the widget on the client side
	Action string `json:"action"`
	// CData is the customer data passed to the widget on the client side
	CData string `json:"cdata"`
}

// Verify verifies a "h-captcha-response" data field, with an optional remote IP set.
func (t *externalTurnstile) Verify(token string) (*turnstileResponse, error) {
	client := resty.New()

	request := client.R()
	request.SetFormData(map[string]string{
		"secret":   config.Get("TURNSTILE_SECRET_KEY"),
		"response": token,
	})

	var result turnstileResponse
	request.SetResult(&result)

	resp, err := request.Post(turnstileURL)

	if resp.IsError() {
		return nil, fmt.Errorf("error response from server: %s%s", resp.Status(), resp.Body())
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
