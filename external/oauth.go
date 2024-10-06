package external

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/usegranthq/backend/config"
)

type ExternalOidc struct{}

var Oidc = &ExternalOidc{}

func (o *ExternalOidc) Request(method, url string, payload interface{}, responseStruct interface{}) error {
	client := resty.New()
	client.SetBaseURL(config.Get("OAUTH_SERVER_URL"))

	request := client.R().
		SetAuthToken(config.Get("OAUTH_SERVER_AUTH_TOKEN")).
		SetResult(responseStruct)

	if method == "POST" {
		request.SetBody(payload)
	}

	resp, err := request.Execute(method, url)
	if err != nil {
		return fmt.Errorf("error response from server: %s", resp.Status())
	}

	return nil
}
