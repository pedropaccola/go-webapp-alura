package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name     string
	Features string
	Price    float64
	Quantity int
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":3000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Shirt", Features: "Black", Price: 29, Quantity: 10},
		{Name: "Pants", Features: "Blue", Price: 49, Quantity: 5},
		{Name: "Jacket", Features: "Denim", Price: 199, Quantity: 2},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
