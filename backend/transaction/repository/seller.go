package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/seller_repository_mock.go -package=mocks . SellerRepository

// SellerRepository manages transaction repository operations
type SellerRepository interface {
	FindBySellerName(ctx context.Context, sellerName string) (*e.Seller, error)
	Save(ctx context.Context, s *e.Seller) error
}
