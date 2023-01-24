package users

import (
	"fmt"
	"html/template"
	"net/http"
)

func PostContact(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../public/templates/contact.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		template.Execute(w, "Contact")
	}
	fmt.Fprintf(w, "Mesagge from CONTACT")

}
