package middlewares

import (
	configs "WebBackend/pkg/api/configs"
	dto "WebBackend/pkg/api/core/dto"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func Jwt(account string, role string) (string, error) {
	now := time.Now()
	jwtID := account + strconv.FormatInt(now.Unix(), 10)
	jwtSecret := configs.GetSecretKey()

	claims := &dto.Claims{
		Account: account,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			Audience:  account,
			ExpiresAt: now.Add(8 * time.Hour).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(9 * time.Hour).Unix(),
			Subject:   account,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
