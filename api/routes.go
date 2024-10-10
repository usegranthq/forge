package api

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api/middlewares"
	"github.com/usegranthq/backend/api/services/auth"
	"github.com/usegranthq/backend/api/services/projects"
	"github.com/usegranthq/backend/api/services/public"
)

func defineProjectRoutes(routerGroup *gin.RouterGroup) {
	projectGroup := routerGroup.Group("/projects")
	projectGroup.POST("/", projects.CreateProject)

	projectIdGroup := projectGroup.Group("/:projectID")
	projectIdGroup.Use(middlewares.ValidateProject())

	projectIdGroup.DELETE("/", projects.DeleteProject)

	clientGroup := projectIdGroup.Group("/clients")
	clientGroup.POST("/", projects.CreateOidcClient)

	clientIdGroup := clientGroup.Group("/:clientID")
	clientIdGroup.Use(middlewares.ValidateClient())

	clientIdGroup.GET("/token", projects.GetToken)
}

func definePublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", public.Root)
	routerGroup.GET("/health", public.HealthCheck)
	routerGroup.GET("/.well-known/openid-configuration", public.WellKnown)
}

func defineAuthRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", auth.Signup)
	routerGroup.POST("/login", auth.Login)
	routerGroup.POST("/verify", auth.Verify)
}

func defineProtectedRoutes(routerGroup *gin.RouterGroup) {
	defineProjectRoutes(routerGroup)
}

func SetupRoutes(router *gin.Engine) {
	defaultRouterGroup := router.Group("/")

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")

	protectedRouterGroup := apiV1RouterGroup.Group("/")
	protectedRouterGroup.Use(middlewares.Auth())

	definePublicRoutes(defaultRouterGroup)
	defineAuthRoutes(apiV1RouterGroup)
	defineProtectedRoutes(protectedRouterGroup)
}
