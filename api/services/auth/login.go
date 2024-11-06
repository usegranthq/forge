package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/forge/db"
	"github.com/usegranthq/forge/ent"
	"github.com/usegranthq/forge/ent/user"
	"github.com/usegranthq/forge/utils"
)

type loginRequest struct {
	Email   string `json:"email" binding:"required"`
	CfToken string `json:"cf_token" binding:"required"`
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Log.Errorf("Error binding login request: %v", err)
		utils.HttpError.BadRequest(c, err.Error())
		return
	}

	l := utils.Log.With(
		"user_email", req.Email,
	)

	if err := VerifyCaptcha(c, req.CfToken); err != nil {
		l.Errorf("Error verifying captcha: %v", err)
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
				l.Errorf("Error doing email signup: %v", err)
				return
			}
			c.JSON(http.StatusCreated, gin.H{})
			return
		}
		l.Errorf("Error getting user: %v", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	StartUserVerification(c, user)
	c.JSON(http.StatusOK, gin.H{})
}
