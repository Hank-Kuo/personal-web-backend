package main

import (
	"fmt"
	"log"

	"github.com/Hank-Kuo/personal-web-backend/config"
	// middlewares "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/middlewares"

	// _ "github.com/Hank-Kuo/personal-web-backend/pkg/docs"
	// "github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/Hank-Kuo/personal-web-backend/internal/server"
	"github.com/Hank-Kuo/personal-web-backend/pkg/database"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
)

func main() {
	log.Println("Starting api server")

	// init config
	cfg, err := config.GetConf()
	if err != nil {
		panic(fmt.Errorf("load config: %v", err))
	}

	// init logger
	apiLogger := logger.NewApiLogger(cfg)
	apiLogger.InitLogger()

	// init database
	db, err := database.ConnectDB(&cfg.Database)
	if err != nil {
		panic(fmt.Errorf("load database: %v", err))
	}
	defer db.Close()

	// init in-memory cache
	cache, err := database.ConnectCacheDB(&cfg.Cache)
	if err != nil {
		panic(fmt.Errorf("load cache: %v", err))
	}

	// init server
	srv := server.NewServer(cfg, db, cache, apiLogger)
	if err = srv.Run(); err != nil {
		apiLogger.Fatal(err)
	}
}
