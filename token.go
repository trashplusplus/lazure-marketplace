package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractToken(authHeader string) (string, bool) {

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", false
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		return "", false
	}

	return token, true
}

func ParseJWTToken(tokenString string) (*jwt.Token, error) {
	signatureKeyEnv := os.Getenv("SIGNATUREKEY")
	signatureKey := []byte(signatureKeyEnv)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return signatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidateToken(c *gin.Context) (*jwt.Token, error) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	token, err := ParseJWTToken(authHeader)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return token, nil
}
