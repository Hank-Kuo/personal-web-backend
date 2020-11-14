package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Status    bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"message"`
}

func HandleNotFound(c *gin.Context) {
	msg := c.Request.Method + " " + c.Request.URL.String()
	StatusCode := GetCode(NOT_FOUND)
	response := &Error{Status: false, ErrorCode: NOT_FOUND, Msg: msg}
	c.JSON(StatusCode, response)
}

func ResponseError(c *gin.Context, code int, msg string, err error) {
	StatusCode := GetCode(code)
	if len(msg) == 0 {
		// default message
		msg = GetMsg(code)
	}
	resp := &Error{Status: false, ErrorCode: code, Msg: msg}
	c.JSON(StatusCode, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(StatusCode, err)
}

type Success struct {
	Status bool        `json:"success"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, msg string, data interface{}) {
	StatusCode := GetCode(SUCCESSCODE)
	resp := &Success{Status: true, Msg: msg, Data: data}
	c.JSON(StatusCode, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
