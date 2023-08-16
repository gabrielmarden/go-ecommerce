package main

import (
	"net/http"

	"io.nedram/lolja/routes"
)

func main() {
	routes.Routes()

	http.ListenAndServe(":8000", nil)
}
