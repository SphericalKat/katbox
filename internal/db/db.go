package db

import (
	"context"
	"sync"

	"github.com/SphericalKat/katbox/ent"
	"github.com/SphericalKat/katbox/internal/config"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DB *ent.Client

// Connect establish a database connection
func Connect(ctx context.Context, wg *sync.WaitGroup) {
	// connect to db
	DB, err := ent.Open("postgres", config.Conf.DatabaseURL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// run automigration
	if err := DB.Schema.Create(ctx); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	log.Info("Connected to the database.")

	// graceful shutdown
	<-ctx.Done()
	log.Info("Shutting down database connection...")
	DB.Close()
	wg.Done()
}
