package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// it hold info of api
type Api struct {
	Addr string
}

func (a *Api) begining(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	Port := ":8080"
	//instance of Api
	h := &Api{Addr: Port}

	//gorilla mux
	gor := mux.NewRouter()

	//configure the http serverrr
	srv := &http.Server{
		Addr:    h.Addr,
		Handler: gor,
	}

	//routing
	gor.HandleFunc("/", h.begining).Methods("GET")

	//start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error in server")
	}

}
