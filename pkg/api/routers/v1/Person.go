package v1

import (
	"github.com/gin-gonic/gin"
	controllers "WebBackend/pkg/api/controllers/v1"
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


