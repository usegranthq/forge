package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/oidcclient"
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/utils"
	"go.uber.org/zap"
)

func ValidateClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.MustGet("projectID").(uuid.UUID)
		clientID := c.Param("clientID")
		l := c.MustGet("logger").(*zap.SugaredLogger)

		client, err := db.Client.OidcClient.Query().Where(
			oidcclient.ID(uuid.MustParse(clientID)),
			oidcclient.HasProjectWith(project.ID(projectID)),
		).Only(c)

		if err != nil {
			if ent.IsNotFound(err) {
				utils.HttpError.NotFound(c, "Client not found")
			} else {
				utils.Log.Errorf("Error getting client: %v", err)
				utils.HttpError.InternalServerError(c)
			}
			return
		}

		cl := l.With(
			"client_id", client.ID,
		)

		c.Set("client", client)
		c.Set("clientID", client.ID)
		c.Set("logger", cl)
		c.Next()
	}
}
