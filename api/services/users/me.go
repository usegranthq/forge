package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/ent"
)

func GetUser(c *gin.Context) {
	meUser := c.MustGet("user").(*ent.User)

	c.JSON(http.StatusOK, gin.H{
		"email":      meUser.Email,
		"last_login": meUser.LastLogin,
	})
}
