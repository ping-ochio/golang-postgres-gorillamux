package users

import (
	"database/sql"
	"errors"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"log"
)

/*

esta funcion debe recibir una struct del tipo user con la password
previamente hasheada, quiere decir que el generador de hash debe ser llamado
antes de que envien los datos para ser guardados en base de datos.

*/

func adduser(u models.User) error {

	//e := models.User{}

	/*reqBody, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &e)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusBadRequest)

	}*/

	var err error
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		log.Fatal(err)
	}

	q := `INSERT INTO
	users (user_name, name, surname, email, password, age, active)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	//*************************************
	// Variables created to handle null data (int, string, bool)
	userNameNull := sql.NullString{}
	nameNull := sql.NullString{}
	surnameNull := sql.NullString{}
	emailNull := sql.NullString{}
	passNull := sql.NullString{}
	ageNull := sql.NullInt64{}
	activeboolNull := sql.NullBool{} //if we need to put a null bool as a data in db

	//*************************************
	db := config.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//*************************************
	// To handle when a data field is not entered because we do not have it and we want to have null in the db,
	// since age 0 cannot be (golang by default gives us 0)

	if u.User_Name == "" {
		userNameNull.Valid = false
	} else {
		userNameNull.Valid = true
		userNameNull.String = u.User_Name
	}

	if u.Name == "" {
		nameNull.Valid = false
	} else {
		nameNull.Valid = true
		nameNull.String = u.Name
	}

	if u.Surname == "" {
		surnameNull.Valid = false
	} else {
		surnameNull.Valid = true
		surnameNull.String = u.Surname
	}

	if u.Email == "" {
		emailNull.Valid = false
	} else {
		emailNull.Valid = true
		emailNull.String = u.Email
	}

	if u.Password == "" {
		passNull.Valid = false
	} else {
		passNull.Valid = true
		passNull.String = u.Password
	}

	if u.Age == 0 {
		ageNull.Valid = false
	} else {
		ageNull.Valid = true
		ageNull.Int64 = int64(u.Age)
	}

	//	NOTE: u.Active --> "Active *bool" in struct User.
	//	"u.Active" is a pointer, is "nil" if empty
	//	and "*u.Active" its content, in this case "true" or "false"
	if u.Active == nil {
		activeboolNull.Valid = false
	} else {
		activeboolNull.Valid = true
		activeboolNull.Bool = *u.Active
	}

	//*************************************

	re, err := stmt.Exec(
		&userNameNull,
		&nameNull,
		&surnameNull,
		&emailNull,
		&passNull,
		&ageNull,
		&activeboolNull,
		//&u.Active,
	)
	if err != nil {
		log.Fatal(err)
	}
	i, _ := re.RowsAffected()

	// Make a custom error with errors.New()
	if i != 1 {
		log.Println(errors.New("error: We expected affect only one row"))
	}
	return nil
}
