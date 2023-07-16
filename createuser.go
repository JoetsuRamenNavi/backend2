// userHandlers.go
package main

import (
	"encoding/json"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	var user User

	// ユーザー情報をデコードします
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert
	insForm, err := db.Prepare("INSERT INTO users(nickname, password, mail) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insForm.Exec(user.Nickname, user.Password, user.Mail)

	w.WriteHeader(http.StatusCreated)
}

