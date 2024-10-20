package external

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/usegranthq/backend/config"
)

type externalOidc struct{}

var Oidc = &externalOidc{}

func (o *externalOidc) serverURL(path string) string {
	return fmt.Sprintf("%s%s", os.Getenv("OAUTH_SERVER_URL"), path)
}

func (o *externalOidc) createAuthenticatedRequest() *resty.Request {
	baseURL := o.serverURL("/")

	client := resty.New()
	client.SetBaseURL(baseURL)

	request := client.R()
	request.SetAuthToken(config.Get("OAUTH_SERVER_AUTH_TOKEN"))

	return request
}

func (o *externalOidc) Request(method, url string, payload interface{}, responseStruct interface{}) error {
	request := o.createAuthenticatedRequest()
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

func (o *externalOidc) RequestToken(c *gin.Context, clientID string, responseStruct interface{}) error {
	request := o.createAuthenticatedRequest()
	request.SetResult(responseStruct)
	request.SetFormData(map[string]string{
		"id": clientID,
	})

	resp, err := request.Execute("POST", "/oauth2/token")
	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("error response from server: %s%s", resp.Status(), resp.Body())
	}

	return nil

}
