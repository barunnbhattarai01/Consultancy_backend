package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
	"gorm.io/gorm"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only post method allowed", http.StatusBadRequest)
		return
	}

	var reg model.Register

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "failed to fetech data from body", http.StatusBadRequest)
		return
	}

	//create a db
	result := intailizer.DB.Session(&gorm.Session{PrepareStmt: false}).Create(&reg)

	if result.Error != nil {
		log.Println("DB create error:", result.Error)
		http.Error(w, "error while stroing in datbase", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sucessfully added",
	})
}
