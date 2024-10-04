package webhooks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @todo: setup authentication
func OryTokenHook(c *gin.Context) {
	session := gin.H{
		"access_token": gin.H{
			"iss":    "http://xee:4444",
			"custom": "test",
		},
		"id_token": gin.H{
			"iss": "http://xee:4444",
			"sub": "the-client-id",
		},
	}

	c.JSON(http.StatusOK, gin.H{"session": session})
}
