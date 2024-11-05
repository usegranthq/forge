package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
	"go.uber.org/zap"
)

func ValidateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(uuid.UUID)
		projectID := c.Param("projectID")
		l := c.MustGet("logger").(*zap.SugaredLogger)

		proj, err := db.Client.Project.Query().Where(
			project.ID(uuid.MustParse(projectID)),
			project.HasUserWith(user.ID(userID)),
		).Only(c)

		if err != nil {
			if ent.IsNotFound(err) {
				utils.HttpError.NotFound(c, "Project not found")
			} else {
				utils.Log.Errorf("Error getting project: %v", err)
				utils.HttpError.InternalServerError(c)
			}
			return
		}

		pl := l.With(
			"project_id", proj.ID,
		)

		c.Set("project", proj)
		c.Set("projectID", proj.ID)
		c.Set("logger", pl)
		c.Next()
	}
}
