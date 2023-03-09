package routes

import (
	"gorillamux/pkg/product"
	"gorillamux/pkg/users"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {

	// Definition of a subrouter to handle and ensure product management.
	blo := r.PathPrefix("/api/").Subrouter()

	// These endpoints can only be accessed through authentication.
	blo.HandleFunc("/product/{id:[0-9]+}", product.Getproduct).Methods("GET")
	blo.HandleFunc("/products", product.Getproducts).Methods("GET")
	blo.HandleFunc("/product/", product.Postproduct).Methods("POST")
	blo.HandleFunc("/product/{id:[0-9]+}", product.Putproduct).Methods("PUT")
	blo.HandleFunc("/product/delete/{id:[0-9]+}", product.Deleteproduct).Methods("DELETE")

	//Mongo route
	blo.HandleFunc("/mongo", product.MongoEndPoint)
	blo.HandleFunc("/desmongo", product.DisconnectMongo)
	blo.HandleFunc("/singlemongo", product.InsertOneMongo)
	blo.HandleFunc("/readlemongo", product.ReadOneMongo)

	// Middleware to check access rights to these routes
	blo.Use(users.ValidateTokenMiddleware)

}
