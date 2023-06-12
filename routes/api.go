package routes

import (
    "github.com/gorilla/mux"
	
	"main/controllers"
)

func New() {
	r := mux.NewRouter()
	r.HandleFunc("/login",controllers.UserLogin).Methods("POST")
	
}