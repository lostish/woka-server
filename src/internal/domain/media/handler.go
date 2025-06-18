package media

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/search", handleSearch)
	rg.POST("/upload", handleUpload)
	rg.PUT("/update", handleUpdate)
	rg.DELETE("/remove", handleRemove)
}

func handleSearch(c *gin.Context) {}
func handleUpload(c *gin.Context) {}
func handleUpdate(c *gin.Context) {}
func handleRemove(c *gin.Context) {}
