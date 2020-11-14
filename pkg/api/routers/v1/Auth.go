package v1

import (
	controllers "WebBackend/pkg/api/controllers/v1"
	middlewares "WebBackend/pkg/api/core/middlewares"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	authController := new(controllers.AuthController)

	r := router.Group("auth")
	{
		r.POST("/", authController.Login)
	}
}

func SetPrivateRoutes(router *gin.RouterGroup) {
	authorized := router.Group("aa")
	authorized.Use(middlewares.AuthRequired)
	{
		authorized.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
		})
	}
}
