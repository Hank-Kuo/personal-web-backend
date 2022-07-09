package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/emoji"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type EmojiHandler struct {
	emojiUC emoji.Usecase
	logger  logger.Logger
}

func NewEmojiHandler(e *gin.RouterGroup, emojiUC emoji.Usecase, logger logger.Logger) {
	handler := &EmojiHandler{
		emojiUC: emojiUC,
		logger:  logger,
	}

	e.PUT("/emoji/:id", handler.Update)
}

func (h *EmojiHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "can't convert string to int", Detail: err}, h.logger).ToJSON(c)
		return
	}
	var body dto.EmojiDto
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "Bad params", Detail: err}, h.logger).ToJSON(c)
		return
	}
	data, err := h.emojiUC.Update(ctx, id, &body)
	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}

	response.OK(http.StatusOK, "Update emoji successfully", data).ToJSON(c)

}
