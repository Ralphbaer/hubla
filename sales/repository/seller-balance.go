package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SellerBalanceRepository manages sales repository operations
type SellerBalanceRepository interface {
	Upsert(ctx context.Context, sb *e.SellerBalance) (*float64, error)
}
