package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractToken(authHeader string) string {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}

	token := authHeader[len("Bearer "):]
	return token
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

func GrabToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	return authHeader
}

func ValidateToken(c *gin.Context) (*jwt.Token, error) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	token, err := ParseJWTToken(ExtractToken(authHeader))
	if err != nil {
		log.Println("TOKEN SUKA: ", token)
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

func GetIdByTokenClaim(c *gin.Context) int {
	token, err := ValidateToken(c)
	if err != nil {
		return -1
	}

	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		userId := claims["http://schemas.microsoft.com/ws/2008/06/identity/claims/userdata"].(string)

		n, err := strconv.Atoi(userId)
		if err != nil {
			log.Println("Error: ", err)
			return -1
		}

		return n

	}

	return -1

}
