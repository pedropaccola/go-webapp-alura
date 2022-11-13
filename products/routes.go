package products

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/newproduct", NewProduct)
	http.HandleFunc("/", Index)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
}
