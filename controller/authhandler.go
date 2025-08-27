package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/barunnbhattarai01/consultancy_backend/intailizer"
	"github.com/barunnbhattarai01/consultancy_backend/model"
	"github.com/golang-jwt/jwt/v5"
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

	//generate jwt
	tokenString, err := generateJWT(user.Email)
	if err != nil {
		http.Error(w, `{"Failde to generate token"}`, http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{Login sucessfully"}`))

	//return token to user
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token"` + tokenString + `"`))
}

// JWT logic
func generateJWT(email string) (string, error) {
	//jwt.Mapclaims is a payload
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	secret := os.Getenv("JWT_TOKEN")

	if secret == "" {
		secret = "default_secret"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
