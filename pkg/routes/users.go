package routes

import (
	"gorillamux/pkg/trafic"
	"gorillamux/pkg/users"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/getip", trafic.GetIpAddress).Methods("GET")
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/contact", users.PostContact).Methods("GET")
	r.HandleFunc("/login/", users.PostLogin).Methods("GET", "POST")
	r.HandleFunc("/signup", users.Signup).Methods("GET", "POST")
}
