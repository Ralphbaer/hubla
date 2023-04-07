package usecase

import (
	"context"
	"log"

	"github.com/Ralphbaer/hubla/backend/transaction/entity"
	"github.com/google/uuid"
)

func (uc *TransactionUseCase) findOrCreateSeller(ctx context.Context, name string, tType entity.TransactionTypeEnum, sellerNameToID map[string]string) (string, error) {
	sellerID, err := uc.findSellerID(ctx, name, sellerNameToID)
	if err != nil {
		return "", err
	}
	if sellerID == "" {
		ID, err := uc.createSeller(ctx, name, tType)
		if err != nil {
			return "", err
		}
		sellerNameToID[name] = ID
	}

	return sellerNameToID[name], nil
}

func (uc *TransactionUseCase) findSellerID(ctx context.Context, name string, sellerNameToID map[string]string) (string, error) {
	if sellerID, found := sellerNameToID[name]; found {
		return sellerID, nil
	}

	seller, err := uc.SellerRepo.FindBySellerName(ctx, name)
	if err != nil {
		return "", err
	}
	if seller == nil {
		return "", nil
	}

	return seller.ID, nil
}

func (uc *TransactionUseCase) createSeller(ctx context.Context, sellerName string, tType entity.TransactionTypeEnum) (string, error) {
	seller := &entity.Seller{
		ID:         uuid.NewString(),
		Name:       sellerName,
		SellerType: entity.TransactionTypeToSellerTypeMap[tType],
	}
	if err := uc.SellerRepo.Save(ctx, seller); err != nil {
		return "", err
	}

	log.Printf("Created seller with name %s and ID %s\n", seller.Name, seller.ID)

	return seller.ID, nil
}

func (uc *TransactionUseCase) updateSellerBalance(ctx context.Context, ct *CreateTransaction, sellerID string) error {
	sellerBalance := &entity.SellerBalance{
		ID:       uuid.NewString(),
		SellerID: sellerID,
		Balance:  entity.TransactionTypeToOperationMap[ct.TType](ct.Amount),
	}

	updatedBalance, err := uc.SellerBalanceRepo.Upsert(ctx, sellerBalance)
	if err != nil {
		return ErrUpsertingSellerBalance
	}
	log.Printf("Balance for seller %s updated by %s to %v\n", sellerBalance.SellerID, sellerBalance.Balance.String(), updatedBalance)

	return nil
}
