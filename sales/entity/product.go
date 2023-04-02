package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Product represents a product that can be created and sold by creators or affiliates.
type Product struct {
	ID        string
	Name      string
	CreatorID string
	CreatedAt time.Time
}

var SellerTypeToOperationMap = map[TransactionTypeEnum]func(amount decimal.Decimal) decimal.Decimal{
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
