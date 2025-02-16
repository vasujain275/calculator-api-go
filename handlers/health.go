package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vasujain275/calculator-api-go/logger"
	"github.com/vasujain275/calculator-api-go/models"
)

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	health := models.Health{
		Status:  "up",
		Version: "v1",
	}
	logger.Log.Info("Fectched HeathCheck for the api")
	rw.Header().Set("Content-type", "application/json")
	json.NewEncoder(rw).Encode(health)
}
