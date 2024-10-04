package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/usegranthq/backend/api"
	"github.com/usegranthq/backend/db"
)

// setup server using gin
func main() {
	// setup db
	db.Connect()

	// enable prod mode
	if os.Getenv("NODE_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.SetTrustedProxies(nil)

	// handle panics
	router.Use(gin.Recovery())

	api.SetupRoutes(router)
	router.Run(":3000")
}
