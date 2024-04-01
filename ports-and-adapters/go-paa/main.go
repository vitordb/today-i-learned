package main

import (
	"go-paa/internal/product/infra/rest"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/product", rest.AddProduct)

	log.Fatal(http.ListenAndServe(":8099", nil))
}
