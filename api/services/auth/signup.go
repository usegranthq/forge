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

	if err := VerifyCaptcha(c, req.CfToken); err != nil {
		return
	}

	if err := DoEmailSignup(c, req.Email); err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
