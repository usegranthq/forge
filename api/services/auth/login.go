package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	user, err := db.Client.User.
		Query().
		Where(user.Email(req.Email)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			utils.HttpError.NotFound(c, "User not found")
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.HttpError.Unauthorized(c, "Invalid password")
		return
	}

	// generate jwt
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	tokenExpiry := time.Now().Add(time.Hour * 24) // 1 day

	// create user session
	_, err = db.Client.UserSession.Create().
		SetUser(user).
		SetToken(token).
		SetExpiresAt(tokenExpiry).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	// set cookie
	cookieExpiry := int(tokenExpiry.Sub(time.Now()).Seconds())
	secure := os.Getenv("NODE_ENV") == "production"
	domain := "/"
	c.SetCookie("auth", token, cookieExpiry, "/", domain, secure, true)

	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}
