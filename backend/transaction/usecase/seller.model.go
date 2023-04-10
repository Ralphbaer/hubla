package usecase

import (
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// CreateSeller is a struct that holds information about a new Seller to be created.
type CreateSeller struct {
	SellerName string
	SellerType e.SellerTypeEnum
}
