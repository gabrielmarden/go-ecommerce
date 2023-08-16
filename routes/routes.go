package routes

import (
	"net/http"

	"io.nedram/lolja/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.HomePageHandler)
	http.HandleFunc("/new", controllers.NewHandler)
	http.HandleFunc("/insert", controllers.InsertHandler)
	http.HandleFunc("/delete", controllers.DeleteHandler)
	http.HandleFunc("/edit", controllers.EditHandler)
	http.HandleFunc("/update", controllers.UpdateHandler)
	http.HandleFunc("/api/products", controllers.ProductsHandler)
}
