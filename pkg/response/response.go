package response

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type response struct {
	Body       *responseBody
	StatusCode int
}

type responseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(statusCode int, message string, data interface{}) *response {
	return &response{&responseBody{Status: "success", Message: message, Data: data}, statusCode}
}

func Fail(err error, logger logger.Logger) *response {
	parseErr := parseError(err)
	logger.Error(parseErr.Detail)
	return &response{&responseBody{Status: "fail", Message: parseErr.GetMessage()}, parseErr.GetStatus()}
}

func (r *response) ToJSON(c *gin.Context) {
	if r.Body.Status == "fail" {
		c.AbortWithStatusJSON(r.StatusCode, r.Body)
	} else {
		c.JSON(r.StatusCode, r.Body)
	}
}

func parseError(err error) *utils.HttpError {
	var parseErr utils.HttpError

	switch {
	case errors.As(err, &parseErr):
		return &parseErr
	default:
		parseErr.Detail = err
		return &parseErr
	}
}
