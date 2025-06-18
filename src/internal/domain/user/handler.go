package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handleListUsers)
	rg.GET("/:id", handleGetUser)

	rg.POST("/", handleCreateUser)

	rg.PUT("/:id", handleUpdateUser)

	rg.DELETE("/:id", handleDeleteUser)
}

func handleGetUser(c *gin.Context) {}

func handleListUsers(c *gin.Context) {}

func handleCreateUser(c *gin.Context) {}

func handleUpdateUser(c *gin.Context) {}

func handleDeleteUser(c *gin.Context) {}
