package users

import (
	"fmt"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var u models.User

	template, err := template.ParseFiles("../public/templates/newUser.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		template.Execute(w, ".")

		u.User_Name = r.FormValue("user_name")
		u.Name = r.FormValue("name")
		u.Surname = r.FormValue("surname")
		u.Email = r.FormValue("email")
		u.Password = r.FormValue("password")
		u.Age, _ = strconv.Atoi(r.FormValue("age"))
		active := false
		u.Active = &active

		if !isEmpty(u) {

			if ifExist(u) {

				fmt.Fprintf(w, "invalid data")
				return

			}
			err = adduser(u) // this block save the user in database
			if err != nil {

				fmt.Println(err)

			}
		} else {
			log.Println("some values are empty")
		}
	}
}

// Check if user data is empty.
// Does the same thing as the "preventDefault()" javascript method.
func isEmpty(v models.User) bool {
	if (v.Name == "") && (v.User_Name == "") &&
		(v.Email == "") && (v.Password == "") &&
		(v.Surname == "") {
		return true
	}
	return false
}

// Check if this user email already exist
func ifExist(v models.User) bool {

	sqlStatement := `SELECT email
	FROM users WHERE email=$1`
	db := config.GetConnection()
	defer db.Close()

	row := db.QueryRow(sqlStatement, v.Email)
	err := row.Scan(
		&v.Email,
	)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
