package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
	"github.com/google/uuid"
)

// CreatorUseCase represents a collection of use cases for sales operations
type SellerUseCase struct {
	SellerRepo r.SellerRepository
}

// StoreFileContent stores a new Sales
func (uc *SellerUseCase) GetCreatorAmount(ctx context.Context, ID string) (*e.Transaction, error) {
	return nil, nil
}

// StoreFileContent stores a new Sales
func (uc *SellerUseCase) CreateSeller(ctx context.Context, cs *CreateSeller) (string, error) {
	s := &e.Seller{
		ID:         uuid.NewString(),
		SellerType: cs.SellerType,
		Name:       cs.SellerName,
	}

	ID, err := uc.SellerRepo.Save(ctx, s)
	if err != nil {
		return "", err
	}

	return ID, nil
}
