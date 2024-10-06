package db

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/usegranthq/backend/config"
	"github.com/usegranthq/backend/ent"
)

var Client *ent.Client

func Connect() {
	var err error

	// Get database connection details from environment variables
	dbHost := config.Get("DB_HOST")
	dbPort := config.Get("DB_PORT")
	dbUser := config.Get("DB_USER")
	dbName := config.Get("DB_NAME")
	dbPassword := config.Get("DB_PASSWORD")

	// Construct the database connection string
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	// Connect to the database
	Client, err = ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run auto migration
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("Successfully connected to the database and ran migrations")
}
