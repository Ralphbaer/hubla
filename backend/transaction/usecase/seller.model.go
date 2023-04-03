package usecase

import (
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

type CreateSeller struct {
	SellerName string
	SellerType e.SellerTypeEnum
}
