package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Apdate a product register
func Putproduct(w http.ResponseWriter, r *http.Request) {
	var upActual models.Product

	vars := mux.Vars(r)
	//upActual := models.Product{}
	//upNew := models.Product{}

	id, _ := strconv.ParseInt(vars["id"], 10, 0)

	nameNull := sql.NullString{}
	descNull := sql.NullString{}
	priceNull := sql.NullFloat64{}
	quantityNull := sql.NullInt64{}

	//------------------------------------------------------------------------
	// This block is to get all the columns and check it for each vlaue change,
	// if not we maintain older value.

	sqlStatement := `SELECT product_id, product_name, description, price, quantity
	FROM products WHERE product_id=$1`
	db := config.GetConnection()
	defer db.Close()
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&upActual.Product_id,
		&nameNull,     //upActual.Product_name,
		&descNull,     //upActual.Description,
		&priceNull,    //upActual.Price,
		&quantityNull, //upActual.Quantity,
	)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(upActual)
	default:
		panic(err)
	}

	if upActual.Quantity == 0 {
		quantityNull.Valid = false
	} else {
		quantityNull.Valid = true
		quantityNull.Int64 = int64(upActual.Quantity)
	}
	upActual.Product_name = nameNull.String
	upActual.Description = descNull.String
	upActual.Price = priceNull.Float64
	upActual.Quantity = quantityNull.Int64

	upActual = isnewData(r, upActual)

	q := `UPDATE products
					SET
					product_name = $1,
					description = $2,
					price = $3,
					quantity = $4,
					updated_at = now()
					WHERE product_id = $5`

	db = config.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	re, err := stmt.Exec(upActual.Product_name, upActual.Description, upActual.Price, upActual.Quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	i, _ := re.RowsAffected()

	// Make a custom error with errors.New()
	if i != 1 {
		log.Println(errors.New("error: We expected affect only one row"))
	}

}

// --------------------------------------------------------------------------
func isnewData(r *http.Request, up models.Product) models.Product {

	datosactual := make(map[string]string)
	datosactual["product_id"] = strconv.FormatInt(up.Product_id, 10)
	datosactual["product_name"] = up.Product_name
	datosactual["description"] = up.Description
	datosactual["price"] = strconv.FormatFloat(up.Price, 'f', 0, 64)
	datosactual["quantity"] = strconv.FormatInt(up.Quantity, 10)
	fmt.Println("DATOSACTUAL", datosactual)

	// read the request to compare with actual columns values
	datosnew := make(map[string]string)
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &datosnew)
	if err != nil {
		fmt.Println(err)

	}

	// Checking if a column value has changes
	for k, val := range datosnew {

		if k != "" && val != "" {
			datosactual[k] = val
		}
	}
	up.Product_id, _ = strconv.ParseInt(datosactual["product_id"], 10, 0)
	up.Product_name = datosactual["product_name"]
	up.Description = datosactual["description"]
	up.Price, _ = strconv.ParseFloat(datosactual["price"], 64)
	up.Quantity, _ = strconv.ParseInt(datosactual["quantity"], 10, 0)
	return up

}
