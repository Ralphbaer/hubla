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

var TransactionTypeToOperationMap = map[TransactionTypeEnum]func(amount decimal.Decimal) decimal.Decimal{
	CREATOR_SALE: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
	AFFILIATE_SALE: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
	COMMISSION_PAID: func(amount decimal.Decimal) decimal.Decimal {
		return amount.Neg()
	},
	COMMISSION_RECEIVED: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
}
