package db

import (
	"context"
	"fmt"
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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
