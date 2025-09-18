package controller

import (
	"encoding/json"
	"log"
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

	query := `insert into interviewregister (name,address,date,images) values($1,$2,$3,$4) returning id`
	var lastid int
	err := intailizer.DB.QueryRow(query, interview.Name, interview.Address, interview.Date, interview.Images).Scan(&lastid)
	if err != nil {
		log.Println("DB create error:", err)
		http.Error(w, "error while creating db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sucessfully added to database",
	})

}
