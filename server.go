package main

import (
	_ "main/routes"

	"main/controllers"
	"main/config"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Run() {
	
	fmt.Println("Listing")
    r := mux.NewRouter()
	r.HandleFunc("/login",controllers.UserLogin).Methods("POST")
	r.HandleFunc("/checkUser",controllers.CheckUser).Methods("POST")
	r.Handle("/home",config.ValidateJWT(controllers.Home)).Methods("GET")
	r.HandleFunc("/products",controllers.AllProducts).Methods("GET")
	r.HandleFunc("/getAlljobs",controllers.Allobs).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))
}