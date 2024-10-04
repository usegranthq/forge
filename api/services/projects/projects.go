package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

type createProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateProject(c *gin.Context) {
	user := c.MustGet("user").(*ent.User)

	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	project, err := db.Client.Project.Create().
		SetUser(user).
		SetName(req.Name).
		SetDescription(req.Description).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
	})
}

func DeleteProject(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	projectID := c.Param("id")

	project, err := db.Client.Project.Query().Where(
		project.ID(uuid.MustParse(projectID)),
		project.HasUserWith(user.ID(userID)),
	).Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c, "Project not found")
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	err = db.Client.Project.DeleteOne(project).Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
