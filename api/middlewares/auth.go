package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/constants"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent/usersession"
	"github.com/usegranthq/backend/utils"
)

func unauthorized(c *gin.Context) {
	utils.Http.DeleteCookie(c, constants.AuthCookie)
	utils.HttpError.Unauthorized(c)
	c.Abort()
}

// Auth middleware to check if the user is authenticated from jwt cookie
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionCookie, err := c.Cookie(constants.AuthCookie)
		if err != nil {
			unauthorized(c)
			return
		}

		// Parse the token
		claims, err := jwt.ParseWithClaims(sessionCookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get("JWT_SECRET")), nil
		})

		if err != nil {
			unauthorized(c)
			return
		}

		// Validate token claims
		if !claims.Valid {
			unauthorized(c)
			return
		}

		// Check token expiration
		expirationTime, err := claims.Claims.GetExpirationTime()
		if err != nil || expirationTime.Before(time.Now()) {
			unauthorized(c)
			return
		}

		// Get the user ID from the claims
		sessionID, err := claims.Claims.GetSubject()
		if err != nil {
			unauthorized(c)
			return
		}

		// find the user session
		session, err := db.Client.UserSession.Query().Where(usersession.Token(sessionID)).Only(c)
		if err != nil {
			unauthorized(c)
			return
		}

		// Check session expiration
		if session.ExpiresAt.Before(time.Now()) {
			unauthorized(c)
			return
		}

		// get user id from session
		user, err := session.QueryUser().Only(c)
		if err != nil {
			unauthorized(c)
			return
		}

		c.Set("user", user)
		c.Set("userID", user.ID)
		c.Set("sessionID", session.ID)
		c.Set("session", session)
		c.Next()
	}
}
