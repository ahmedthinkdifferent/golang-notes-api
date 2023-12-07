package util

import (
	"github.com/golang-jwt/jwt/v5"
	"noteapp/core/env"
	"time"
)

func GenerateJwtToken(claims map[string]any) string {
	jwtClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}
	for key, value := range claims {
		jwtClaims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, _ := token.SignedString([]byte(env.AppEnvironment.App.JwtToken))
	return tokenString
}

func VerifyJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.AppEnvironment.App.JwtToken), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
