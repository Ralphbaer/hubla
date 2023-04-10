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

// SellerTypeToOperationMap is a mapping between TransactionTypeEnum values and
// corresponding functions that manipulate the transaction amount.
var SellerTypeToOperationMap = map[TransactionTypeEnum]func(amount decimal.Decimal) decimal.Decimal{
	CreatorSale: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
	AffiliateSale: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
	CommissionPaid: func(amount decimal.Decimal) decimal.Decimal {
		return amount.Neg()
	},
	CommissionReceived: func(amount decimal.Decimal) decimal.Decimal {
		return amount
	},
}
