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
	TType     TransactionTypeEnum `json:"type"`
	TDate     time.Time           `json:"date"`
	ProductID string              `json:"product_id"`
	Amount    decimal.Decimal     `json:"amount"`
	SellerID  string              `json:"seller_id"`
	CreatedAt time.Time           `json:"created_at"`
}

var TransactionTypeToSellerTypeMap = map[TransactionTypeEnum]SellerTypeEnum{
	CREATOR_SALE:        CREATOR,
	AFFILIATE_SALE:      AFFILIATE,
	COMMISSION_PAID:     CREATOR,
	COMMISSION_RECEIVED: AFFILIATE,
}

type TransactionTypeEnum uint8

const (
	CREATOR_SALE        TransactionTypeEnum = 1
	AFFILIATE_SALE      TransactionTypeEnum = 2
	COMMISSION_PAID     TransactionTypeEnum = 3
	COMMISSION_RECEIVED TransactionTypeEnum = 4
)

var TransactionTypeMap = map[uint8]TransactionTypeEnum{
	1: CREATOR_SALE,
	2: AFFILIATE_SALE,
	3: COMMISSION_PAID,
	4: COMMISSION_RECEIVED,
}

var TransactionTypeMapString = map[TransactionTypeEnum]string{
	CREATOR_SALE:        "CREATOR_SALE",
	AFFILIATE_SALE:      "AFFILIATE_SALE",
	COMMISSION_PAID:     "COMMISSION_PAID",
	COMMISSION_RECEIVED: "COMMISSION_RECEIVED",
}
