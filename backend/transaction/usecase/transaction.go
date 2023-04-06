package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
	r "github.com/Ralphbaer/hubla/backend/transaction/repository"
	"github.com/google/uuid"
)

// TransactionUseCase represents a collection of use cases for Transaction operations
type TransactionUseCase struct {
	FileMetadataRepo    r.FileMetadataRepository
	SellerRepo          r.SellerRepository
	ProductRepo         r.ProductRepository
	TransactionRepo     r.TransactionRepository
	FileTransactionRepo r.FileTransactionRepository
	SellerBalanceRepo   r.SellerBalanceRepository
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileContent(ctx context.Context, binaryData []byte) ([]*e.Transaction, error) {
	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileMetadata(ctx context.Context, ctfm *CreateFileMetadata) (string, error) {
	hash := common.CalculateSHA256Hash(ctfm.BinaryData)

	fm, err := uc.FileMetadataRepo.FindByHash(ctx, hash)
	if fm != nil {
		return "", fmt.Errorf(ErrFileMetadataAlreadyExists.Error(), hash)
	}
	if err != nil {
		if _, ok := err.(common.EntityNotFoundError); !ok {
			return "", err
		}
	}

	fileSize, err := strconv.Atoi(ctfm.FileSize)
	if err != nil {
		return "", err
	}

	tfm := &e.FileMetadata{
		ID:          uuid.NewString(),
		FileSize:    fileSize,
		Disposition: ctfm.Disposition,
		Hash:        hash,
		BinaryData:  ctfm.BinaryData,
	}

	if err := uc.FileMetadataRepo.Save(ctx, tfm); err != nil {
		return "", err
	}

	return tfm.ID, nil
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) GetFileTransactions(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	return uc.TransactionRepo.List(ctx, fileID)
}
