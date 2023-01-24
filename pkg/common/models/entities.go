package models

import "time"

type Product struct {
	Product_id   int64   `json:"product_id"`
	Product_name string  `json:"product_name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Quantity     int64   `json:"quantity,omitempty"`
	//CreateAt time.Time `json:"createdat"`
	UpdateAt time.Time `json:"updated_at"`
}

type User struct {
	User_ID   int    `json:"user_id"`
	User_Name string `json:"user_name"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
	Age       int    `json:"user_age"`
	Active    bool   `json:"active"`
}
