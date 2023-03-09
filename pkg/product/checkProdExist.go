package product

import (
	"fmt"
	"gorillamux/pkg/common/config"
)

func CheckProdExists(prodId int64) bool {

	sqlStatement := `SELECT product_id
	FROM products WHERE product_id=$1`
	db := config.GetConnection()
	defer db.Close()

	row := db.QueryRow(sqlStatement, prodId)
	err := row.Scan(
		&prodId,
	)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
