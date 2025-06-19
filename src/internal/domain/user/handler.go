package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handleListUsers)
	rg.GET("/:id", handleGetUser)

	rg.POST("/", handleCreateUser)
	rg.POST("/:id/friend-request", handleFriendRequest)
	rg.POST("/:id/friend-reply", handleFriendReply)

	rg.PUT("/:id", handleUpdateUser)
	rg.PUT("/:id/friend-update", handleFriendUpdate)

	rg.DELETE("/:id", handleDeleteUser)
}

func handleListUsers(c *gin.Context) {}

func handleGetUser(c *gin.Context) {}

func handleCreateUser(c *gin.Context) {}

func handleFriendRequest(c *gin.Context) {}

func handleFriendReply(c *gin.Context) {}

func handleUpdateUser(c *gin.Context) {}

func handleFriendUpdate(c *gin.Context) {}

func handleDeleteUser(c *gin.Context) {}
