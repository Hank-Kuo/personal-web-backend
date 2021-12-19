package v1

import (
	controllers "github.com/Hank-Kuo/personal-web-backend/pkg/api/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetBlogRoutes(router *gin.RouterGroup) {
	blogController := new(controllers.BlogController)

	r := router.Group("blog")
	{
		r.GET("/", blogController.GetAll)
		r.GET("/:id", blogController.Get)
		r.POST("/", blogController.Post)
		r.PUT("/:id", blogController.Put)
	}
}
