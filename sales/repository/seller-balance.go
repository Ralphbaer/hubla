package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SellerBalanceRepository manages sales repository operations
type SellerBalanceRepository interface {
	Find(ctx context.Context, sellerID string) (*e.SellerBalanceView, error)
	Upsert(ctx context.Context, sb *e.SellerBalance) (*float64, error)
}
