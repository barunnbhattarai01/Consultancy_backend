package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
	"gorm.io/gorm"
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

	result := intailizer.DB.Session(&gorm.Session{PrepareStmt: false, SkipDefaultTransaction: false}).Create(&interview)
	if result.Error != nil {
		log.Println("DB create error:", result.Error)
		http.Error(w, "error while creating db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sucessfully added to database",
	})

}
