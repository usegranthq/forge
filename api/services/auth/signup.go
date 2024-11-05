package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/utils"
)

type signupRequest struct {
	Email   string `json:"email" binding:"required,email"`
	CfToken string `json:"cf_token" binding:"required"`
}

func Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
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

	if err := DoEmailSignup(c, req.Email); err != nil {
		l.Errorf("Error doing email signup: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
