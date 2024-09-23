package main

import (
	"fmt"
	"encoding/json"
	"math"
	"net/http"
	"unicode"
	"strings"
	"time"
	"strconv"
	"github.com/google/uuid"
)

// Receipt represents the structure of a receipt
type Receipt struct {
	Retailer     string   `json:"retailer"`
	PurchaseDate string   `json:"purchaseDate"`
	PurchaseTime string   `json:"purchaseTime"`
	Items        []Item   `json:"items"`
	Total        string   `json:"total"`
}

// Item represents an item on the receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// IdResponse represents the JSON response structure containing a unique UUID format ID.
type IdResponse struct {
	Id string `json:"id"`
}

// setPoints generates a new ID, associates it with the provided points in the GlobalDBStore, and returns the ID.
func setPoints(points int) string {
	id := uuid.New().String()
	GlobalDBStore[id] = points
	return id
}

// One point for every alphanumeric character in the retailer name.
func calcAlphaNum(retailer string) int {
	total := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			total++
		}
	}
	return total
}

// 6 points if the day in the purchase date is odd.
func addDatePoints(d string) (points int, err error) {
	purchaseDate, err := time.Parse("2006-01-02", d)
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else if purchaseDate.Day() % 2 != 0 {
		return 6, nil
	} else {
		return 0, nil
	}
	
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func addTimePoints(t string) (points int, err error) {
	purchaseTime, err := time.Parse("15:04", t)
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else if purchaseTime.Hour() >= 14 && purchaseTime.Hour() <= 16 {
		return 10, nil
	} else {
		return 0, nil
	}
}
// calculatePoints calculates and returns the points based on the provided receipt information.
func calculatePoints(receipt Receipt) int {

	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" {
		return -1
	}

	points := 0

	points += calcAlphaNum(receipt.Retailer)

	dpt, err := addDatePoints(receipt.PurchaseDate)
	if err != nil {
		return -1
	}
	points += dpt

	tpt, err := addTimePoints(receipt.PurchaseTime)
	if err != nil {
		return -1
	}
	points += tpt


	total := parseFloat(receipt.Total)
	if total == -1 {
		return -1
	}
	_, dec := math.Modf(total)

	// 50 points if the total is a round dollar amount with no cents.
	if dec == 0 {
		points += 50
	}
	
	// 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += (5 * int(len(receipt.Items) / 2))

	for _, item := range receipt.Items {

		if item.ShortDescription == "" || item.Price == "" {
			return -1
		}

		price := parseFloat(item.Price)
		if price == -1 {
			return -1
		}

		// trimmed length of the item description
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		// is a multiple of 3
		if trimmedLen % 3 == 0 {
			// multiply the price by 0.2 and round up to the nearest integer
			points += int(math.Ceil(price * 0.2))
		}
	}

	return points

}

// parseFloat converts a string to a floating-point number.
func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return val
}

// receiptProcess is an HTTP handler function that processes a receipt and calculates points.
func receiptProcess(w http.ResponseWriter, r *http.Request) {

	var receipt Receipt

	err := json.NewDecoder(r.Body).Decode(&receipt)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	points := calculatePoints(receipt)

	if points == -1 {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	id := setPoints(points)
	response := IdResponse{Id: id}

	fmt.Println("Saved a new receipt; Calculated Points - " + strconv.Itoa(points) + "; Generated ID - " + id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
