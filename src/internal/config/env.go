package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfEnv struct {
	SURREAL_URL         string
	SURREAL_USER        string
	SURREAL_PASSWORD    string
	SURREAL_NAMESPACE   string
	SURREAL_DB          string
	MINIO_ENDPOINT      string
	MINIO_ROOT_USER     string
	MINIO_ROOT_PASSWORD string
	MINIO_USE_SSL       bool
	ENV                 string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar .env localmente, pero no es crítico.")
	}
}

var env ConfEnv

func GetEnv() ConfEnv {
	env = ConfEnv{
		SURREAL_URL:         os.Getenv("SURREAL_URL"),
		SURREAL_USER:        os.Getenv("SURREAL_USER"),
		SURREAL_PASSWORD:    os.Getenv("SURREAL_PASSWORD"),
		SURREAL_NAMESPACE:   os.Getenv("SURREAL_NAMESPACE"),
		SURREAL_DB:          os.Getenv("SURREAL_DB"),
		MINIO_ENDPOINT:      os.Getenv("MINIO_ENDPOINT"),
		MINIO_ROOT_USER:     os.Getenv("MINIO_ROOT_USER"),
		MINIO_ROOT_PASSWORD: os.Getenv("MINIO_ROOT_PASSWORD"),
		MINIO_USE_SSL:       os.Getenv("MINIO_USE_SSL") == "true",
		ENV:                 os.Getenv("ENV"),
	}
	return env
}
