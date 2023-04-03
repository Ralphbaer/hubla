package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
)

// CreatorUseCase represents a collection of use cases for sales operations
type SellerUseCase struct {
	SellerBalanceRepo r.SellerBalanceRepository
	SellerRepo        r.SellerRepository
}

// StoreFileContent stores a new Transaction
func (uc *SellerUseCase) GetSellerBalanceByID(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	return uc.SellerBalanceRepo.Find(ctx, sellerID)
}
