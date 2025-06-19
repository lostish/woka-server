package payments

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handleListPaymentMethods)
	rg.POST("/", handleCreatePaymentMethod)
	rg.PUT("/:payment_id", handleUpdatePaymentMethod)
	rg.DELETE("/:payment_id", handleDeletePaymentMethod)
}

func handleListPaymentMethods(c *gin.Context) {}

func handleCreatePaymentMethod(c *gin.Context) {}

func handleUpdatePaymentMethod(c *gin.Context) {}

func handleDeletePaymentMethod(c *gin.Context) {}
