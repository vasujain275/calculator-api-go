package main

import (
	"github.com/vasujain275/calculator-api-go/handlers"
	"github.com/vasujain275/calculator-api-go/logger"
	"net/http"
)

func main() {
	logger.Log.Info("Starting the Http Server!.......")

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/add", handlers.AdditionHandler)
	http.HandleFunc("/subtract", handlers.SubtractionHandler)
	http.HandleFunc("/multiply", handlers.MultiplicationHandler)
	http.HandleFunc("/divide", handlers.DivisionHandler)
	http.HandleFunc("/sum", handlers.SumHandler)

	fs := http.FileServer(http.Dir("./swagger-ui"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log.Debug(err.Error())
	}
}
