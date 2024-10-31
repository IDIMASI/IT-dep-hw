package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func addData(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		os.Exit(1)
	}
	fmt.Fprintf(w, "Data added: %s", name)
}

func getData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		panic(1)
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			http.Error(w, "Failed to scan data", http.StatusInternalServerError)
			panic(1)
		}
		fmt.Fprintf(w, "User: %s\n", name)
	}
}

func main() {
	http.HandleFunc("/add", addData)
	http.HandleFunc("/get", getData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
