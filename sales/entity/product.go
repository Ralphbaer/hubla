package entity

import "time"

// Product represents a product that can be created and sold by creators or affiliates.
type Product struct {
	ID        string
	Name      string
	CreatorID string
	CreatedAt time.Time
}
