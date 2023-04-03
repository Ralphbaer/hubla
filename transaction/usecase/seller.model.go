package usecase

import (
	e "github.com/Ralphbaer/hubla/transaction/entity"
)

type CreateSeller struct {
	SellerName string
	SellerType e.SellerTypeEnum
}
