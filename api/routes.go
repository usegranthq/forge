package api

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api/middlewares"
	"github.com/usegranthq/backend/api/services/auth"
	"github.com/usegranthq/backend/api/services/projects"
	"github.com/usegranthq/backend/api/services/public"
	"github.com/usegranthq/backend/api/services/users"
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

func defineUserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/me", users.GetUser)
	routerGroup.GET("/refresh", users.Refresh)
	routerGroup.POST("/logout", users.Logout)
}

func definePublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", public.Root)
	routerGroup.GET("/health", public.HealthCheck)
	routerGroup.GET("/.well-known/openid-configuration", public.WellKnown)
}

func defineAuthRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", auth.Signup)
	routerGroup.POST("/login", auth.Login)
	routerGroup.GET("/login/github", auth.GithubLogin)
	routerGroup.GET("/login/google", auth.GoogleLogin)
	routerGroup.POST("/verify", auth.Verify)
	routerGroup.POST("/verify/github", auth.VerifyGithub)
	routerGroup.POST("/verify/google", auth.VerifyGoogle)
}

func defineProtectedRoutes(routerGroup *gin.RouterGroup) {
	defineProjectRoutes(routerGroup)
}

func SetupRoutes(router *gin.Engine) {
	defaultRouterGroup := router.Group("/")

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")

	protectedRouterGroup := apiV1RouterGroup.Group("/u")
	protectedRouterGroup.Use(middlewares.Auth())

	definePublicRoutes(defaultRouterGroup)
	defineAuthRoutes(apiV1RouterGroup)
	defineUserRoutes(protectedRouterGroup)
	defineProtectedRoutes(protectedRouterGroup)
}
