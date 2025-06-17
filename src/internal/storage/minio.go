package storage

import (
	"log"

	"github.com/lostish/woka-server/src/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitStorageClient() {
	cfg := config.GetEnv()
	var err error

	minioClient, err = minio.New(cfg.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MINIO_ACCESS_KEY, cfg.MINIO_SECRET_KEY, ""),
		Secure: cfg.MINIO_USE_SSL,
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func GetStorageClient() *minio.Client {
	return minioClient
}
