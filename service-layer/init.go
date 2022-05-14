package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func AddDb() (*pgx.Conn, error) {
	connStr := "postgresql://postgres:1@localhost:5432/okAuth"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		conn.Close(context.Background())
	}

	return conn, err
}

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}
