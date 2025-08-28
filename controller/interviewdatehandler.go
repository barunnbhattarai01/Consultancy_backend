package controller

import (
	"encoding/json"
	"net/http"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
)

func InterviewDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Only post method is allowed", http.StatusBadRequest)
		return
	}

	var interview model.InterviewDate
	if err := json.NewDecoder(r.Body).Decode(&interview); err != nil {
		http.Error(w, "Error while fetching", http.StatusBadRequest)
		return
	}

	result := intailizer.DB.Create(&interview)
	if result.Error != nil {
		http.Error(w, "error while creating db", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sucessfully added to database",
	})

}
