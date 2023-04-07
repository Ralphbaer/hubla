package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/seller_balance_repository_mock.go -package=mocks . SellerBalanceRepository

// SellerBalanceRepository manages transaction repository operations
type SellerBalanceRepository interface {
	Find(ctx context.Context, sellerID string) (*e.SellerBalanceView, error)
	Upsert(ctx context.Context, sb *e.SellerBalance) (*float64, error)
}
