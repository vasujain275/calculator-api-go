package handlers

import (
	"encoding/json"
	"github.com/vasujain275/calculator-api-go/logger"
	"github.com/vasujain275/calculator-api-go/models"
	"net/http"
)

func AdditionHandler(rw http.ResponseWriter, r *http.Request) {
	var b models.RequestBody

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		logger.Log.Debug(err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	result := models.Result{
		Result: b.Number1 + b.Number2,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(result)

	return
}
