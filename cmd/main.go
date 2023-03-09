package main

import (
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/db"
	"gorillamux/pkg/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func init() {

	config.GetConnection()
	db.LoadEnvVariables()
}
func main() {

	//--- Create a new router and some rotes
	r := mux.NewRouter().StrictSlash(false)

	routes.UserRoutes(r)
	//blo := r.PathPrefix("/api/").Subrouter()
	routes.ProductRoutes(r)
	//blo.Use(users.ValidateTokenMiddleware)

	//r.Use(users.ValidateTokenMiddleware)

	fs := http.FileServer(http.Dir("../public/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//--- Create a new server http using server struct.
	server := &http.Server{
		Addr:           os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max 1 MegaByte = 1048576 = (1 << 20)
	}

	log.Println("Listening...")

	log.Fatal(server.ListenAndServe())
}
