package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
	SetAuthRoutes(g)
	SetPrivateRoutes(g)
	SetBlogRoutes(g)
	SetEmojiRoutes(g)
	SetCommentsRoutes(g)
}
