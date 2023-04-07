package usecase

import (
	"context"
	"log"
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
	log.Printf("Starting the process of StoreFileContent")

	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("The file content has been processed successfully. Number of entries: %v", len(entries))

	return entries, nil
}

// StoreFileContent stores a new Transaction
func (uc *TransactionUseCase) StoreFileMetadata(ctx context.Context, ctfm *CreateFileMetadata) (string, error) {
	log.Printf("Starting the process of StoreFileMetadata: %v", ctfm)

	hash := common.CalculateSHA256Hash(ctfm.BinaryData)
	fileSize, err := strconv.Atoi(ctfm.FileSize)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		if err, ok := err.(common.EntityConflictError); ok {
			return "", common.EntityConflictError{
				Message: ErrFileMetadataAlreadyExists.Error(),
				Err:     err,
			}
		}
		return "", err
	}

	log.Printf("The file metadata has been processed successfully.	: %v", ctfm)

	return tfm.ID, nil
}

func (uc *TransactionUseCase) GetFileTransactions(ctx context.Context, fileID string) ([]*e.Transaction, error) {
	log.Printf("GetFileTransactions() with fileID %s", fileID)

	transactions, err := uc.TransactionRepo.List(ctx, fileID)
	if err != nil {
		log.Printf("GetFileTransactions().err %v", err)
		return nil, err
	}

	log.Println("GetFileTransactions.success")

	return transactions, nil
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

	if ct.TType == e.AFFILIATE_SALE {
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
