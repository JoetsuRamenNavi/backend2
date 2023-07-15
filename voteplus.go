package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func incrementVote(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	vars := mux.Vars(r)
	battleTerm := vars["battle_term"]
	storeID := vars["store_id"]

	fmt.Printf("Received request to increment vote for battle_term: %s, store_id: %s\n", battleTerm, storeID)

	var existingVoteID int
	err := db.QueryRow("SELECT votes_id FROM votes WHERE battle_term = ? AND store_id = ?", battleTerm, storeID).Scan(&existingVoteID)

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == sql.ErrNoRows {
		// If record doesn't exist, create new one with vote count 1
		stmt, err := db.Prepare("INSERT INTO votes (battle_term, store_id, vote) VALUES (?, ?, 1)")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing INSERT query: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(battleTerm, storeID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing INSERT query: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Println("Record didn't exist, inserted new one with vote count 1")
	} else {
		// If record exists, increment vote count
		stmt, err := db.Prepare("UPDATE votes SET vote = vote + 1 WHERE votes_id = ?")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error preparing UPDATE query: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(existingVoteID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing UPDATE query: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Println("Record existed, incremented vote count")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Finished processing request")
}

