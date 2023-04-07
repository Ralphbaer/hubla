package usecase

import (
	"context"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	r "github.com/Ralphbaer/hubla/backend/transaction/repository"
)

// CreatorUseCase represents a collection of use cases for transaction operations
type SellerUseCase struct {
	SellerBalanceRepo r.SellerBalanceRepository
	SellerRepo        r.SellerRepository
}

// StoreFileContent stores a new Transaction
func (uc *SellerUseCase) GetSellerBalanceByID(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	hlog.NewLoggerFromContext(ctx).Infof("Retrieving Health Pulse for profile ID %s", sellerID)

	return uc.SellerBalanceRepo.Find(ctx, sellerID)
}
