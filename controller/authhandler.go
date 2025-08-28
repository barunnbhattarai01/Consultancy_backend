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
		http.Error(w, `{"message":"Failed to read body"}`, http.StatusBadRequest)
		return
	}

	//hashing the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http.Error(w, `{"message":"Failed to hash Password"}`, http.StatusBadRequest)
		return
	}

	//sign upinngg
	user := model.User{Email: body.Email, Password: string(hash)}
	result := intailizer.DB.Session(&gorm.Session{PrepareStmt: false, SkipDefaultTransaction: false}).Create(&user)

	if result.Error != nil {
		http.Error(w, `{"message":"failed to create user"}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

// login logic
func Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, `{"message":"invalid information:}`, http.StatusBadRequest)
		return
	}

	//find emaul form datbase
	var user model.User

	email := strings.ToLower(strings.TrimSpace(body.Email))
	//session create the new db session
	result := intailizer.DB.Where("email=?", email).First(&user)
	if result.Error != nil {
		http.Error(w, `{"message":"Email not found"}`, http.StatusBadRequest)
		return
	}

	//check a password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, `{"message":"invalid password"}`, http.StatusBadRequest)
		return
	}

	//generate jwt
	tokenString, err := generateJWT(user.Email)
	if err != nil {
		http.Error(w, `{"message":"Fail to generate token"}`, http.StatusBadRequest)
		return
	}

	//return token to user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   tokenString,
	})

}

// JWT logic
func generateJWT(email string) (string, error) {
	//jwt.Mapclaims map that stores the data inside jwt called payload
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	secret := os.Getenv("JWT_TOKEN")

	if secret == "" {
		secret = "default_secret"
	}

	//siginingMethodHS256 is a algorthims for sigining(symetic,simple and fast)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
