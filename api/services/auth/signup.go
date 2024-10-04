package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	// check if user already exists
	_, err := db.Client.User.
		Query().
		Where(user.Email(req.Email)).
		Only(c)

	if err == nil {
		utils.HttpError.Conflict(c, "User already exists")
		return
	} else if !ent.IsNotFound(err) {
		utils.HttpError.InternalServerError(c)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	newUser, err := db.Client.User.
		Create().
		SetEmail(req.Email).
		SetPassword(string(hashedPassword)).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": newUser.ID,
	})
}
