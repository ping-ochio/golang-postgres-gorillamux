package users

import (
	"database/sql"
	"fmt"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"html/template"
	"net/http"
)

type JWTResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "Application/json")

	//reqBody, _ := io.ReadAll(r.Body)

	//err := json.Unmarshal(reqBody, &u)
	//if err != nil {
	//	http.Error(w, "Cannot decode json", http.StatusBadRequest)
	//		return
	//}*/
	template, err := template.ParseFiles("../public/templates/login.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		var u models.User
		u.User_Name = r.FormValue("user_name")
		u.Password = r.FormValue("password")
		fmt.Println(u)

		password := getUserHash(u.User_Name)
		resul, err := CheckPassword(password, u.Password)
		if err == nil {

			fmt.Println(" Welcome", u.Name, u.Surname)
			cookie := &http.Cookie{
				Name:     "bearer",
				Value:    resul.Token + "Refresh=" + resul.Refresh,
				Secure:   true,
				HttpOnly: true,
				Path:     "/api/",
			}
			http.SetCookie(w, cookie)
			//template.Execute(w, "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			//json.NewEncoder(w).Encode(&resul)

		} else {
			fmt.Println("invalid credentials")
			template.Execute(w, "/login/")

		}
		if u.Password != "" {

			u.Password, err = HashPassword(u.Password)
			if err != nil {

				fmt.Println(err)

			}
		}

		/*err = adduser(u) // this block save the user in database
		if err != nil {

			fmt.Println(err)

			}*/
	}

}
func getUserHash(u string) string {
	//product := models.User{}

	passNull := sql.NullString{}

	//------------------------------------------------------------------------
	sqlStatement := `SELECT password
	FROM users WHERE user_name=$1`
	db := config.GetConnection()
	defer db.Close()

	row := db.QueryRow(sqlStatement, u)
	err := row.Scan(
		&passNull,
	)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		return passNull.String
	default:
		panic(err)
	}

	return ""

}
