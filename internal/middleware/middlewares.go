package middleware

import (
	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/user"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userUC user.Usecase
	cfg    *config.Config
	logger logger.Logger
}

func NewMiddlewares(cfg *config.Config, userUC user.Usecase, logger logger.Logger) *Middleware {
	return &Middleware{cfg: cfg, userUC: userUC, logger: logger}
}

func NewGlobalMiddlewares(engine *gin.Engine) {
	engine.Use(gin.Logger(), gin.Recovery(), CORSMiddleware())
	engine.NoMethod(httpNotFound)
	engine.NoRoute(httpNotFound)
}
