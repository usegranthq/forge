package external

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/usegranthq/forge/config"
)

type externalPostman struct{}

var Postman = &externalPostman{}

func (p *externalPostman) serverURL(path string) string {
	return fmt.Sprintf("%s%s", os.Getenv("POSTMAN_SERVER_URL"), path)
}

func (p *externalPostman) createAuthenticatedRequest() *resty.Request {
	baseURL := p.serverURL("/api")

	client := resty.New()
	client.SetBaseURL(baseURL)

	request := client.R()
	request.SetAuthToken(config.Get("POSTMAN_SERVER_AUTH_TOKEN"))

	return request
}

func (p *externalPostman) Request(method, url string, payload interface{}, responseStruct interface{}) error {
	request := p.createAuthenticatedRequest()
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

func (p *externalPostman) SendLoginEmail(c *gin.Context, to, code string) error {
	payload := map[string]string{
		"to":   to,
		"code": code,
	}

	return p.Request("POST", "/login", payload, nil)
}
