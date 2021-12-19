package v1

import (
	controllers "github.com/Hank-Kuo/personal-web-backend/pkg/api/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetPersonRoutes(router *gin.RouterGroup) {
	peopleController := new(controllers.PeopleController)

	r := router.Group("people")
	{
		r.GET("/", peopleController.GetAll)
		r.GET("/:id", peopleController.Get)
		r.POST("/", peopleController.Post)
		r.PUT("/:id", peopleController.Put)
		r.DELETE("/:id", peopleController.Delete)
	}
}
