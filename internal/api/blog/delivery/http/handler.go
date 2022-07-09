package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/blog"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/middleware"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type BlogHandler struct {
	blogUC blog.Usecase
	logger logger.Logger
}

func NewBlogHandler(e *gin.RouterGroup, blogUC blog.Usecase, mid *middleware.Middleware, logger logger.Logger) {
	handler := &BlogHandler{
		blogUC: blogUC,
		logger: logger,
	}

	e.GET("/blogs", handler.Fetch)
	e.GET("/blog/:id", handler.GetByID)
	e.POST("/blog", mid.JWTAuthMiddleware(), handler.Create)
	e.PUT("/blog/:id", mid.JWTAuthMiddleware(), handler.Update)
	e.GET("/visit", mid.RateLimit(), handler.Visit)

}

func (h *BlogHandler) Fetch(c *gin.Context) {
	ctx := c.Request.Context()
	paginationQuery, err := utils.GetPaginationFromGin(c)

	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "Bad params", Detail: err}, h.logger).ToJSON(c)
		return
	}

	data, err := h.blogUC.Fetch(ctx, paginationQuery)

	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	response.OK(http.StatusOK, "Get all blogs successfully", data).ToJSON(c)

}

func (h *BlogHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}
	data, err := h.blogUC.GetByID(ctx, id)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	html, err := utils.Crawler(data.Link)
	if err != nil {
		response.Fail(utils.HttpError{Message: "can't crawl html", Detail: err}, h.logger).ToJSON(c)
	}

	data.Html = html

	response.OK(http.StatusOK, "Get all blogs successfully", data).ToJSON(c)

}

func (h *BlogHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var body dto.CreateBlogDto
	if err := c.ShouldBindJSON(&body); err != nil {
		errMsg := utils.GetValidatMsg(err)
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: errMsg, Detail: err}, h.logger).ToJSON(c)
		return
	}

	if err := h.blogUC.Create(ctx, &body); err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	response.OK(http.StatusOK, "create blog successfully", nil).ToJSON(c)

}

func (h *BlogHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}

	var body dto.UpdateBlogDto
	if err := c.ShouldBindJSON(&body); err != nil {
		errMsg := utils.GetValidatMsg(err)
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: errMsg, Detail: err}, h.logger).ToJSON(c)
		return
	}

	if err := h.blogUC.Update(ctx, id, &body); err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	response.OK(http.StatusOK, "update blog successfully", nil).ToJSON(c)
}

func (h *BlogHandler) Visit(c *gin.Context) {
	ctx := c.Request.Context()

	blogId := c.Query("blogId")
	if len(blogId) == 0 {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "not found blogId", Detail: errors.New("not found blogId")}, h.logger).ToJSON(c)
		return
	}
	id, err := strconv.Atoi(blogId)
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}
	data, err := h.blogUC.GetByID(ctx, id)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	if err := h.blogUC.Update(ctx, id, &dto.UpdateBlogDto{Visitor: data.Visitor + 1}); err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	response.OK(http.StatusOK, "visit blog successfully", nil).ToJSON(c)
}
