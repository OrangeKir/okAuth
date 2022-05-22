package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	"okAuth/models"
	"okAuth/types"
	"time"
)

func GetUserToken(conn *pgx.Conn, login string, password string, tokenSecret string) (string, error) {
	userInfo := GetUserInfo(conn, login)

	hashedPassword := HashPassword(password)

	if hashedPassword != userInfo.password {
		return "", fmt.Errorf("auth error")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: userInfo.login,
		Role:  userInfo.role,
	})

	return token.SignedString([]byte(tokenSecret))
}

type CustomerInfo struct {
	login string
	role  RoleType.Role
}
