package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usegranthq/forge/api"
	"github.com/usegranthq/forge/config"
	"github.com/usegranthq/forge/db"
	"github.com/usegranthq/forge/external"
	"github.com/usegranthq/forge/utils"
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
	utils.Log.Infof("Starting server on port %s", port)

	router.Run(":" + port)
}
