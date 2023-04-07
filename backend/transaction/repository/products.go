package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// SellerBalanceRepository manages transaction repository operations
type ProductRepository interface {
	FindByProductName(ctx context.Context, productName string) (*e.Product, error)
	Save(ctx context.Context, p *e.Product) error
}
