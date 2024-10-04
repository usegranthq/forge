package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent/usersession"
	"github.com/usegranthq/backend/utils"
)

// Auth middleware to check if the user is authenticated from jwt cookie
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth")
		if err != nil {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Parse the token
		claims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Validate token claims
		if !claims.Valid {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Check token expiration
		expirationTime, err := claims.Claims.GetExpirationTime()
		if err != nil || expirationTime.Before(time.Now()) {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Get the user ID from the claims
		userID, err := claims.Claims.GetSubject()
		if err != nil {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Get the user from the database
		user, err := db.Client.User.Get(c, uuid.MustParse(userID))
		if err != nil {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Get the session from the database
		session, err := db.Client.UserSession.Query().Where(usersession.Token(token)).Only(c)
		if err != nil {
			utils.HttpError.Unauthorized(c)
			return
		}

		// Check session expiration
		if session.ExpiresAt.Before(time.Now()) {
			utils.HttpError.Unauthorized(c)
			return
		}

		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Next()
	}
}
