package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/pedropaccola/go-webapp-alura/db"
	"github.com/pedropaccola/go-webapp-alura/products"
)

func init() {
	database := db.ConnectDB()
	defer database.Close()
	query := `create table if not exists products (
		id serial primary key,
		name varchar,
		features varchar,
		price decimal,
		quantity integer)`
	_, err := database.Exec(query)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	products.HandleRoutes()
	http.ListenAndServe(":3000", nil)
}
