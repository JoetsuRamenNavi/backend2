package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type StoreVotes struct {
	ID           int    `json:"store_id"`
	StoreName    string `json:"store_name"`
	Tel          string `json:"tel"`
	Address      string `json:"address"`
	StoreURL     string `json:"store_url"`
	Image        string `json:"image"`
	EntryItem    string `json:"entry_item"`
	EntryPrice   string `json:"entry_price"`
	EntryComment string `json:"entry_comment"`
	UserID       int    `json:"id"`
	Votes        int    `json:"votes"`
}

func getStoresWithVotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    defer db.Close()

    vars := mux.Vars(r)
    battleTerm := vars["battle_term"]

    rows, err := db.Query("SELECT stores.store_id, store_name, tel, address, store_url, image, entry_item, entry_price, entry_comment, id, SUM(vote) as votes FROM stores LEFT JOIN votes ON stores.store_id = votes.store_id WHERE votes.battle_term = ? GROUP BY stores.store_id", battleTerm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var stores []StoreVotes
    for rows.Next() {
        var s StoreVotes
        err := rows.Scan(&s.ID, &s.StoreName, &s.Tel, &s.Address, &s.StoreURL, &s.Image, &s.EntryItem, &s.EntryPrice, &s.EntryComment, &s.UserID, &s.Votes)
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

