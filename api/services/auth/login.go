package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

type loginRequest struct {
	Email string `json:"email" binding:"required"`
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
			DoSignup(c, req.Email)
			return
		} else {
			utils.HttpError.InternalServerError(c)
		}
		return
	}

	StartUserVerification(c, user)
}
