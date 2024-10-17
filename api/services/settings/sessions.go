package settings

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/ent/usersession"
	"github.com/usegranthq/backend/utils"
)

type listSessionsResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func ListSessions(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)
	currentSessionID := c.MustGet("sessionID").(uuid.UUID)

	sessions, err := db.Client.UserSession.Query().
		Where(
			usersession.HasUserWith(user.ID(currentUser.ID)),
			usersession.Not(usersession.ID(currentSessionID)),
		).
		Order(ent.Desc(usersession.FieldCreatedAt)).
		All(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	response := make([]listSessionsResponse, len(sessions))
	for i, session := range sessions {
		response[i] = listSessionsResponse{
			ID:        session.ID,
			CreatedAt: session.CreatedAt,
			ExpiresAt: session.ExpiresAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

func DeleteSession(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)

	sessionID := c.Param("sessionID")
	if sessionID == "" {
		utils.HttpError.BadRequest(c, "Invalid session ID")
		return
	}

	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		utils.HttpError.BadRequest(c, "Invalid session ID")
		return
	}

	session, err := db.Client.UserSession.Query().
		Where(
			usersession.ID(parsedSessionID),
			usersession.HasUserWith(user.ID(currentUser.ID)),
		).
		Only(c)
	if err != nil {
		utils.HttpError.NotFound(c, "Session not found")
		return
	}

	_, err = db.Client.UserSession.Delete().Where(usersession.ID(session.ID)).Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteAllOtherSessions(c *gin.Context) {
	currentUser := c.MustGet("user").(*ent.User)
	currentSessionID := c.MustGet("sessionID").(uuid.UUID)

	_, err := db.Client.UserSession.Delete().
		Where(
			usersession.HasUserWith(user.ID(currentUser.ID)),
			usersession.Not(usersession.ID(currentSessionID)),
		).
		Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
