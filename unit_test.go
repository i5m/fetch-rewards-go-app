package main

import (
	"testing"
)

func TestFirstParse(t *testing.T) {

	ansParse := parseFloat("22.67")
	wantParse := 22.67
	if ansParse != wantParse {
		t.Errorf("Failed first test on parseFloat :( Expected points: %f, got: %f", wantParse, ansParse)
	}

}

func TestSecondParse(t *testing.T) {

	ansParse := parseFloat("ab.cd")
	wantParse := -1.0
	if ansParse != wantParse {
		t.Errorf("Failed second test on parseFloat :( Expected points: %f, got: %f", wantParse, ansParse)
	}

}

func TestFirstReceipt(t *testing.T) {

	r := Receipt{
		Retailer: "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			Item{ ShortDescription: "Mountain Dew 12PK", Price: "6.49" },
			Item{ ShortDescription: "Emils Cheese Pizza", Price: "12.25" },
			Item{ ShortDescription: "Knorr Creamy Chicken", Price: "1.26" },
			Item{ ShortDescription: "Doritos Nacho Cheese", Price: "3.35" },
			Item{ ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00" },
		},
		Total: "35.35",
	}

	ansReceipt := calculatePoints(r)
	wantReceipt := 28
	if ansReceipt != wantReceipt {
		t.Errorf("Failed first test on Receipt :( Expected points: %d, got: %d", wantReceipt, ansReceipt)
	}

}

func TestSecondReceipt(t *testing.T) {

	r := Receipt{
		Retailer: "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			Item{ ShortDescription: "Gatorade", Price: "2.25" },
			Item{ ShortDescription: "Gatorade", Price: "2.25" },
			Item{ ShortDescription: "Gatorade", Price: "2.25" },
			Item{ ShortDescription: "Gatorade", Price: "2.25" },
		},
		Total: "9.00",
	}

	ansReceipt := calculatePoints(r)
	wantReceipt := 109
	if ansReceipt != wantReceipt {
		t.Errorf("Failed second test on Receipt :( Expected points: %d, got: %d", wantReceipt, ansReceipt)
	}

}

func TestFirstGetSet(t *testing.T) {

	toSet := 23
	ansSet := setPoints(toSet)
	wantGet := getPoints(ansSet)

	if wantGet != toSet {
		t.Errorf("Failed fisrt test on getters and setters :( Expected points: %d, got: %d", wantGet, toSet)
	}

}

func TestSecondGetSet(t *testing.T) {

	toSet := 716
	ansSet := setPoints(toSet)
	wantGet := getPoints(ansSet)

	if wantGet != toSet {
		t.Errorf("Failed second test on getters and setters :( Expected points: %d, got: %d", wantGet, toSet)
	}

}

func TestFirstStrAlphaNum(t *testing.T) {

	total := calcAlphaNum("ABCD!@#123")
	want := 7

	if total != want {
		t.Errorf("Failed fisrt test on TestFirstStrAlphaNum :( Expected points: %d, got: %d", want, total)
	}

}

func TestSecondStrAlphaNum(t *testing.T) {

	total := calcAlphaNum("!@#0R$#!")
	want := 2

	if total != want {
		t.Errorf("Failed second test on TestFirstStrAlphaNum :( Expected points: %d, got: %d", want, total)
	}

}

func TestFirstDate(t *testing.T) {

	pt, _ := addDatePoints("2022-01-01")
	want := 6

	if pt != want {
		t.Errorf("Failed second test on TestFirstDate :( Expected points: %d, got: %d", want, pt)
	}

}

func TestSecondDate(t *testing.T) {

	pt, _ := addDatePoints("2022-02-04")
	want := 0

	if pt != want {
		t.Errorf("Failed second test on TestSecondDate :( Expected points: %d, got: %d", want, pt)
	}

}

func TestFirstTime(t *testing.T) {

	pt, _ := addTimePoints("12:05")
	want := 0

	if pt != want {
		t.Errorf("Failed second test on TestFirstTime. Expected points: %d, got: %d", want, pt)
	}

}

func TestSecondTime(t *testing.T) {

	pt, _ := addTimePoints("15:05")
	want := 10

	if pt != want {
		t.Errorf("Failed second test on TestSecondTime :( Expected points: %d, got: %d", want, pt)
	}

}
