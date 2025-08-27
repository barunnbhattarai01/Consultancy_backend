package controller

import (
	"net/http"
)

// it hold info of api
type Api struct {
	Addr string
}

func (a *Api) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
