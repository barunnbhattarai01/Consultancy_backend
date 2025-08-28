package main

import (
	"log"
	"net/http"
	"os"

	"github.com/barunnbhattarai01/consultancy_backend/controller"
	"github.com/barunnbhattarai01/consultancy_backend/intailizer"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func init() {
	intailizer.Loadenv()
	intailizer.Connection()
	intailizer.Syncdatabase()
}

func main() {
	Port := ":" + os.Getenv("Port")
	//instance of Api
	h := &controller.Api{Addr: Port}

	//gorilla mux
	gor := mux.NewRouter()

	//cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	//c is middleware and we passed as gor as argument to enable cors in all route that gor handle

	handler := c.Handler(gor)

	//configure the http serverrr
	srv := &http.Server{
		Addr:    h.Addr,
		Handler: handler,
	}

	//routing
	gor.HandleFunc("/signup", controller.Signup).Methods("POST")
	gor.HandleFunc("/login", controller.Login).Methods("POST")
	gor.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	//start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error in server")
	}

}
