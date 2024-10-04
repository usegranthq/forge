package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api/services/auth"
	"github.com/usegranthq/backend/api/services/projects"
	"github.com/usegranthq/backend/api/services/webhooks"
	"github.com/usegranthq/backend/middlewares"
)

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func defineOtherRoutes(router *gin.RouterGroup) {
	router.GET("/health", healthCheckHandler)
	router.POST("/ory-token-hook", webhooks.OryTokenHook)
}

func defineAuthRoutes(router *gin.RouterGroup) {
	router.POST("/signup", auth.Signup)
	router.POST("/login", auth.Login)
}

func defineProjectsRoutes(router *gin.RouterGroup) {
	router.POST("/projects", projects.CreateProject)
	router.DELETE("/projects/:id", projects.DeleteProject)
	router.POST("/projects/:id/clients", projects.CreateOidcClient)
	router.GET("/projects/:id/clients/:client_id/token", projects.GetToken)
}

func SetupRoutes(router *gin.Engine) {
	router.GET("/", rootHandler)

	apiRouter := router.Group("/api")
	apiV1Router := apiRouter.Group("/v1")

	protectedRouter := apiV1Router.Group("/")
	protectedRouter.Use(middlewares.Auth())

	defineOtherRoutes(apiV1Router)
	defineAuthRoutes(apiV1Router)
	defineProjectsRoutes(protectedRouter)
}
