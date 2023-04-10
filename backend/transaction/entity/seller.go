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

// SellerBalanceView is a struct representing a seller's balance,
// including the seller ID, name, balance, and the last update time.
type SellerBalanceView struct {
	SellerID      string          `json:"seller_id"`
	SellerName    string          `json:"seller_name"`
	SellerBalance decimal.Decimal `json:"balance"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

// SellerTypeEnum is a custom type representing different types of sellers.
type SellerTypeEnum uint8

// Constants representing different types of sellers.
const (
	CREATOR   SellerTypeEnum = 1
	AFFILIATE SellerTypeEnum = 2
)

// SellerTypeMap maps SellerTypeEnum values to their corresponding string representations.
var SellerTypeMap = map[SellerTypeEnum]string{
	CREATOR:   "CREATOR",
	AFFILIATE: "AFFILIATE",
}

// SellerTypeFromString maps string representations of seller types to their corresponding SellerTypeEnum values.
var SellerTypeFromString = map[string]SellerTypeEnum{
	"CREATOR":   CREATOR,
	"AFFILIATE": AFFILIATE,
}
