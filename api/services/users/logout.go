package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/forge/constants"
	"github.com/usegranthq/forge/db"
	"github.com/usegranthq/forge/ent"
	"github.com/usegranthq/forge/utils"
	"go.uber.org/zap"
)

func Logout(c *gin.Context) {
	session, _ := c.MustGet("session").(*ent.UserSession)
	l := c.MustGet("logger").(*zap.SugaredLogger)

	// delete the session
	err := db.Client.UserSession.DeleteOneID(session.ID).Exec(c)
	if err != nil {
		l.Errorf("Error deleting session: %v", err)
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	utils.Http.DeleteCookie(c, constants.AuthCookie)
	c.JSON(http.StatusOK, gin.H{})
}
