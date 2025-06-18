package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lostish/woka-server/src/internal/domain/auth"
	"github.com/lostish/woka-server/src/internal/domain/business"
	"github.com/lostish/woka-server/src/internal/domain/media"
	"github.com/lostish/woka-server/src/internal/domain/payments"
	"github.com/lostish/woka-server/src/internal/domain/user"
)

func InitRouting(r *gin.Engine) {
	// Root path generation for route generation tests
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Woka Server",
		})
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")
	auth.RegisterRoutes(v1.Group("/auth"))
	user.RegisterRoutes(v1.Group("/user"))
	business.RegisterRoutes(v1.Group("/businesss"))
	payments.RegisterRoutes(v1.Group("/payments"))
	media.RegisterRoutes(v1.Group("/media"))
}
