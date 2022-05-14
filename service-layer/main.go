package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"okAuth/controllers"
	"okAuth/models"
)

func main() {
	conn, _ := AddDb()
	logger := initZapLog()
	zap.ReplaceGlobals(logger)

	http.HandleFunc("/get-token", func(w http.ResponseWriter, r *http.Request) {
		var authInfo models.AuthInfo
		json.NewDecoder(r.Body).Decode(&authInfo)

		var token, _ = controllers.GetUserToken(conn, authInfo.Login, authInfo.Password)

		json.NewEncoder(w).Encode(token)
	})

	http.HandleFunc("/set-user", func(w http.ResponseWriter, r *http.Request) {
		var info models.CreateUserInfo
		json.NewDecoder(r.Body).Decode(&info)

		controllers.CreateUser(conn, info)
	})

	http.ListenAndServe("localhost:8181", nil)
}
