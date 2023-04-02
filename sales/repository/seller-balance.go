package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
	"github.com/shopspring/decimal"
)

// SellerBalanceRepository manages sales repository operations
type SellerBalanceRepository interface {
	Upsert(ctx context.Context, sb *e.SellerBalance) (*decimal.Decimal, error)
}
