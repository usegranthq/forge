package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

func ValidateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(uuid.UUID)
		projectID := c.Param("projectID")

		proj, err := db.Client.Project.Query().Where(
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

		c.Set("project", proj)
		c.Set("projectID", proj.ID)
		c.Next()
	}
}
