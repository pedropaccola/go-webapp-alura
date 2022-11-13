package products

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts, err := SearchProducts()
	if err != nil {
		log.Printf("couldn't fetch products: %s\n", err)
	}
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		feat := r.FormValue("features")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Printf("couldn't parse product price: %s\n", err)
		}
		qty, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Printf("couldn't parse product quantity: %s\n", err)
		}
		if err := CreateProduct(name, feat, price, qty); err != nil {
			log.Printf("couldn't create new product: %s\n", err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := DeleteProduct(id); err != nil {
		log.Printf("couldn't delete the product: %s\n", err)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	p, err := EditProduct(id)
	if err != nil {
		log.Printf("couldn't edit the product: %s\n", err)
	}

	temp.ExecuteTemplate(w, "Edit", p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Printf("couldn't parse product id: %s\n", err)
		}
		name := r.FormValue("name")
		feat := r.FormValue("features")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Printf("couldn't parse product price: %s\n", err)
		}
		qty, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Printf("couldn't parse product quantity: %s\n", err)
		}
		if err := UpdateProduct(id, name, feat, price, qty); err != nil {
			log.Printf("couldn't update product: %s\n", err)
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
