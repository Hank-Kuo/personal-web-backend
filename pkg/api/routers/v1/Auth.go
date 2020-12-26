package v1

import (
	controllers "WebBackend/pkg/api/controllers/v1"
	middlewares "WebBackend/pkg/api/core/middlewares"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	authController := new(controllers.AuthController)

	r := router.Group("")
	{
		r.POST("/login", authController.Login)
		r.POST("/register", authController.Register)
	}
}

func SetPrivateRoutes(router *gin.RouterGroup) {
	blogController := new(controllers.BlogController)

	authorized := router.Group("auth")
	authorized.Use(middlewares.AuthRequired)
	{
		authorized.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
		})
		authorized.POST("/blog", blogController.Post)
	}
}
