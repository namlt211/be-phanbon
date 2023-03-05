package routers

import (
	"green/controllers/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() {
	post := ":8080"
	r := mux.NewRouter()
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(post, r))

}