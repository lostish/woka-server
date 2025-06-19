package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lostish/woka-server/src/internal/config"
	"github.com/lostish/woka-server/src/internal/db"
	"github.com/lostish/woka-server/src/internal/domain"
	"github.com/lostish/woka-server/src/internal/storage"
)

func init() {
	config.LoadEnv()
	db.InitClient()
	storage.InitStorageClient()
}

func main() {
	router := gin.Default()

	domain.InitRouting(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
