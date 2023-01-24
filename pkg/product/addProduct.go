package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"io"
	"log"
	"net/http"
)

// POST USERS
func Postproduct(w http.ResponseWriter, r *http.Request) {

	e := models.Product{}

	reqBody, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &e)
	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusBadRequest)

	}

	q := `INSERT INTO
	products (product_name, description, price, quantity)
	VALUES ($1, $2, $3, $4)`

	//*************************************
	// Variables created to handle null data (int, string, bool)
	strNull := sql.NullString{}
	intNull := sql.NullInt64{}
	// boolNull := sql.NullBool{} if we need to put a null bool as a data in db

	//*************************************
	db := config.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//*************************************
	// To handle when a data is not entered because we do not have it and we want to have null in the db,
	// since age 0 cannot be (golang by default gives us 0)
	if e.Quantity == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Quantity)
	}

	if e.Description == "" {

		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = e.Description
	}

	//*************************************

	re, err := stmt.Exec(e.Product_name, strNull, e.Price, intNull)
	if err != nil {
		log.Fatal(err)
	}
	i, _ := re.RowsAffected()

	// Make a custom error with errors.New()
	if i != 1 {
		log.Println(errors.New("error: We expected affect only one row"))
	}

}
