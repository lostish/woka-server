package storage

import (
	"context"
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
		Creds:  credentials.NewStaticV4(cfg.MINIO_ROOT_USER, cfg.MINIO_ROOT_PASSWORD, ""),
		Secure: cfg.MINIO_USE_SSL,
	})

	if err != nil {
		log.Fatalf("failed to init MinIO client: %v", err)
	}

	ctx := context.Background()
	_, err = minioClient.ListBuckets(ctx)
	if err != nil {
		log.Fatalf("MinIO ping failed: %v", err)
	}
	log.Println("MinIO connected OK")
}

func GetStorageClient() *minio.Client {
	return minioClient
}
