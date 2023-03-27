package entity

import "time"

// Sales represents a collection of identification data about a Zé Delivery Sales,
// including its coordinates represented by the coverageArea and address fields.
// swagger:model Sales
type Sales struct {
	ID                 string
	TType              uint8
	TDate              time.Time
	ProductDescription string
	Amount             int
	Seller             string
}
