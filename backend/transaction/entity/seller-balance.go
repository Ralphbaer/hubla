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
