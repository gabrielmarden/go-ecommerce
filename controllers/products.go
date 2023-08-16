package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"io.nedram/lolja/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	products := models.RetrieveProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		product := getProductFromRequest(r)

		models.AddProduct(product)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.RetrieveProduct(id)

	templates.ExecuteTemplate(w, "Edit", product)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		product := getProductFromRequest(r)

		models.UpdateProduct(product)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

func getProductFromRequest(r *http.Request) models.Product {
	name := r.FormValue("nome")
	description := r.FormValue("descricao")
	price := r.FormValue("preco")
	quantity := r.FormValue("quantidade")
	id := r.FormValue("id")

	priceConverted, err := strconv.ParseFloat(price, 64)
	if err != nil {
		panic(err.Error())
	}

	quantityConverted, err := strconv.Atoi(quantity)
	if err != nil {
		panic(err.Error())
	}

	idConverted, err := strconv.Atoi(id)
	if err != nil {
		idConverted = -1
	}

	product := models.Product{Id: idConverted, Name: name, Description: description, Price: priceConverted, Quantity: quantityConverted}

	return product
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := models.RetrieveProducts()

	json.NewEncoder(w).Encode(products)
}
