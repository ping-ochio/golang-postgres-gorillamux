package product

import (
	"errors"
	"gorillamux/pkg/common/config"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DELETE PRODUCT
func Deleteproduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		log.Println(err)
	}
	q := `DELETE FROM products WHERE product_id = $1`

	db := config.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	re, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	i, _ := re.RowsAffected()

	// Make a custom error with errors.New()
	if i != 1 {
		log.Println(errors.New("error: We expected affect only one row"))
	}
}
