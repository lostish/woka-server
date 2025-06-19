package db

import (
	"log"

	"github.com/lostish/woka-server/src/internal/config"
	"github.com/surrealdb/surrealdb.go"
)

var db *surrealdb.DB

func InitClient() {
	cfg := config.GetEnv()
	var err error

	log.Println("Connecting to SurrealDB...")

	// Create connection
	db, err = surrealdb.New(cfg.SURREAL_URL)
	if err != nil {
		panic(err)
	}

	// Set namespace and database
	if err = db.Use(cfg.SURREAL_NAMESPACE, cfg.SURREAL_DB); err != nil {
		panic(err)
	}

	// Try root authentication
	_, err = db.SignIn(&surrealdb.Auth{
		Username: cfg.SURREAL_USER,
		Password: cfg.SURREAL_PASSWORD,
	})
	if err != nil {
		log.Printf("Root SignIn failed: %v", err)
		panic(err)
	}

	log.Println("Successfully connected to SurrealDB")
}

func GetClient() *surrealdb.DB {
	return db
}
