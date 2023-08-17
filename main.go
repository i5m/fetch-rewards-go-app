package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Use this to set and get all the requests points
var GlobalDBStore = make(map[string]int)

func main() {

	// Using mux to efficiently handle routing and dynamic URLs
	r := mux.NewRouter()

	// To handle new receipts -> stores points in db
	r.HandleFunc("/receipts/process", receiptProcess).Methods("POST")

	// To fetch points from id -> gets points from db
	r.HandleFunc("/receipts/{id}/points", receiptPoints).Methods("GET")

	// Start listening to the requests
	fmt.Println("Starting to listen now at: http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
