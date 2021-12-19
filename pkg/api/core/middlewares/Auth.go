package middlewares

import (
	configs "github.com/Hank-Kuo/personal-web-backend/pkg/api/configs"
	dto "github.com/Hank-Kuo/personal-web-backend/pkg/api/core/dto"

	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if len(auth) == 0 {
		// not found token
		err := errors.New("Unauthorized")
		ResponseError(c, 1006, "Not found any Token", err)
	}
	token := strings.Split(auth, "Bearer ")[1]
	jwtSecret := configs.GetSecretKey()

	tokenClaims, err := jwt.ParseWithClaims(token, &dto.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	// fmt.Println(tokenClaims, err)

	if err != nil {
		// invalid token
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
		ResponseError(c, 1007, message, err)
	}

	if claims, ok := tokenClaims.Claims.(*dto.Claims); ok && tokenClaims.Valid {
		c.Set("account", claims.Account)
		c.Set("role", claims.Role)
		c.Next()
	}
}
