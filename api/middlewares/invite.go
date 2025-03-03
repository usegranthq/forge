package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/forge/config"
	"github.com/usegranthq/forge/constants"
	"github.com/usegranthq/forge/utils"
)

func ValidateInvite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get invite token from cookie
		inviteToken, err := c.Cookie(constants.InviteCookie)
		if err != nil {
			utils.HttpError.Forbidden(c, "Signup requires invite token")
			c.Abort()
			return
		}

		if inviteToken != config.Get("INVITE_SECRET") {
			utils.HttpError.Forbidden(c, "Invalid invite token")
			c.Abort()
			return
		}

		c.Next()
	}
}
