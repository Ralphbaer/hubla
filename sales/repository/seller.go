package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SellerRepository manages sales repository operations
type SellerRepository interface {
	Find(ctx context.Context, sellerName string, productName string, sellerType e.SellerTypeEnum) (*e.Seller, error)
	Save(ctx context.Context, s *e.Seller) (string, error)
}
