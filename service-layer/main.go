package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"okAuth/controllers"
	"okAuth/models"
)

func main() {
	logger := initZapLog()
	zap.ReplaceGlobals(logger)

	conn, err := AddDb()

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	tokenSecret := controllers.RandStringRunes(256)

	http.HandleFunc("/get-token", func(w http.ResponseWriter, r *http.Request) {
		var authInfo models.AuthInfo
		json.NewDecoder(r.Body).Decode(&authInfo)

		var token, err = controllers.GetUserToken(conn, authInfo.Login, authInfo.Password, tokenSecret)

		if err != nil {
			zap.L().Error(err.Error())
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

	http.ListenAndServe("localhost:8181", nil)
}
