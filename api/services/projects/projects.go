package projects

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
}

func toProjectResponse(project *ent.Project) projectResponse {
	return projectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func CreateProject(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	type createProjectRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	project, err := db.Client.Project.Create().
		SetUserID(userID).
		SetName(req.Name).
		SetDescription(req.Description).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// register project in oidc provider
	payload := map[string]interface{}{
		"id":     project.ID.String(),
		"domain": "",
	}
	err = external.Oidc.Request("POST", "/projects", payload, nil)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, toProjectResponse(project))
}

func DeleteProject(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)
	currentUser := c.MustGet("user").(*ent.User)

	err := db.Client.Project.DeleteOneID(projectID).
		Where(project.HasUserWith(user.ID(currentUser.ID))).
		Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
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
