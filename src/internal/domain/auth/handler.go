package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", Login)
	rg.POST("/register", Register)
}

func Login(c *gin.Context) {
	// Login
}

func Register(c *gin.Context) {
	// similar a Login, validando y llamando al RegisterService
}
