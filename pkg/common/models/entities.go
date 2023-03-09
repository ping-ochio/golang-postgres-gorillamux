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
	User_ID   int    `json:"user_id" bson:"-"`
	User_Name string `json:"user_name" bson:"-"`
	Email     string `json:"email" bson:"-"`
	Name      string `json:"name" bson:"name"`
	Surname   string `json:"surname" bson:"surname"`
	Password  string `json:"password" bson:"-"`
	Age       int    `json:"age" bson:"age"`
	Active    *bool  `json:"active,omitempty" bson:"-"`
}
