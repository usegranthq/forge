package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/constants"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/utils"
)

func GetUser(c *gin.Context) {
	meUser := c.MustGet("user").(*ent.User)

	c.JSON(http.StatusOK, gin.H{
		"email":      meUser.Email,
		"last_login": meUser.LastLogin,
		"uid":        meUser.UID,
	})
}

func DeleteUser(c *gin.Context) {
	meUser := c.MustGet("user").(*ent.User)

	// delete also clients in knox

	err := db.Client.User.DeleteOneID(meUser.ID).Exec(c)
	if err != nil {
		utils.HttpError.InternalServerError(c, err.Error())
		return
	}

	utils.Http.DeleteCookie(c, constants.AuthCookie)
	c.JSON(http.StatusOK, gin.H{})
}
