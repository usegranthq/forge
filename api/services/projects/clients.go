package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/oidcclient"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
	"golang.org/x/oauth2/clientcredentials"
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
	userId := c.MustGet("user_id").(uuid.UUID)
	projectID := c.Param("id")

	// validate request
	var req createOidcClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	project, err := db.Client.Project.Query().Where(
		project.ID(uuid.MustParse(projectID)),
		project.HasUserWith(user.ID(userId)),
	).Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c, "Project not found")
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	// create oidc client in hydra
	url := fmt.Sprintf("%s/admin/clients", os.Getenv("HYDRA_ADMIN_URL"))
	payload := map[string]interface{}{
		"audience":                   []string{"sts.amazonaws.com"},
		"client_id":                  req.ClientID,
		"client_secret":              req.ClientSecret,
		"grant_types":                []string{"client_credentials"},
		"token_endpoint_auth_method": "client_secret_post",
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// is http ok
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		utils.HttpError.InternalServerError(c)
		return
	}

	// create oidc client in the database
	client, err := db.Client.OidcClient.Create().
		SetProject(project).
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
	userId := c.MustGet("user_id").(uuid.UUID)
	projectID := c.Param("id")
	clientID := c.Param("client_id")

	// validate request
	var req getTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	_, err := db.Client.Project.Query().Where(
		project.ID(uuid.MustParse(projectID)),
		project.HasUserWith(
			user.ID(userId),
		),
	).Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c, "Project not found")
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	_, err = db.Client.OidcClient.Query().Where(
		oidcclient.ClientID(clientID),
		oidcclient.HasProjectWith(
			project.ID(uuid.MustParse(projectID)),
		),
	).Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c, "Client not found")
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	tokenEndpoint := fmt.Sprintf("%s/oauth2/token", os.Getenv("HYDRA_PUBLIC_URL"))
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: req.ClientSecret,
		TokenURL:     tokenEndpoint,
		EndpointParams: map[string][]string{
			"audience": {"sts.amazonaws.com"},
		},
	}

	token, err := conf.Token(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token.AccessToken,
	})
}
