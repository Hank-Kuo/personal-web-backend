package server

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) registerSwagger() {
	if s.cfg.Server.Debug {
		url := ginSwagger.URL("http://localhost:" + s.cfg.Server.Port + "/swagger/doc.json")
		s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
