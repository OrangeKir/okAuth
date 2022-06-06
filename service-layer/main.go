package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"okAuth/controllers"
	"okAuth/models"
)

func addDb() (*pgx.Conn, error) {
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
	logger, err := config.Build()

	if err != nil {
		fmt.Printf(err.Error())
	}

	return logger
}

func main() {
	logger := initZapLog()
	zap.ReplaceGlobals(logger)

	conn, err := addDb()

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	tokenSecret := controllers.RandStringRunes(256)

	http.HandleFunc("/get-token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		var authInfo models.AuthInfo
		json.NewDecoder(r.Body).Decode(&authInfo)

		var token, err = controllers.GetUserToken(conn, authInfo.Login, authInfo.Password, tokenSecret)

		if err != nil {
			zap.L().Info(err.Error())
			w.WriteHeader(400)
			return
		}

		json.NewEncoder(w).Encode(token)
		w.WriteHeader(200)
	})

	http.HandleFunc("/set-user", func(w http.ResponseWriter, r *http.Request) {
		var request models.CreateUserInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		err := controllers.CreateUser(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(200)
	})

	http.HandleFunc("/change-password", func(w http.ResponseWriter, r *http.Request) {
		var request models.ChangeUserPasswordRequest
		json.NewDecoder(r.Body).Decode(&request)

		err := controllers.ChangeUserPassword(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(400)
			return
		}

		w.WriteHeader(200)
	})

	http.HandleFunc("/change-role", func(w http.ResponseWriter, r *http.Request) {
		var request models.ChangeUserRoleRequest
		json.NewDecoder(r.Body).Decode(&request)

		err := controllers.ChangeUserRole(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/validate-token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		var request models.ValidateTokenRequest
		json.NewDecoder(r.Body).Decode(&request)

		isValid, err := controllers.ValidateToken(request.Token, tokenSecret)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(400)
			return
		}

		response := models.ValidateTokenResponse{IsValid: isValid}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})

	http.ListenAndServe("localhost:5080", nil)
}
