package api

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api/middlewares"
	"github.com/usegranthq/backend/api/services/auth"
	"github.com/usegranthq/backend/api/services/projects"
	"github.com/usegranthq/backend/api/services/public"
	"github.com/usegranthq/backend/api/services/settings"
	"github.com/usegranthq/backend/api/services/users"
)

func defineProjectRoutes(routerGroup *gin.RouterGroup) {
	projectGroup := routerGroup.Group("/projects")
	projectGroup.POST("", projects.CreateProject)
	projectGroup.GET("", projects.ListProjects)

	projectIdGroup := projectGroup.Group("/:projectID")
	projectIdGroup.Use(middlewares.ValidateProject())

	projectIdGroup.DELETE("", projects.DeleteProject)
	projectIdGroup.PUT("", projects.UpdateProject)

	clientGroup := projectIdGroup.Group("/clients")
	clientGroup.POST("", projects.CreateOidcClient)
	clientGroup.GET("", projects.ListOidcClients)

	clientIdGroup := clientGroup.Group("/:clientID")
	clientIdGroup.Use(middlewares.ValidateClient())

	clientIdGroup.DELETE("", projects.DeleteOidcClient)
	clientIdGroup.GET("/token", projects.GetToken)
}

func defineUserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/me", users.GetUser)
	routerGroup.GET("/refresh", users.Refresh)
	routerGroup.POST("/logout", users.Logout)
	routerGroup.DELETE("/", users.DeleteUser)
}

func defineSettingsRoutes(routerGroup *gin.RouterGroup) {
	settingsRouterGroup := routerGroup.Group("/settings")

	settingsRouterGroup.GET("/tokens", settings.ListTokens)
	settingsRouterGroup.POST("/tokens", settings.CreateToken)
	settingsRouterGroup.DELETE("/tokens/:tokenID", settings.DeleteToken)

	settingsRouterGroup.GET("/sessions", settings.ListSessions)
	settingsRouterGroup.DELETE("/sessions/:sessionID", settings.DeleteSession)
	settingsRouterGroup.DELETE("/sessions", settings.DeleteAllOtherSessions)
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
	defineUserRoutes(routerGroup)
	defineProjectRoutes(routerGroup)
	defineSettingsRoutes(routerGroup)
}

func SetupRoutes(router *gin.Engine) {
	defaultRouterGroup := router.Group("/")

	apiRouterGroup := router.Group("/api")
	apiV1RouterGroup := apiRouterGroup.Group("/v1")

	protectedRouterGroup := apiV1RouterGroup.Group("/u")
	protectedRouterGroup.Use(middlewares.Auth())

	definePublicRoutes(defaultRouterGroup)
	defineAuthRoutes(apiV1RouterGroup)
	defineProtectedRoutes(protectedRouterGroup)
}
