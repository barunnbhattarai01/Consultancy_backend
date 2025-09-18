package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
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

	query := `insert into studentregister (name,address,phone,age,join_date) values ($1,$2,$3,$4,$5) returning id`
	var lastid int
	err := intailizer.DB.QueryRow(query, reg.Name, reg.Address, reg.Age, reg.Phone, reg.JOIN_DATE).Scan(&lastid)

	if err != nil {
		log.Println("DB create error:", err)
		http.Error(w, "error while stroing in datbase", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sucessfully added",
	})
}
