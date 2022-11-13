package products

import (
	"github.com/pedropaccola/go-webapp-alura/db"
)

type Product struct {
	ID       int
	Name     string
	Features string
	Price    float64
	Quantity int
}

func SearchProducts() ([]Product, error) {
	database := db.ConnectDB()
	defer database.Close()

	rows, err := database.Query("select * from products order by id asc")
	if err != nil {
		return nil, err
	}

	p := Product{}
	products := []Product{}

	for rows.Next() {
		var id, qty int
		var name, feat string
		var price float64

		err = rows.Scan(&id, &name, &feat, &price, &qty)
		if err != nil {
			return nil, err
		}

		p.ID = id
		p.Name = name
		p.Features = feat
		p.Price = price
		p.Quantity = qty

		products = append(products, p)
	}
	return products, nil
}

func CreateProduct(name, feat string, price float64, qty int) error {
	database := db.ConnectDB()
	defer database.Close()

	query := `
	insert into products (name, features, price, quantity)
	values ($1, $2, $3, $4)`

	_, err := database.Query(
		query,
		name,
		feat,
		price,
		qty)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id string) error {
	database := db.ConnectDB()
	defer database.Close()

	query := `delete from products where id = $1`
	_, err := database.Query(
		query,
		id)
	if err != nil {
		return err
	}
	return nil
}

func EditProduct(id string) (Product, error) {
	database := db.ConnectDB()
	defer database.Close()

	p := Product{}

	query := `select * from products where id = $1`
	rows, err := database.Query(
		query,
		id)
	if err != nil {
		return p, err
	}

	for rows.Next() {
		var id, qty int
		var name, feat string
		var price float64

		err = rows.Scan(&id, &name, &feat, &price, &qty)
		if err != nil {
			return p, err
		}

		p.ID = id
		p.Name = name
		p.Features = feat
		p.Price = price
		p.Quantity = qty
	}
	return p, nil
}

func UpdateProduct(id int, name, feat string, price float64, qty int) error {
	database := db.ConnectDB()
	defer database.Close()

	query := `
	update products set name=$1, features=$2, price=$3, quantity=$4 
	where id=$5`

	_, err := database.Query(
		query,
		name,
		feat,
		price,
		qty,
		id)
	if err != nil {
		return err
	}
	return nil
}
