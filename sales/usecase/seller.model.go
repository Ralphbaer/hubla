package usecase

import (
	e "github.com/Ralphbaer/hubla/sales/entity"
)

type CreateSeller struct {
	SellerName string
	SellerType e.SellerTypeEnum
}
