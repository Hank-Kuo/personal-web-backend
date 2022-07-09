package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/comment"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type CommentHandler struct {
	commentUC comment.Usecase
	logger    logger.Logger
}

func NewCommentHandler(e *gin.RouterGroup, commentUC comment.Usecase, logger logger.Logger) {
	handler := &CommentHandler{
		commentUC: commentUC,
		logger:    logger,
	}

	e.GET("/comments", handler.Fetch)
	e.POST("/comment", handler.Create)
}

type fetchReq struct {
	BlogID int `form:"blogId"`
}

func (h *CommentHandler) Fetch(c *gin.Context) {
	ctx := c.Request.Context()
	paginationQuery, err := utils.GetPaginationFromGin(c)
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "Bad params", Detail: err}, h.logger).ToJSON(c)
		return
	}

	var query fetchReq

	if err := c.BindQuery(&query); err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "blog id query string error", Detail: err}, h.logger).ToJSON(c)
		return
	}

	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}

	data, err := h.commentUC.GetByBlogID(ctx, query.BlogID, paginationQuery)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	response.OK(http.StatusOK, "GET comments successfully", data).ToJSON(c)
}

func (h *CommentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var body dto.CraeteCommentDto
	if err := c.ShouldBindJSON(&body); err != nil {
		errMsg := utils.GetValidatMsg(err)
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: errMsg, Detail: err}, h.logger).ToJSON(c)
		return
	}

	data, err := h.commentUC.Create(ctx, &body)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	response.OK(http.StatusOK, "Create comment successfully", data).ToJSON(c)
}
