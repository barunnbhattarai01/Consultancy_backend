package controller

import (
	"encoding/json"
	"net/http"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Only POST method allowed",
		})
		return
	}

	var reg model.Register

	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "error while decoding data from body",
		})
		return
	}

	query := `insert into studentregister (name,address,phone,age,join_date) values ($1,$2,$3,$4,$5) returning id`
	var lastid int
	err = intailizer.DB.QueryRow(query, reg.Name, reg.Address, reg.Phone, reg.Age, reg.JOIN_DATE).Scan(&lastid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "error while sending data to db",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sucessfully added",
	})
}
