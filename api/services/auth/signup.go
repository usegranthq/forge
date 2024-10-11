package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := DoSignup(c, req.Email); err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
