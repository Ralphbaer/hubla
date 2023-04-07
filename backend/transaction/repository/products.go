package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/product_repository_mock.go -package=mocks . ProductRepository

// ProductRepository manages transaction repository operations
type ProductRepository interface {
	FindByProductName(ctx context.Context, productName string) (*e.Product, error)
	Save(ctx context.Context, p *e.Product) error
}
