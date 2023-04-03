package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
)

// TransactionUseCase represents a collection of use cases for Transaction operations
type TransactionUseCase struct {
	TransactionRepo   r.TransactionRepository
	SellerRepo        r.SellerRepository
	ProductRepo       r.ProductRepository
	SellerBalanceRepo r.SellerBalanceRepository
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileContent(ctx context.Context, binaryData []byte) ([]*e.Transaction, error) {
	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return entries, err
	}

	return entries, nil
}
