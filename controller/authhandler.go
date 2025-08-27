package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// it hold info of api
type Api struct {
	Addr string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	//headers are key-value pairs that tells the client likes cookies etc
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, `{"Failed to read body"}`, http.StatusBadRequest)
		return
	}

	//hashing the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http.Error(w, `{"Failed to hash Password"}`, http.StatusBadRequest)
		return
	}

	//sign upinngg
	user := model.User{Email: body.Email, Password: string(hash)}
	result := intailizer.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, `{"failed to create user"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"User created sucessfully"}`))
}

// login logic
func Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, `{"invalid information:}`, http.StatusBadRequest)
		return
	}

	//find emaul form datbase
	var user model.User

	email := strings.ToLower(strings.TrimSpace(body.Email))
	//session create the new db session
	result := intailizer.DB.Session(&gorm.Session{PrepareStmt: false}).Where("email=?", email).First(&user)
	if result.Error != nil {
		http.Error(w, `{"Email npot found"}`, http.StatusBadRequest)
		return
	}

	//check a password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, `{"invalid password"}`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{Login sucessfully"}`))
}
