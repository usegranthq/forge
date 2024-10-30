package public

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

type wellKnownResponse struct {
	Issuer                           string   `json:"issuer"`
	JwksUri                          string   `json:"jwks_uri"`
	SubjectTypesSupported            []string `json:"subject_types_supported"`
	ResponseTypesSupported           []string `json:"response_types_supported"`
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                  []string `json:"claims_supported"`
	ScopesSupported                  []string `json:"scopes_supported"`
}

func getHostFromRequest(request *http.Request) string {
	host := request.Header.Get("X-Forwarded-Host")
	if host == "" {
		host = request.Host
	}

	return host
}

func getHostName(host string) string {
	hostWithOutProtocol := strings.TrimPrefix(host, "https://")

	parsedURL, err := url.Parse("https://" + hostWithOutProtocol)
	if err != nil {
		return ""
	}
	return parsedURL.Hostname()
}

// return .auth.usegrant.local from <PROJECT_URL_ID>.auth.usegrant.local
func getDefaultProjectUrlSuffix() string {
	oidcDomain := config.Get("PROJECT_DEFAULT_DOMAIN")
	hostname := getHostName(oidcDomain)
	return strings.TrimPrefix(hostname, "<PROJECT_URL_ID>")
}

func validateProjectUrlID(c *gin.Context) bool {
	host := getHostFromRequest(c.Request)
	hostname := getHostName(host)
	projectUrlSuffix := getDefaultProjectUrlSuffix()

	if !strings.HasSuffix(hostname, projectUrlSuffix) {
		return false
	}

	projectUrlID := strings.TrimSuffix(hostname, projectUrlSuffix)

	// check if any project has the host
	_, err := db.Client.Project.Query().Where(project.URLID(projectUrlID)).First(c)
	if err != nil {
		return false
	}

	return true
}

func WellKnown(c *gin.Context) {
	if !validateProjectUrlID(c) {
		utils.HttpError.InternalServerError(c, "Invalid host")
		return
	}

	var response interface{}
	err := external.Oidc.Request("GET", "/oauth2/.well-known/openid-configuration", nil, &response)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	host := getHostFromRequest(c.Request)
	hostname := getHostName(host)
	hostWithProtocol := "https://" + hostname

	c.JSON(http.StatusOK, wellKnownResponse{
		Issuer:                           hostWithProtocol,
		JwksUri:                          fmt.Sprintf("%s/.well-known/jwks", hostWithProtocol),
		SubjectTypesSupported:            []string{"public"},
		ResponseTypesSupported:           []string{},
		IdTokenSigningAlgValuesSupported: []string{"RS256"},
		ClaimsSupported:                  []string{"sub", "aud", "exp", "iat", "iss", "jti", "nbf"},
		ScopesSupported:                  []string{"openid"},
	})
}
