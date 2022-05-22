package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"okAuth/models"
	"time"
)

func ValidateToken(tokenString string, tokenSecret string) (bool, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(tokenSecret), nil
	})

	claims, _ := token.Claims.(*models.TokenClaims)
	isValid := token.Valid && claims.ExpiresAt > time.Now().Unix()

	return isValid, err
}
