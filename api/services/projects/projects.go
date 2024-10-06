package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/utils"
)

type createProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateProject(c *gin.Context) {
	userId := c.MustGet("userId").(uuid.UUID)

	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	project, err := db.Client.Project.Create().
		SetUserID(userId).
		SetName(req.Name).
		SetDescription(req.Description).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
	})
}

func DeleteProject(c *gin.Context) {
	projectID := c.MustGet("projectID").(uuid.UUID)

	err := db.Client.Project.DeleteOneID(projectID).Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
