package users

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/usegranthq/backend/constants"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/utils"
	"go.uber.org/zap"
)

const sessionRefreshThreshold = 12 * time.Hour

func unauthorized(c *gin.Context) {
	utils.Http.DeleteCookie(c, constants.AuthCookie)
	utils.HttpError.Unauthorized(c)
}

// refresh user session,
// if session is less than 12 hours from expiry, refresh the session,
// else do nothing
func Refresh(c *gin.Context) {
	l := c.MustGet("logger").(*zap.SugaredLogger)
	session, ok := c.MustGet("session").(*ent.UserSession)
	if !ok {
		unauthorized(c)
		return
	}

	// if session expiry is greater than 12 hours from now, do nothing
	if session.ExpiresAt.Sub(time.Now()) > sessionRefreshThreshold {
		return
	}

	// refresh session
	sessionExpiry := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub": session.Token,
		"exp": sessionExpiry.Unix(),
	}

	sessionCookie, err := utils.Jwt.SignToken(claims)
	if err != nil {
		l.Errorf("Error signing session token: %v", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	_, err = session.Update().SetExpiresAt(sessionExpiry).Save(c)
	if err != nil {
		l.Errorf("Error updating session: %v", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	cookieExpiry := int(sessionExpiry.Sub(time.Now()).Seconds())
	utils.Http.SetCookie(c, constants.AuthCookie, sessionCookie, cookieExpiry)
}
