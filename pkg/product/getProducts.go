package product

import (
	"database/sql"
	"fmt"
	"gorillamux/pkg/common/config"
	"gorillamux/pkg/common/models"
	"html/template"
	"log"
	"net/http"
)

/********************************************
*** REMEMBER TO MAKE POSSIBLE PAGINATION ****
***         LIMIT $1 OFFSET $2           ****
*********************************************
 */
// GET USERS
func Getproducts(w http.ResponseWriter, r *http.Request) {
	Products := []models.Product{}

	q := `SELECT product_id, product_name, description, price, quantity
			FROM products order by product_id asc`

	nameNull := sql.NullString{}
	descNull := sql.NullString{}
	priceNull := sql.NullFloat64{}
	quantityNull := sql.NullInt64{}

	//timeNull := pq.NullTime{} // here we use pq package to handle time null
	//str1Null := sql.NullString{}
	//boolNull := sql.NullBool{}

	db := config.GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	// Next() it returns a false if there is no new data, in this case new row.
	// and if false, the for loop ends.
	for rows.Next() {

		e := models.Product{}
		err = rows.Scan( // Use 'rows.Scan()' from http://go-database-sql.org
			&e.Product_id,
			&nameNull,
			&descNull,
			&priceNull,
			&quantityNull,
			//&boolNull,
			////&e.CreateAt,
			//&timeNull,
			// Before we sent --> &e.UpdateAt, but it gave an error because
			// in the queries it does not read the null time data.
		)
		if err != nil {
			log.Println(err)
		}

		if e.Quantity == 0 {
			quantityNull.Valid = false
		} else {
			quantityNull.Valid = true
			quantityNull.Int64 = int64(e.Quantity)
		}

		e.Product_name = nameNull.String
		e.Description = descNull.String
		e.Price = priceNull.Float64
		e.Quantity = quantityNull.Int64
		//e.UpdateAt = timeNull.Time

		Products = append(Products, e)

	}

	templ, err := template.ParseFiles("../public/templates/index.html")
	if err != nil {
		fmt.Fprintf(w, " PAGE NOT FOUND..!")
	} else {
		templ.Execute(w, Products)
	}
}
