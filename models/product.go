package models

import (
	"fmt"

	"io.nedram/lolja/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func RetrieveProducts() []Product {
	products := []Product{}
	var name, description string
	var id, quantity int
	var price float64

	connectionDB := db.DbConnector()
	defer connectionDB.Close()

	query, err := connectionDB.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	for query.Next() {
		err := query.Scan(&id, &name, &price, &quantity, &description)
		if err != nil {
			panic(err.Error())
		}

		product := Product{id, name, description, price, quantity}
		products = append(products, product)
	}

	return products
}

func AddProduct(product Product) {
	connectionDB := db.DbConnector()
	defer connectionDB.Close()

	query, err := connectionDB.Prepare("insert into products(name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(product.Name, product.Description, product.Price, product.Quantity)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteProduct(id string) {
	connectionDB := db.DbConnector()
	defer connectionDB.Close()

	query, err := connectionDB.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
}

func RetrieveProduct(productId string) Product {
	connectionDB := db.DbConnector()
	defer connectionDB.Close()

	var name, description string
	var price float64
	var quantity, id int

	query, err := connectionDB.Prepare("select * from products where id=$1;")
	if err != nil {
		panic(err.Error())
	}

	err = query.QueryRow(productId).Scan(&id, &name, &price, &quantity, &description)
	if err != nil {
		panic(err.Error())
	}

	return Product{Id: id, Name: name, Price: price, Quantity: quantity, Description: description}
}

func UpdateProduct(product Product) {
	connectionDB := db.DbConnector()
	defer connectionDB.Close()

	fmt.Println(product)

	query, err := connectionDB.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	res, _ := query.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
	fmt.Println(res.RowsAffected())

}
