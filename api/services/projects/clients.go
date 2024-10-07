package projects

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

type createOidcClientRequest struct {
	Name     string `json:"name" binding:"required"`
	Audience string `json:"audience" binding:"required"`
}

type createOidcClientResponse struct {
	ClientID                string `json:"client_id"`
	ClientSecret            string `json:"client_secret"`
	RegistrationClientURI   string `json:"registration_client_uri"`
	RegistrationAccessToken string `json:"registration_access_token"`
}

type getTokenRequest struct {
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

func CreateOidcClient(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)

	// validate request
	var req createOidcClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	payload := map[string]interface{}{
		"client_name":                req.Name,
		"grant_types":                []string{"client_credentials"},
		"token_endpoint_auth_method": "client_secret_basic",
		"redirect_uris":              []string{"https://placeholder.usegranthq.com/callback"},
		"response_types":             []string{"none"},
		"audience":                   req.Audience,
	}

	var response createOidcClientResponse
	if err := external.Oidc.Request("POST", "oauth2/reg", payload, &response); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// create oidc client in the database
	client, err := db.Client.OidcClient.Create().
		SetProjectID(projectID).
		SetName(req.Name).
		SetClientID(response.ClientID).
		SetClientSecret(response.ClientSecret).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            client.ID,
		"name":          client.Name,
		"client_id":     client.ClientID,
		"client_secret": response.ClientSecret,
	})
}

func GetToken(c *gin.Context) {
	// validate request
	var req getTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	token, err := external.Oidc.RequestToken(c, req.ClientID, req.ClientSecret)
	if err != nil {
		fmt.Println("error getting token", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
