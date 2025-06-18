package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", handleLogin)
	rg.POST("/register", handleRegister)
	rg.POST("/reset-pwd", handleResetPwd)
	rg.POST("/verify", handleVerify)
	rg.POST("/logout", handleLogout)
}

func handleLogin(c *gin.Context) {
	// Login
}

func handleRegister(c *gin.Context) {
	// similar a Login, validando y llamando al RegisterService
}

func handleResetPwd(c *gin.Context) {
	// reset pwd here
}

func handleLogout(c *gin.Context) {
	// logout user here
}

func handleVerify(c *gin.Context) {
	// verify email, access and others..
}
