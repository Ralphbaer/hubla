package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Seller represents a seller (Creator or Affiliate) who offers products for sale.
type SellerBalance struct {
	ID        string
	SellerID  string
	Balance   decimal.Decimal
	UpdatedAt time.Time
	CreatedAt time.Time
}
