package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type StatsResponse struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type InteractionRequest struct {
	IsLike bool `json:"is_like"`
}

func main() {
	connStr := "user=postgres password=mysecret dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	initSQL, err := os.ReadFile("db/init.sql")
	if err != nil {
		log.Printf("Warning: Can't read init.sql file: %v", err)
	} else {
		_, err = db.Exec(string(initSQL))
		if err != nil {
			log.Fatalf("Initialize DB failed: %v", err)
		}
		fmt.Println("Initialize DB Success")
	}

	http.HandleFunc("/api/interactions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost {
			var body InteractionRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				return
			}

			_, err := db.Exec("INSERT INTO interaction_history (is_like) VALUES ($1)", body.IsLike)
			if err != nil {
				log.Print("Insert failed: %v", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
		}

		if r.Method == http.MethodGet || r.Method == http.MethodPost {
			query := `SELECT
						count(*) FILTER (WHERE is_like = TRUE), count(*) FILTER (WHERE is_like = FALSE)
						FROM interaction_history`

			var stats StatsResponse

			err = db.QueryRow(query).Scan(&stats.Likes, &stats.Dislikes)
			if err != nil {
				log.Printf("Query failing %v", err)
				http.Error(w, "Query Error", http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(stats)

		}

	})

	http.HandleFunc("/api/reset", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodPost {
			_, err := db.Exec("TRUNCATE TABLE interaction_history RESTART IDENTITY")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Reset successful"})
		}
	})
	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
