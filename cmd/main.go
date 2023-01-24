package main

import (
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/db"
	"gorillamux/pkg/product"
	"gorillamux/pkg/users"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// var id int

func init() {

	config.GetConnection()
	db.LoadEnvVariables()
}
func main() {

	//--- Create a new router and some rotes
	r := mux.NewRouter() //.StrictSlash(false)

	r.HandleFunc("/api/product/{id:[0-9]+}", product.Getproduct).Methods("GET")
	r.HandleFunc("/api/products", product.Getproducts).Methods("GET")
	r.HandleFunc("/api/product/", product.Postproduct).Methods("POST")
	r.HandleFunc("/api/product/{id:[0-9]+}", product.Putproduct).Methods("PUT")
	r.HandleFunc("/api/product/delete/{id:[0-9]+}", product.Deleteproduct).Methods("DELETE")

	r.HandleFunc("/contact", users.PostContact).Methods("GET")
	r.HandleFunc("/login", users.PostLogin).Methods("GET")

	fs := http.FileServer(http.Dir("../public/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//--- Create a new server http using server struct.
	server := &http.Server{
		Addr:           ":7551",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max 1 MegaBytes = 1048576 = (1 << 20)
	}

	log.Println("Listening...")

	// but if we want to stop the server when an error occurs we can put.
	log.Fatal(server.ListenAndServe())
}
