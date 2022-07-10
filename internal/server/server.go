package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/middleware"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
	db     *sqlx.DB
	cache  *bigcache.BigCache
	logger logger.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, cache *bigcache.BigCache, logger logger.Logger) *Server {
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Server{engine: gin.Default(), cfg: cfg, db: db, logger: logger}
}

func (s *Server) Run() error {
	middleware.NewGlobalMiddlewares(s.engine)

	server := &http.Server{
		Addr:           ":" + s.cfg.Server.Port,
		Handler:        s.engine,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.registerHanders()
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Error http ListenAndServe: %s", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
