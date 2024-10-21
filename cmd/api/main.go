package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/db"
	"github.com/usegranthq/backend/external"
	"github.com/usegranthq/backend/utils"
)

func initDependencies() {
	config.Init()
	external.Init()
	utils.Init()
}

// setup server using gin
func main() {
	initDependencies()

	// setup db
	db.Connect()

	// enable prod mode
	if config.Get("NODE_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())

	api.SetupRoutes(router)

	port := config.Get("SERVER_PORT")
	router.Run(":" + port)
}
