package utils

import (
	"strings"
	"time"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

func GetJwt(cfg *config.Config, user *model.User, tokenType string) (string, error) {
	now := time.Now()
	jwtSecret := ""
	if tokenType == "access" {
		jwtSecret = cfg.Server.AccessJwt
	} else {
		jwtSecret = cfg.Server.RefreshJwt
	}

	claims := &Claims{
		Account: user.Account,
		Email:   user.Email,
		Role:    user.Role,
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Account,
			ExpiresAt: now.Add(cfg.Server.JwtExpireTime * time.Second).Unix(),
			Id:        user.UUID,
			IssuedAt:  now.Unix(),
			Issuer:    "hank-kuo",
			NotBefore: now.Unix() - 1000,
			Subject:   user.Account,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

func ParseJwt(cfg *config.Config, token string, tokenType string) (*jwt.Token, error) {
	t := strings.Split(token, "Bearer ")[1]

	jwtSecret := ""
	if tokenType == "access" {
		jwtSecret = cfg.Server.AccessJwt
	} else {
		jwtSecret = cfg.Server.RefreshJwt
	}
	tokenClaims, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(jwtSecret), nil
	})
	return tokenClaims, err
}
