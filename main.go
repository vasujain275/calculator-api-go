package main

import (
	"net/http"

	"github.com/vasujain275/calculator-api-go/handlers"
	"github.com/vasujain275/calculator-api-go/logger"
)

func main() {
	logger.Log.Info("Starting the Http Server!.......")

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/add", handlers.AdditionHandler)
	http.HandleFunc("/subtract", handlers.SubtractionHandler)
	http.HandleFunc("/multiply", handlers.MultiplicationHandler)
	http.HandleFunc("/divide", handlers.DivisionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log.Debug(err.Error())
	}
}
