package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/config"
)

type HttpUtil struct{}

var Http = HttpUtil{}

func (h HttpUtil) SetCookie(c *gin.Context, name, value string, maxAge int) {
	secure := config.Get("NODE_ENV") != "development"
	c.SetCookie(name, value, maxAge, "/", "", secure, true)
}

func (h HttpUtil) DeleteCookie(c *gin.Context, name string) {
	secure := config.Get("NODE_ENV") != "development"
	c.SetCookie(name, "", -1, "/", "", secure, true)
}
