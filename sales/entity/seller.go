package entity

import "time"

// Seller represents a seller (Creator or Affiliate) who offers products for sale.
type Seller struct {
	ID         string
	Name       string
	SellerType SellerTypeEnum
	CreatedAt  time.Time
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
