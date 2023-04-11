package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Transaction represents a collection of identification data about a Hubla Sales,
// including its coordinates represented by the coverageArea and address fields.
// swagger:model Transaction
type Transaction struct {
	ID        string              `json:"id"`
	TType     TransactionTypeEnum `json:"t_type"`
	TDate     time.Time           `json:"t_date"`
	ProductID string              `json:"product_id"`
	Amount    decimal.Decimal     `json:"amount"`
	SellerID  string              `json:"seller_id"`
	CreatedAt time.Time           `json:"created_at"`
}

// TransactionTypeEnum is a custom type that represents the type of transaction.
type TransactionTypeEnum uint8

// These constants represent the different types of transactions.
const (
	CreatorSale        TransactionTypeEnum = 1 // The transaction type for creator sales.
	AffiliateSale      TransactionTypeEnum = 2 // The transaction type for affiliate sales.
	CommissionPaid     TransactionTypeEnum = 3 // The transaction type for commission paid.
	CommissionReceived TransactionTypeEnum = 4 // The transaction type for commission received.
)

// TransactionTypeToSellerTypeMap maps TransactionTypeEnum to SellerTypeEnum.
var TransactionTypeToSellerTypeMap = map[TransactionTypeEnum]SellerTypeEnum{
	CreatorSale:        CREATOR,
	AffiliateSale:      AFFILIATE,
	CommissionPaid:     CREATOR,
	CommissionReceived: AFFILIATE,
}

// TransactionTypeMap maps uint8 to TransactionTypeEnum.
var TransactionTypeMap = map[uint8]TransactionTypeEnum{
	1: CreatorSale,
	2: AffiliateSale,
	3: CommissionPaid,
	4: CommissionReceived,
}

// SellerTypeMapString maps SellerTypeEnum to string.
var SellerTypeMapString = map[SellerTypeEnum]string{
	CREATOR:   "CREATOR",
	AFFILIATE: "AFFILIATE",
}

// TransactionTypeMapString maps TransactionTypeEnum to string.
var TransactionTypeMapString = map[TransactionTypeEnum]string{
	CreatorSale:        "CREATOR_SALE",
	AffiliateSale:      "AFFILIATE_SALE",
	CommissionPaid:     "COMMISSION_PAID",
	CommissionReceived: "COMMISSION_RECEIVED",
}

// TransactionTypeMapEnum is a map that allows for getting a transaction type enum value
// from its string representation.
var TransactionTypeMapEnum = map[string]TransactionTypeEnum{
	"CREATOR_SALE":        CreatorSale,
	"AFFILIATE_SALE":      AffiliateSale,
	"COMMISSION_PAID":     CommissionPaid,
	"COMMISSION_RECEIVED": CommissionReceived,
}
