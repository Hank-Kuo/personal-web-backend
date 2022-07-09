package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/tag"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type TagHandler struct {
	tagUC  tag.Usecase
	logger logger.Logger
}

func NewTagHandler(e *gin.RouterGroup, tagUC tag.Usecase, logger logger.Logger) {
	handler := &TagHandler{
		tagUC:  tagUC,
		logger: logger,
	}

	e.GET("/tags", handler.Fetch)
	e.GET("/tag/:id", handler.GetByID)
}

func (h *TagHandler) Fetch(c *gin.Context) {
	ctx := c.Request.Context()
	paginationQuery, err := utils.GetPaginationFromGin(c)

	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "Bad params", Detail: err}, h.logger).ToJSON(c)
		return
	}

	data, err := h.tagUC.Fetch(ctx, paginationQuery)

	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	response.OK(http.StatusOK, "Get all tags", data).ToJSON(c)

}

func (h *TagHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}
	data, err := h.tagUC.GetByID(ctx, id)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	response.OK(http.StatusOK, "Get tag", data).ToJSON(c)

}
