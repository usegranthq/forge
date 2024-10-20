package projects

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/oidcclient"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

type clientResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	Audience  string    `json:"audience"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toClientResponse(client *ent.OidcClient) clientResponse {
	return clientResponse{
		ID:        client.ID.String(),
		Name:      client.Name,
		Subject:   client.ClientID,
		Audience:  client.Audience,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
}

func CreateOidcClient(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)

	type createOidcClientRequest struct {
		Name     string `json:"name" binding:"required"`
		Audience string `json:"audience" binding:"required"`
	}

	type createOidcClientResponse struct {
		ID           string `json:"id"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}

	// validate request
	var req createOidcClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	payload := map[string]interface{}{
		"audience": req.Audience,
	}

	var registerUrl = fmt.Sprintf("projects/%s/clients", projectID.String())

	var response createOidcClientResponse
	if err := external.Oidc.Request("POST", registerUrl, payload, &response); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// create oidc client in the database
	client, err := db.Client.OidcClient.Create().
		SetProjectID(projectID).
		SetName(req.Name).
		SetAudience(req.Audience).
		SetClientRefID(response.ID).
		SetClientID(response.ClientID).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, toClientResponse(client))
}

func GetToken(c *gin.Context) {
	clientID := c.MustGet("clientID").(uuid.UUID)
	projectID := c.MustGet("projectID").(uuid.UUID)

	var response struct {
		AccessToken string `json:"access_token"`
	}

	oidcClient, err := db.Client.OidcClient.Query().
		Where(
			oidcclient.ID(clientID),
			oidcclient.HasProjectWith(project.ID(projectID)),
		).
		Only(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	if err := external.Oidc.RequestToken(c, oidcClient.ClientRefID, &response); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": response.AccessToken,
	})
}

func ListOidcClients(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)

	clients, err := db.Client.OidcClient.
		Query().
		Where(oidcclient.HasProjectWith(project.ID(projectID))).
		All(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	response := make([]clientResponse, len(clients))
	for i, client := range clients {
		response[i] = toClientResponse(client)
	}

	c.JSON(http.StatusOK, response)
}

func DeleteOidcClient(c *gin.Context) {
	clientID := c.MustGet("clientID").(uuid.UUID)
	projectID := c.MustGet("projectID").(uuid.UUID)

	oidcClient, err := db.Client.OidcClient.Get(c, clientID)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	var registerUrl = fmt.Sprintf("projects/%s/clients/%s", projectID.String(), oidcClient.ClientRefID)

	if err := external.Oidc.Request("DELETE", registerUrl, nil, nil); err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	err = db.Client.OidcClient.DeleteOneID(clientID).
		Where(oidcclient.HasProjectWith(project.ID(projectID))).
		Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
