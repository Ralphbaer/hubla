package entity

import "time"

// Seller represents a seller (Creator or Affiliate) who offers products for sale.
type Seller struct {
	ID         string
	Name       string
	SellerType SellerTypeEnum
	CreatedAt  time.Time
}
