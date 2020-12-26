package v1

import (
	controllers "WebBackend/pkg/api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func SetCommentsRoutes(router *gin.RouterGroup) {
	commentsController := new(controllers.CommentsController)

	r := router.Group("comments")
	{
		r.GET("/:id", commentsController.Get)
		r.POST("/", commentsController.Post)
	}
}
