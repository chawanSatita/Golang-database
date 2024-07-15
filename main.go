package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "postgres"
	password     = "aom10022"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)

	}
	db = sdb
	err = db.Ping()

	if err != nil {
		log.Fatal(err)

	}

	//connection database successful
	print("connection database successful \n")

	products, err := getProducts()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get Product Successful", products)

	///fmt.Println("Get successful", product)

	// err = createProduct(&Product{Name: "Go Product2", Price: 444})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// print("create succesful")

}

// func createProduct(product *Product) error {

// 	_, err := db.Exec("INSERT INTO public.products(name, price)VALUES ($1, $2);",
// 		product.Name,
// 		product.Price)

// 	return err
// }

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id,name,price FROM products WHERE id = $1;",
		id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}
	return p, nil

}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE public.products SET name = $1 , price = $2 WHERE id = $3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price from products;")

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil

}
