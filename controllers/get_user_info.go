package controllers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"okAuth/types"
)

func GetUserInfo(conn *pgx.Conn, login string) Response {

	query := "SELECT * FROM users WHERE login = $1"
	result := Response{}
	conn.QueryRow(context.Background(), query, login).Scan(&result.login, &result.password, &result.role)

	return result
}

type Response struct {
	login    string
	password string
	role     RoleType.Role
}
