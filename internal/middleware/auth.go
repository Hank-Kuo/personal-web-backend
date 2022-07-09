package middleware

import (
	"errors"
	"net/http"

	"github.com/Hank-Kuo/personal-web-backend/pkg/response"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (m *Middleware) JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if len(auth) == 0 {
			response.Fail(utils.HttpError{Status: http.StatusBadRequest, Message: "token not found", Detail: errors.New("toke is empty")}, m.logger).ToJSON(c)
			return
		}
		tokenClaims, err := utils.ParseJwt(m.cfg, auth, "access")

		if err != nil {
			var message string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					message = "token is malformed"
				} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
					message = "token could not be verified because of signing problems"
				} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					message = "signature validation failed"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					message = "token is expired"
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					message = "token is not yet valid before sometime"
				} else {
					message = "can not handle this token"
				}
			}
			response.Fail(utils.HttpError{Status: http.StatusUnauthorized, Message: message, Detail: err}, m.logger).ToJSON(c)
		}

		if claims, ok := tokenClaims.Claims.(*utils.Claims); ok && tokenClaims.Valid {
			_, err := m.userUC.GetUser(c, claims.Account)
			if err != nil {
				response.Fail(err, m.logger).ToJSON(c)
				return
			}
			c.Next()
		}
	}
}
