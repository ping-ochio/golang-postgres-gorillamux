package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("../public/templates/index.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		template.Execute(w, "/")
	}
}
