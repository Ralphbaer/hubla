package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/transaction/entity"
)

// SellerRepository manages transaction repository operations
type SellerRepository interface {
	Find(ctx context.Context, sellerName string) (*e.Seller, error)
	Save(ctx context.Context, s *e.Seller) (string, error)
}
