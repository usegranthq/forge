package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/ent"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/utils"
)

type loginRequest struct {
	Email   string `json:"email" binding:"required"`
	CfToken string `json:"cf_token" binding:"required"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	if err := VerifyCaptcha(c, req.CfToken); err != nil {
		return
	}

	userEmail := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := db.Client.User.
		Query().
		Where(user.Email(userEmail)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			if err := DoEmailSignup(c, userEmail); err != nil {
				return
			}
			c.JSON(http.StatusCreated, gin.H{})
			return
		}
		utils.HttpError.InternalServerError(c)
		return
	}

	StartUserVerification(c, user)
	c.JSON(http.StatusOK, gin.H{})
}
