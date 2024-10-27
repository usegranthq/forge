package projects

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/sqids/sqids-go"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

type projectResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	URLID       string    `json:"url_id"`
}

func toProjectResponse(project *ent.Project) projectResponse {
	return projectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		URLID:       project.URLID,
	}
}

var CreateProject = db.GinHandlerWithTx(createProjectHandler)
var DeleteProject = db.GinHandlerWithTx(deleteProjectHandler)
var UpdateProject = db.GinHandlerWithTx(updateProjectHandler)

func createProjectHandler(c *gin.Context, tx *ent.Tx) error {
	userID := c.MustGet("userID").(uuid.UUID)

	type createProjectRequest struct {
		Name        string `json:"name" binding:"required,min=3,max=32"`
		Description string `json:"description" binding:"required,min=3,max=100"`
	}

	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return err
	}

	sqids, _ := sqids.New(
		sqids.Options{
			MinLength: 10,
			Alphabet:  "abcdefghijklmnopqrstuvwxyz1234567890",
		},
	)
	urlUniq, _ := sqids.Encode([]uint64{uint64(time.Now().Unix())})
	urlID := slug.Make(req.Name) + "-" + urlUniq

	project, err := tx.Project.Create().
		SetUserID(userID).
		SetName(req.Name).
		SetDescription(req.Description).
		SetURLID(urlID).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	domain := config.Get("PROJECT_DEFAULT_DOMAIN")
	domain = strings.Replace(domain, "<PROJECT_URL_ID>", urlID, 1)

	// register project in oidc provider
	payload := map[string]interface{}{
		"id":     project.ID.String(),
		"domain": domain,
	}
	err = external.Oidc.Request("POST", "/projects", payload, nil)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	c.JSON(http.StatusCreated, toProjectResponse(project))
	return nil
}

func deleteProjectHandler(c *gin.Context, tx *ent.Tx) error {
	projectID := c.MustGet("projectID").(uuid.UUID)
	currentUser := c.MustGet("user").(*ent.User)

	err := tx.Project.DeleteOneID(projectID).
		Where(project.HasUserWith(user.ID(currentUser.ID))).
		Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	err = external.Oidc.Request("DELETE", "/projects/"+projectID.String(), nil, nil)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	c.JSON(http.StatusNoContent, gin.H{})
	return nil
}

func updateProjectHandler(c *gin.Context, tx *ent.Tx) error {
	projectID := c.MustGet("projectID").(uuid.UUID)
	currentUser := c.MustGet("user").(*ent.User)

	type updateProjectRequest struct {
		Name        string `json:"name" binding:"required,min=3,max=32"`
		Description string `json:"description" binding:"required,min=3,max=100"`
	}

	var req updateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return err
	}

	project, err := tx.Project.
		UpdateOneID(projectID).
		Where(project.HasUserWith(user.ID(currentUser.ID))).
		SetName(req.Name).
		SetDescription(req.Description).
		Save(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	payload := map[string]interface{}{
		"id":     project.ID.String(),
		"domain": project,
	}
	err = external.Oidc.Request("PUT", "/projects/"+projectID.String(), payload, nil)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return err
	}

	c.JSON(http.StatusOK, toProjectResponse(project))
	return nil
}

func ListProjects(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)

	projects, err := db.Client.Project.Query().
		Where(project.HasUserWith(user.ID(currentUser.ID))).
		All(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	projectResponses := make([]projectResponse, len(projects))
	for i, project := range projects {
		projectResponses[i] = toProjectResponse(project)
	}

	c.JSON(http.StatusOK, projectResponses)
}

func GetProject(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)
	currentUser := c.MustGet("user").(*ent.User)

	project, err := db.Client.Project.Query().
		Where(
			project.ID(projectID),
			project.HasUserWith(user.ID(currentUser.ID)),
		).
		Only(c)
	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c)
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	c.JSON(http.StatusOK, toProjectResponse(project))
}
