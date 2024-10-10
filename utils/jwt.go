package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/usegranthq/backend/config"
)

type JwtUtils struct{}

var Jwt = JwtUtils{}

func (j *JwtUtils) SignToken(claims jwt.MapClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Get("JWT_SECRET")))
}

func (j *JwtUtils) DecodeToken(inputToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(inputToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get("JWT_SECRET")), nil
	})
}
