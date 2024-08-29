package application

import (
	"context"
)

type StoreInterface interface {
	Store()
}

type Product struct {
	repo StoreInterface
}

func NewProduct(r StoreInterface) *Product {
	return &Product{
		repo: r,
	}
}

func (p *Product) Save(ctx context.Context) error {

	p.repo.Store()

	return nil
}
