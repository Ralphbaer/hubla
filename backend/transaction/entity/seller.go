package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Seller represents a seller (Creator or Affiliate) who offers products for sale.
type Seller struct {
	ID         string
	Name       string
	SellerType SellerTypeEnum
	CreatedAt  time.Time
}

// SellerBalanceView represents a seller (Creator or Affiliate) who offers products for sale.
type SellerBalanceView struct {
	SellerID      string          `json:"seller_id"`
	SellerName    string          `json:"seller_name"`
	SellerBalance decimal.Decimal `json:"balance"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type SellerTypeEnum uint8

const (
	CREATOR   SellerTypeEnum = 1
	AFFILIATE SellerTypeEnum = 2
)

var SellerTypeMap = map[SellerTypeEnum]string{
	CREATOR:   "CREATOR",
	AFFILIATE: "AFFILIATE",
}

var SellerTypeFromString = map[string]SellerTypeEnum{
	"CREATOR":   CREATOR,
	"AFFILIATE": AFFILIATE,
}
