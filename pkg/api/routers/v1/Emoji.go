package v1

import (
	controllers "github.com/Hank-Kuo/personal-web-backend/pkg/api/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetEmojiRoutes(router *gin.RouterGroup) {
	emojiController := new(controllers.EmojiController)

	r := router.Group("emoji")
	{
		r.GET("/:id", emojiController.Get)
		r.PUT("/:id", emojiController.Put)
	}
}
