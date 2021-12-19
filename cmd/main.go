package main

import (
	// configs "https://github.com/Hank-Kuo/personal-web-backend/pkg/api/configs"
	configs "github.com/Hank-Kuo/personal-web-backend/pkg/api/configs"
	middlewares "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/middlewares"
	routers "github.com/Hank-Kuo/personal-web-backend/pkg/api/routers/v1"
	_ "github.com/Hank-Kuo/personal-web-backend/pkg/docs"

	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin API Swagger
// @version 1.0
// @description This is a backend server.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// init DB
	db := configs.ConnectDB()
	defer configs.CloseDB()

	// init Server
	port := 8080
	fmt.Println("Server Running on Port: ", port)
	engine := gin.New()

	// middleware
	engine.Use(gin.Logger(), gin.Recovery(), middlewares.CORSMiddleware())
	engine.NoMethod(middlewares.HandleNotFound)
	engine.NoRoute(middlewares.HandleNotFound)
	engine.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// init Routes
	v1 := engine.Group("/api/v1")
	routers.InitRoutes(v1)

	// import swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// start server
	engine.Run(":3000")
}
