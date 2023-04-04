package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

func calculateSHA256Hash(binaryData []byte) string {
	hasher := sha256.New()
	hasher.Write(binaryData)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

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
