package users

import (
	"fmt"
	"gorillamux/pkg/common/models"
	"html/template"
	"net/http"
)

func PostContact(w http.ResponseWriter, r *http.Request) {
	p := models.User{}
	template, err := template.ParseFiles("../public/templates/contact.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		p.Name = r.FormValue("name")
		p.Surname = r.FormValue("surname")
		p.Password = r.FormValue("password")
		template.Execute(w, "login")
		fmt.Println(p)
	}
}
