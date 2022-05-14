package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	"okAuth/types"
	"time"
)

func GetUserToken(conn *pgx.Conn, login string, password string) (string, error) {
	userInfo := GetUserInfo(conn, login)

	hashedPassword := HashPassword(password)

	if hashedPassword != userInfo.password {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Login":     userInfo.login,
		"ExpiresAt": time.Now().Add(time.Minute * 1).Unix(),
		"Role":      userInfo.role,
	})

	return token.SigningString()
}

type CustomerInfo struct {
	login string
	role  RoleType.Role
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	CustomerInfo
}
