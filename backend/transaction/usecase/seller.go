package usecase

import (
	"context"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	r "github.com/Ralphbaer/hubla/backend/transaction/repository"
)

// SellerUseCase is a use case struct that contains the necessary repositories for performing business logic related to Sellers.
type SellerUseCase struct {
	SellerBalanceRepo r.SellerBalanceRepository
	SellerRepo        r.SellerRepository
}

// GetSellerBalanceByID retrieves the SellerBalanceView entity for the given seller ID using the SellerBalanceRepo.
// It returns a SellerBalanceView object if found, or an error if the query fails.
func (uc *SellerUseCase) GetSellerBalanceByID(ctx context.Context, sellerID string) (*e.SellerBalanceView, error) {
	hlog.NewLoggerFromContext(ctx).Infof("Retrieving Health Pulse for profile ID %s", sellerID)

	return uc.SellerBalanceRepo.Find(ctx, sellerID)
}
