package dto

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}
