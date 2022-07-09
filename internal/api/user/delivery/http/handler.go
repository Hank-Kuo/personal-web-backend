package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/user"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type TagHandler struct {
	userUC user.Usecase
	logger logger.Logger
}

func NewUserHandler(e *gin.RouterGroup, userUC user.Usecase, logger logger.Logger) {
	handler := &TagHandler{
		userUC: userUC,
		logger: logger,
	}

	e.POST("/login", handler.Login)
}

// Login godoc
// @Summary Login
// @Tags Auth
// @Description Login
// @version 1.0
// @Param polygon body dto.LoginReqDto true "body"
// @produce application/json
// @Success 200 {object} response.OK(http.StatusOK, "login successfully", data) "success"
// @Router /api/login [post]
func (h *TagHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var body dto.LoginReqDto
	if err := c.ShouldBindJSON(&body); err != nil {
		errMsg := utils.GetValidatMsg(err)
		response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: errMsg, Detail: err}, h.logger).ToJSON(c)
		return
	}

	data, err := h.userUC.Login(ctx, &body)

	if err != nil {
		response.Fail(err, h.logger).ToJSON(c)
		return
	}
	response.OK(http.StatusOK, "login successfully", data).ToJSON(c)

}
