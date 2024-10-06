package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

type createOidcClientRequest struct {
	Name         string `json:"name" binding:"required"`
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	Audience     string `json:"audience" binding:"required"`
}

type getTokenRequest struct {
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
		"audience":                   []string{"sts.amazonaws.com"},
		"client_id":                  req.ClientID,
		"client_secret":              req.ClientSecret,
		"grant_types":                []string{"client_credentials"},
		"token_endpoint_auth_method": "client_secret_post",
	}

	var response map[string]interface{}
	if err := external.Oidc.Request("POST", "/admin/clients", payload, &response); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// create oidc client in the database
	client, err := db.Client.OidcClient.Create().
		SetProjectID(projectID).
		SetName(req.Name).
		SetClientID(response["client_id"].(string)).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            client.ID,
		"name":          client.Name,
		"client_id":     client.ClientID,
		"client_secret": response["client_secret"],
		"audience":      response["audience"],
	})
}

func GetToken(c *gin.Context) {
	clientID := c.Param("clientID")

	// validate request
	var req getTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	var response map[string]interface{}
	payload := map[string]interface{}{
		"client_id":     clientID,
		"client_secret": req.ClientSecret,
		"grant_type":    "client_credentials",
		"audience":      "sts.amazonaws.com",
	}
	if err := external.Oidc.Request("POST", "/oauth2/token", payload, &response); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": response["access_token"],
	})
}
