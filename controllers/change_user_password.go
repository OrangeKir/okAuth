package controllers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"okAuth/models"
)

func ChangeUserPassword(conn *pgx.Conn, request models.ChangeUserPasswordRequest) error {
	userInfo := GetUserInfo(conn, request.Login)

	if userInfo.password != HashPassword(request.OldPassword) {
		return fmt.Errorf("auth error")
	}

	query := "Update users SET password = $1 WHERE login = $2"
	_, err := conn.Exec(context.Background(), query, HashPassword(request.NewPassword), request.Login)

	return err
}
