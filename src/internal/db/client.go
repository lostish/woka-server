package db

import (
	"github.com/lostish/woka-server/src/internal/config"
	"github.com/surrealdb/surrealdb.go"
)

var db *surrealdb.DB

func InitClient() {
	cfg := config.GetEnv()
	var err error

	db, err = surrealdb.New(cfg.SURREAL_URL)
	if err != nil {
		panic(err)
	}

	if err = db.Use(cfg.SURREAl_NAMESPACE, cfg.SURREAL_DB); err != nil {
		panic(err)
	}

	authData := &surrealdb.Auth{
		Username: cfg.SURREAL_USER,     // use your setup username
		Password: cfg.SURREAL_PASSWORD, // use your setup password
	}

	token, err := db.SignIn(authData)
	if err != nil {
		panic(err)
	}

	if err := db.Authenticate(token); err != nil {
		panic(err)
	}
}

func GetClient() *surrealdb.DB {
	return db
}
