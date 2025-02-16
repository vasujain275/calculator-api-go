package handlers

import (
	"encoding/json"
	"github.com/vasujain275/calculator-api-go/logger"
	"net/http"
)

func SumHandler(rw http.ResponseWriter, r *http.Request) {
	var numbers []int
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(rw, "Invalid Request Body", http.StatusBadRequest)
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	logger.Log.Info("Got Sum Request")
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]int{"result": sum})
}
