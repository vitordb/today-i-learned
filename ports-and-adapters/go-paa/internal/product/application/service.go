package product

import "context"

type Product struct {
	repo ProductRepository
}

func (p *Product) Save(ctx context.Context) error {

	return nil
}
