package middlewares

import "net/http"

const (
	SUCCESSCODE                    = 1000 // 200
	SERVER_ERROR                   = 1001 // 500系统错误
	NOT_FOUND                      = 1002 // 404错误
	UNKNOWN_ERROR                  = 1003 // 未知错误
	PARAMETER_ERROR                = 1004 // 400参数错误
	FORBIDDEN_ERROR                = 1005 // 403
	ERROR_AUTH                     = 1006 // 401错误
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 1007
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 1008
	ERROR_AUTH_TOKEN               = 1009
)

var CodeFlags = map[int]int{
	SUCCESSCODE:                    http.StatusOK,
	SERVER_ERROR:                   http.StatusInternalServerError,
	NOT_FOUND:                      http.StatusNotFound,
	UNKNOWN_ERROR:                  http.StatusInternalServerError,
	PARAMETER_ERROR:                http.StatusBadRequest,
	FORBIDDEN_ERROR:                http.StatusForbidden,
	ERROR_AUTH:                     http.StatusUnauthorized,
	ERROR_AUTH_CHECK_TOKEN_FAIL:    http.StatusUnauthorized,
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: http.StatusUnauthorized,        // 401
	ERROR_AUTH_TOKEN:               http.StatusInternalServerError, // 500
}

var MsgFlags = map[int]string{
	SUCCESSCODE:                    "Successful",
	SERVER_ERROR:                   "Server Error",
	NOT_FOUND:                      "Not Found",
	UNKNOWN_ERROR:                  "Unknow Error",
	PARAMETER_ERROR:                "請求參數錯誤",
	FORBIDDEN_ERROR:                "Forbidden Error",
	ERROR_AUTH:                     "Unauthorized",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token判斷失敗",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超時",
	ERROR_AUTH_TOKEN:               "Token生成失敗",
}

func GetCode(code int) int {
	errorCode, ok := CodeFlags[code]
	if ok {
		return errorCode
	}
	return CodeFlags[UNKNOWN_ERROR]
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[UNKNOWN_ERROR]
}

/*
当GET, PUT和PATCH请求成功时，要返回对应的数据，及状态码200，即SUCCESS
当POST创建数据成功时，要返回创建的数据，及状态码201，即CREATED
当DELETE删除数据成功时，不返回数据，状态码要返回204，即NO CONTENT
当GET 不到数据时，状态码要返回404，即NOT FOUND
任何时候，如果请求有问题，如校验请求数据时发现错误，要返回状态码 400，即BAD REQUEST
当API 请求需要用户认证时，如果request中的认证信息不正确，要返回状态码 401，即NOT AUTHORIZED
当API 请求需要验证用户权限时，如果当前用户无相应权限，要返回状态码 403，即FORBIDDEN
*/
