package public

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func WellKnown(c *gin.Context) {
	// host url
	// @TODO: set correct oauth url as host
	// first check if the host is registered to any projects, if yes process that
	// if registered, but subscription is expired, return 403
	// else return the env url
	host := c.Request.Host

	var response interface{}
	err := external.Oidc.Request("GET", "/oauth2/.well-known/openid-configuration", nil, &response)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, wellKnownResponse{
		Issuer:                           host,
		JwksUri:                          fmt.Sprintf("%s/jwks", host),
		SubjectTypesSupported:            []string{"public"},
		ResponseTypesSupported:           []string{},
		IdTokenSigningAlgValuesSupported: []string{"RS256"},
		ClaimsSupported:                  []string{"sub", "aud", "exp", "iat", "iss", "jti", "nbf"},
		ScopesSupported:                  []string{"openid"},
	})
}
