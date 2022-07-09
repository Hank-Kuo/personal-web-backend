package server

import (
	"github.com/Hank-Kuo/personal-web-backend/internal/middleware"

	blogDelivery "github.com/Hank-Kuo/personal-web-backend/internal/api/blog/delivery/http"
	blogRepository "github.com/Hank-Kuo/personal-web-backend/internal/api/blog/repository"
	blogUsecase "github.com/Hank-Kuo/personal-web-backend/internal/api/blog/usecase"

	tagDelivery "github.com/Hank-Kuo/personal-web-backend/internal/api/tag/delivery/http"
	tagRepository "github.com/Hank-Kuo/personal-web-backend/internal/api/tag/repository"
	tagUsecase "github.com/Hank-Kuo/personal-web-backend/internal/api/tag/usecase"

	emojiDelivery "github.com/Hank-Kuo/personal-web-backend/internal/api/emoji/delivery/http"
	emojiRepository "github.com/Hank-Kuo/personal-web-backend/internal/api/emoji/repository"
	emojiUsecase "github.com/Hank-Kuo/personal-web-backend/internal/api/emoji/usecase"

	commentDelivery "github.com/Hank-Kuo/personal-web-backend/internal/api/comment/delivery/http"
	commentRepository "github.com/Hank-Kuo/personal-web-backend/internal/api/comment/repository"
	commentUsecase "github.com/Hank-Kuo/personal-web-backend/internal/api/comment/usecase"

	userDelivery "github.com/Hank-Kuo/personal-web-backend/internal/api/user/delivery/http"
	userRepository "github.com/Hank-Kuo/personal-web-backend/internal/api/user/repository"
	userUsecase "github.com/Hank-Kuo/personal-web-backend/internal/api/user/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerHanders() {
	api := s.engine.Group("/api")

	blogRepo := blogRepository.NewReposity(s.db)
	tagRepo := tagRepository.NewReposity(s.db)
	emojiRepo := emojiRepository.NewReposity(s.db)
	commentRepo := commentRepository.NewReposity(s.db)
	userRepo := userRepository.NewReposity(s.db)

	blogUC := blogUsecase.NewUsecase(s.cfg, blogRepo, tagRepo, emojiRepo, s.logger)
	emojiUC := emojiUsecase.NewUsecase(s.cfg, emojiRepo, s.logger)
	tagUC := tagUsecase.NewUsecase(s.cfg, tagRepo, s.logger)
	commentUC := commentUsecase.NewUsecase(s.cfg, commentRepo, blogRepo, s.logger)
	userUC := userUsecase.NewUsecase(s.cfg, userRepo, s.logger)

	mid := middleware.NewMiddlewares(s.cfg, userUC, s.logger)

	blogDelivery.NewBlogHandler(api, blogUC, mid, s.logger)
	tagDelivery.NewTagHandler(api, tagUC, s.logger)
	emojiDelivery.NewEmojiHandler(api, emojiUC, s.logger)
	commentDelivery.NewCommentHandler(api, commentUC, s.logger)
	userDelivery.NewUserHandler(api, userUC, s.logger)

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome Hank kuo Server")
	})

}
