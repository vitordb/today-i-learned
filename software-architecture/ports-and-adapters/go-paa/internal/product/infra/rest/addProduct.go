package rest

import (
	"go-paa/internal/product/application"
	"go-paa/internal/product/infra/mysql"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	db := mysql.NewProductRepository()
	product := application.NewProduct(db)

	product.Save(r.Context())

}
