package external

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/usegranthq/backend/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ExternalOidc struct{}

var Oidc = &ExternalOidc{}

func serverURL(path string) string {
	return fmt.Sprintf("%s%s", os.Getenv("OAUTH_SERVER_URL"), path)
}

func createAuthenticatedRequest() *resty.Request {
	baseURL := serverURL("/")

	client := resty.New()
	client.SetBaseURL(baseURL)

	request := client.R()
	request.SetAuthToken(config.Get("OAUTH_SERVER_AUTH_TOKEN"))

	return request
}

func (o *ExternalOidc) Request(method, url string, payload interface{}, responseStruct interface{}) error {
	request := createAuthenticatedRequest()
	request.SetResult(responseStruct)

	if method == "POST" {
		request.SetBody(payload)
	}

	resp, err := request.Execute(method, url)
	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("error response from server: %s%s", resp.Status(), resp.Body())
	}

	return nil
}

func (o *ExternalOidc) RequestToken(c *gin.Context, clientID, clientSecret string, audience []string) (string, error) {
	tokenEndpoint := serverURL("/oauth2/token")

	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenEndpoint,
		AuthStyle:    oauth2.AuthStyleInHeader,
		EndpointParams: map[string][]string{
			"audience": {"sts.amazonaws.com"},
		},
	}

	token, err := conf.Token(c)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
