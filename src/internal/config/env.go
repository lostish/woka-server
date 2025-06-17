package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ConfEnv struct {
	SURREAL_URL       string `env:"SURREAL_URL,required"`
	SURREAL_USER      string `env:"SURREAL_USER,required"`
	SURREAL_PASSWORD  string `env:"SURREAL_PASSWORD,required"`
	SURREAl_NAMESPACE string `env:"SURREAl_NAMESPACE,required"`
	SURREAL_DB        string `env:"SURREAL_DB,required"`
	MINIO_ENDPOINT    string `env:"MINIO_ENDPOINT,required"`
	MINIO_ACCESS_KEY  string `env:"MINIO_ACCESS_KEY,required"`
	MINIO_SECRET_KEY  string `env:"MINIO_SECRET_KEY,required"`
	MINIO_USE_SSL     bool   `env:"MINIO_USE_SSL,required"`
	GIN_MODE          string `env:"GIN_MODE,required"`
}

// var cfg ConfEnv

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func GetEnv() ConfEnv {
	cfg, err := env.ParseAs[ConfEnv]()
	if err != nil {
		log.Fatal("Error parsing environment variables: ", err)
	}
	fmt.Println(cfg)
	return cfg
}
