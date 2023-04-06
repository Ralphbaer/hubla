package usecase

import (
	"context"

	"github.com/google/uuid"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

func (uc *TransactionUseCase) CreateFileTransactions(ctx context.Context, fileID string, transactions []*e.Transaction) error {
	for _, v := range transactions {
		fileTransaction := &e.FileTransaction{
			ID:            uuid.NewString(),
			FileID:        fileID,
			TransactionID: v.ID,
		}

		if err := uc.FileTransactionRepo.Save(ctx, fileTransaction); err != nil {
			return err
		}
	}

	return nil
}
