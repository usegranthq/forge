package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

func GoogleLogin(c *gin.Context) {
	url, err := external.Google.GenerateOauthUrl()
	if err != nil {
		utils.Log.Errorf("Error generating google oauth url: %v", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url)
}
