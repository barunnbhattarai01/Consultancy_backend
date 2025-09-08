package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//get token from authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"message":"missing authorization"}`, http.StatusUnauthorized)
			return
		}

		//check bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"message":"invalid authorization"}`, http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		secret := os.Getenv("JWT_TOKEN")
		if secret == "" {
			secret = "default_secret"
		}

		//validate the jwt
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sigining method")
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, `{"message":"invalid token"}`, http.StatusUnauthorized)
			return
		}
		//valid process to the next handler
		next(w, r)

	}
}
