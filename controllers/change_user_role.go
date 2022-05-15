package controllers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"okAuth/models"
)

func ChangeUserRole(conn *pgx.Conn, request models.ChangeUserRoleRequest) error {
	userInfo := GetUserInfo(conn, request.Login)

	if userInfo.password != HashPassword(request.Password) {
		return fmt.Errorf("auth error")
	}

	query := "Update users SET role = $1 WHERE login = $2"
	_, err := conn.Exec(context.Background(), query, request.Role, request.Login)

	return err
}
