package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SellerBalanceRepository manages sales repository operations
type ProductRepository interface {
	Find(ctx context.Context, sellerName string) (*e.Product, error)
	Save(ctx context.Context, p *e.Product) error
}
