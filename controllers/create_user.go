package controllers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"okAuth/models"
)

func CreateUser(conn *pgx.Conn, info models.CreateUserInfoRequest) error {
	password := HashPassword(info.Password)

	query := "INSERT INTO users (login, password, role) values ($1, $2, $3)"
	_, err := conn.Exec(context.Background(), query, info.Login, password, info.Role)

	return err
}
