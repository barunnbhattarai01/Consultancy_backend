package main

import (
	"log"
	"net/http"
	"os"

	"github.com/barunnbhattarai01/consultancy_backend/controller"
	"github.com/barunnbhattarai01/consultancy_backend/intailizer"

	"github.com/gorilla/mux"
)

func init() {
	intailizer.Loadenv()
	intailizer.Connection()
}

func main() {
	Port := ":" + os.Getenv("Port")
	//instance of Api
	h := &controller.Api{Addr: Port}

	//gorilla mux
	gor := mux.NewRouter()

	//configure the http serverrr
	srv := &http.Server{
		Addr:    h.Addr,
		Handler: gor,
	}

	//routing
	gor.HandleFunc("/signup", controller.Signup).Methods("POST")
	gor.HandleFunc("/login", controller.Login).Methods("POST")

	//start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error in server")
	}

}
