package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type RequestBody struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type Result struct {
	Result int `json:"result"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Starting the Http Server!.......")

	http.HandleFunc("/add", AdditionHandler)
	http.HandleFunc("/subtract", SubtractionHandler)
	http.HandleFunc("/multiply", MultiplicationHandler)
	http.HandleFunc("/divide", DivisionHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Debug(err.Error())
	}
}

func AdditionHandler(rw http.ResponseWriter, r *http.Request) {
	var b RequestBody

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	result := Result{
		Result: b.Number1 + b.Number2,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)

	return
}

func SubtractionHandler(rw http.ResponseWriter, r *http.Request) {
	var b RequestBody

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	result := Result{
		Result: b.Number1 - b.Number2,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}

func MultiplicationHandler(rw http.ResponseWriter, r *http.Request) {
	var b RequestBody

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	result := Result{
		Result: b.Number1 * b.Number2,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}

func DivisionHandler(rw http.ResponseWriter, r *http.Request) {
	var b RequestBody

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	result := Result{
		Result: b.Number1 / b.Number2,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)
}
