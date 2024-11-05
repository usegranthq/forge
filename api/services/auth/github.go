package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

func GithubLogin(c *gin.Context) {
	url, err := external.Github.GenerateOauthUrl()
	if err != nil {
		utils.Log.Errorf("Error generating github oauth url: %v", err)
		utils.HttpError.InternalServerError(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url)
}
