package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/transaction/entity"
	r "github.com/Ralphbaer/hubla/transaction/repository"
)

// CreatorUseCase represents a collection of use cases for transaction operations
type SellerUseCase struct {
	SellerBalanceRepo r.SellerBalanceRepository
	SellerRepo        r.SellerRepository
}

// StoreFileContent stores a new Transaction
func (uc *SellerUseCase) GetSellerBalanceByID(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	return uc.SellerBalanceRepo.Find(ctx, sellerID)
}
