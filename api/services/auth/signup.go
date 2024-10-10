package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

type signupRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	DoSignup(c, req.Email)
}

func DoSignup(c *gin.Context, email string) {
	// check if email is disposable
	if utils.IsDisposableEmail(email) {
		utils.HttpError.BadRequest(c, "Disposable emails are not allowed. Note: when you delete your account, we delete everything without a trace.")
		return
	}

	// check if user already exists
	_, err := db.Client.User.
		Query().
		Where(user.Email(email)).
		Only(c)

	if err == nil {
		utils.HttpError.Conflict(c, "User already exists")
		return
	} else if !ent.IsNotFound(err) {
		utils.HttpError.InternalServerError(c)
		return
	}

	newUser, err := db.Client.User.
		Create().
		SetEmail(email).
		Save(c)

	if err != nil {
		utils.HttpError.InternalServerError(c)
		return
	}

	StartUserVerification(c, newUser)
}
