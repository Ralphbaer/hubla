package usecase

import (
	"context"
	"strconv"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	r "github.com/Ralphbaer/hubla/backend/transaction/repository"
	"github.com/google/uuid"
)

// TransactionUseCase represents a collection of use cases for Transaction operations
type TransactionUseCase struct {
	FileMetadataRepo  r.FileMetadataRepository
	TransactionRepo   r.TransactionRepository
	SellerRepo        r.SellerRepository
	ProductRepo       r.ProductRepository
	SellerBalanceRepo r.SellerBalanceRepository
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileContent(ctx context.Context, hash string, binaryData []byte) ([]*e.Transaction, error) {
	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return entries, err
	}

	return entries, nil
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileMetadata(ctx context.Context, ctfm *CreateFileMetadata) (string, error) {
	fileSize, err := strconv.Atoi(ctfm.FileSize)
	if err != nil {
		return "", err
	}

	tfm := &e.FileMetadata{
		ID:          uuid.NewString(),
		FileSize:    fileSize,
		Disposition: ctfm.Disposition,
		Hash:        calculateSHA256Hash(ctfm.BinaryData),
		BinaryData:  ctfm.BinaryData,
	}

	if err := uc.FileMetadataRepo.Save(ctx, tfm); err != nil {
		return "", err
	}

	return tfm.ID, nil
}
