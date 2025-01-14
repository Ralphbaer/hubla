package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/common/hlog"
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
	hlog.NewLoggerFromContext(ctx).Infof("Starting the process of StoreFileContent")

	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// StoreFileMetadata stores the given file metadata information in the database using the FileMetadataRepo.
// It returns the file ID if successfully saved, or an error if the query fails.
func (uc *TransactionUseCase) StoreFileMetadata(ctx context.Context, ctfm *CreateFileMetadata) (*FileID, error) {
	hlog.NewLoggerFromContext(ctx).Infof("Starting the process of StoreFileMetadata: %v", ctfm)

	hash := common.CalculateSHA256Hash(ctfm.BinaryData)
	fileSize, err := strconv.Atoi(ctfm.FileSize)
	if err != nil {
		return nil, err
	}

	tfm := &e.FileMetadata{
		ID:          ctfm.ID,
		FileSize:    fileSize,
		Disposition: ctfm.Disposition,
		Hash:        hash,
		BinaryData:  ctfm.BinaryData,
	}

	if err := uc.FileMetadataRepo.Save(ctx, tfm); err != nil {
		if err, ok := err.(common.EntityConflictError); ok {
			return nil, common.EntityConflictError{
				ErrCode: "ErrFileMetadataAlreadyExists",
				Message: fmt.Sprintf(ErrFileMetadataAlreadyExists.Error(), tfm.ID),
				Err:     err,
			}
		}
		return nil, err
	}

	return &FileID{ID: tfm.ID}, nil
}

// GetFileTransactions retrieves all transactions associated with the provided file ID.
// Returns a slice of Transaction pointers and an error if there's any issue.
func (uc *TransactionUseCase) GetFileTransactions(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	hlog.NewLoggerFromContext(ctx).Infof("Retrieving file transaction by fileID %s", fileID)

	transactions, err := uc.TransactionRepo.ListTransactionsByFileID(ctx, fileID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// ListFileTransactions retrieves all file transactions.
// Returns a slice of Transaction pointers and an error if there's any issue.
func (uc *TransactionUseCase) ListFileTransactions(ctx context.Context) ([]*e.FileTransaction, error) {
	hlog.NewLoggerFromContext(ctx).Infof("Retrieving a list of file transactions")

	ft, err := uc.FileTransactionRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return ft, nil
}

func (uc *TransactionUseCase) handleTransaction(ctx context.Context, ct *CreateTransaction, sellerNameToID map[string]string, productIDToName map[string]*e.Product) (*e.Transaction, error) {
	sellerID, err := uc.findOrCreateSeller(ctx, ct.SellerName, ct.TType, sellerNameToID)
	if err != nil {
		return nil, err
	}

	product, err := uc.findOrCreateProduct(ctx, ct.ProductName, productIDToName, sellerID)
	if err != nil {
		return nil, err
	}

	t, err := uc.saveTransaction(ctx, ct, product.ID, sellerID)
	if err != nil {
		return nil, err
	}

	if ct.TType == e.AffiliateSale {
		sellerID = product.CreatorID
	}

	err = uc.updateSellerBalance(ctx, ct, sellerID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (uc *TransactionUseCase) saveTransaction(ctx context.Context, ct *CreateTransaction, productID, sellerID string) (*e.Transaction, error) {
	transaction := &e.Transaction{
		ID:        uuid.NewString(),
		TType:     ct.TType,
		TDate:     ct.TDate,
		ProductID: productID,
		Amount:    ct.Amount,
		SellerID:  sellerID,
	}

	if err := uc.TransactionRepo.Save(ctx, transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
