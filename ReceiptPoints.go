package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

// PointsResponse represents the JSON response structure.
type PointsResponse struct {
	Points int `json:"points"`
}

// getPoints retrieves the points for a given ID from the GlobalDBStore variable.
// If the ID exists, it returns the associated points; otherwise, it returns 0.
func getPoints(id string) int {
	points, exists := GlobalDBStore[id]
	if exists {
		return points
	}
	return -1
}

// receiptPoints is an HTTP handler function that responds to a request for retrieving points.
func receiptPoints(w http.ResponseWriter, r *http.Request) {

	// Extracting route variables (such as "id") from the request.
	vars := mux.Vars(r)
	// Extracting the "id" variable from the route.
	id := vars["id"]

	// Retrieving points for the provided ID using the getPoints function.
	points := getPoints(id)

	if points == -1 {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return	
	}

	// Creating a PointsResponse with the retrieved points.
	response := PointsResponse{Points: points}

	fmt.Println("Fetched receipt; Available Points - " + strconv.Itoa(points) + "; For ID - " + id)

	// Setting the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Encoding the response as JSON and sending it to the client.
	json.NewEncoder(w).Encode(response)

}
