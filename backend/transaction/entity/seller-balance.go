package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// SellerBalance is a struct representing a seller's balance,
// including the balance ID, seller ID, balance amount, and creation and update times.
type SellerBalance struct {
	ID        string
	SellerID  string
	Balance   decimal.Decimal
	UpdatedAt time.Time
	CreatedAt time.Time
}

// TransactionTypeToOperationMap is a mapping between TransactionTypeEnum values and
// corresponding functions that manipulate the transaction amount.
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
