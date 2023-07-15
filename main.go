package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type Store struct {
	StoreID      int    `json:"store_id"`
	StoreName    string `json:"store_name"`
	Tel          string `json:"tel"`
	Address      string `json:"address"`
	StoreURL     string `json:"store_url"`
	Image        string `json:"image"`
	EntryItem    string `json:"entry_item"`
	EntryPrice   string `json:"entry_price"`
	EntryComment string `json:"entry_comment"`
	UserID       int    `json:"user_id"`
}

type Vote struct {
	VotesID   int    `json:"votes_id"`
	BattleTerm string `json:"battle_term"`
	Vote       int    `json:"vote"`
	StoreID    int    `json:"store_id"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "laravel"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := dbConn()
	defer db.Close()

	router := mux.NewRouter()

	// Route

	router.HandleFunc("/users/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/stores", getStores).Methods("GET")
	router.HandleFunc("/storesvotes/{battle_term}", getStoresWithVotes).Methods("GET")
	router.HandleFunc("/incrementvote/{battle_term}/{store_id}", incrementVote).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
// /users
func getUsers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Nickname, &u.Password, &u.Mail)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, u)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}


// /user/1
func getUser(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    id := mux.Vars(r)["id"]

    row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
    var u User
    err := row.Scan(&u.ID, &u.Nickname, &u.Password, &u.Mail)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    json.NewEncoder(w).Encode(u)
}

// /stores
func getStores(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    rows, err := db.Query("SELECT * FROM stores")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var stores []Store
    for rows.Next() {
        var s Store
	err := rows.Scan(&s.StoreID, &s.StoreName, &s.Tel, &s.Address, &s.StoreURL, &s.Image, &s.EntryItem, &s.EntryPrice, &s.EntryComment, &s.UserID)
	if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        stores = append(stores, s)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(stores)
}

