package product

import (
	"database/sql"
	"fmt"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GET PRODUCT
func Getproduct(w http.ResponseWriter, r *http.Request) {

	product := models.Product{}
	vars := mux.Vars(r)
	var err1 error
	product.Product_id, err1 = strconv.ParseInt(vars["id"], 10, 0)
	if err1 != nil {
		log.Println(err1)
	}

	nameNull := sql.NullString{}
	descNull := sql.NullString{}
	priceNull := sql.NullFloat64{}
	quantityNull := sql.NullInt64{}

	//------------------------------------------------------------------------
	sqlStatement := `SELECT product_id, product_name, description, price, quantity
	FROM products WHERE product_id=$1`
	db := config.GetConnection()
	defer db.Close()

	row := db.QueryRow(sqlStatement, product.Product_id)
	err := row.Scan(
		&product.Product_id,
		&nameNull,
		&descNull,
		&priceNull,
		&quantityNull,
	)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(product)
	default:
		panic(err)
	}
	if product.Quantity == 0 {
		quantityNull.Valid = false
	} else {
		quantityNull.Valid = true
		quantityNull.Int64 = int64(product.Quantity)
	}

	product.Product_name = nameNull.String
	product.Description = descNull.String
	product.Price = priceNull.Float64
	product.Quantity = quantityNull.Int64
	fmt.Println("FINAL: ", product)

	template, err := template.ParseFiles("../public/templates/singleproduct.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		template.Execute(w, product)
	}

}
