package usecase

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	e "github.com/Ralphbaer/hubla/sales/entity"
	r "github.com/Ralphbaer/hubla/sales/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SalesUseCase represents a collection of use cases for sales operations
type SalesUseCase struct {
	TransactionRepo   r.TransactionRepository
	SellerRepo        r.SellerRepository
	ProductRepo       r.ProductRepository
	SellerBalanceRepo r.SellerBalanceRepository
}

// StoreFileContent stores a new Sales
func (uc *SalesUseCase) StoreFileContent(ctx context.Context, binaryData []byte) (*TransactionLine, error) {
	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return nil, err
	}

	fmt.Println(entries)

	return nil, nil
}

func (uc *SalesUseCase) processFileData(ctx context.Context, binaryData []byte) ([]*e.Transaction, error) {
	var transactions []*e.Transaction
	sellers := make(map[string]*e.Seller)
	products := make(map[string]*e.Product)

	scanner := bufio.NewScanner(bytes.NewReader(binaryData))
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			log.Printf("Error parsing line: %v", err)
			continue
		}
		if err := uc.handleTransaction(ctx, entry, sellers, products); err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return transactions, nil
}

func parseLine(line string) (*TransactionLine, error) {
	if len(line) < 70 {
		return nil, fmt.Errorf("invalid line format")
	}

	ID := uuid.New().String()
	tTypeCode, err := strconv.Atoi(string(line[0]))
	if err != nil {
		return nil, fmt.Errorf("error parsing transaction type: %v", err)
	}
	tType, ok := e.TransactionTypeMap[uint8(tTypeCode)]
	if !ok {
		return nil, fmt.Errorf("invalid transaction type")
	}

	tDate, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])
	if err != nil {
		return nil, fmt.Errorf("error parsing date: %v", err)
	}

	productName := strings.TrimSpace(line[26:50])

	amount, err := decimal.NewFromString(strings.TrimSpace(line[50:66]))
	if err != nil {
		return nil, fmt.Errorf("error parsing amount: %v", err)
	}

	sellerName := strings.TrimSpace(line[66:])
	if len(sellerName) == 0 {
		return nil, fmt.Errorf("invalid seller name")
	}

	return &TransactionLine{
		ID:          ID,
		TType:       tType,
		TDate:       tDate,
		ProductName: productName,
		Amount:      amount,
		SellerName:  sellerName,
	}, nil
}

func (uc *SalesUseCase) handleTransaction(ctx context.Context, entry *TransactionLine, sellers map[string]*e.Seller, products map[string]*e.Product) error {
	seller, err := uc.findOrCreateSeller(ctx, entry, sellers)
	if err != nil {
		return err
	}

	product, err := uc.findOrCreateProduct(ctx, entry, products, seller.ID)
	if err != nil {
		return err
	}

	err = uc.saveTransaction(ctx, entry, product.ID, seller.ID)
	if err != nil {
		return err
	}

	err = uc.updateSellerBalance(ctx, entry, seller.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SalesUseCase) findOrCreateSeller(ctx context.Context, entry *TransactionLine, sellers map[string]*e.Seller) (*e.Seller, error) {
	seller, found := sellers[entry.SellerName]
	if found {
		return seller, nil
	}

	seller, err := uc.SellerRepo.Find(ctx, entry.SellerName)
	if err != nil {
		return nil, err
	}

	if seller == nil {
		seller = &e.Seller{
			ID:         uuid.NewString(),
			Name:       entry.SellerName,
			SellerType: e.TransactionTypeToSellerTypeMap[entry.TType],
		}
		_, err = uc.SellerRepo.Save(ctx, seller)
		if err != nil {
			return nil, err
		}
	}

	sellers[entry.SellerName] = seller

	return seller, nil
}

func (uc *SalesUseCase) findOrCreateProduct(ctx context.Context, entry *TransactionLine, products map[string]*e.Product, sellerID string) (*e.Product, error) {
	product, found := products[entry.ProductName]
	if found {
		return product, nil
	}

	product, err := uc.ProductRepo.Find(ctx, entry.ProductName)
	if err != nil {
		return nil, err
	}

	if product == nil {
		product = &e.Product{
			ID:        uuid.NewString(),
			Name:      entry.ProductName,
			CreatorID: sellerID,
		}
		err = uc.ProductRepo.Save(ctx, product)
		if err != nil {
			return nil, err
		}
	}

	products[entry.ProductName] = product

	return product, nil
}

func (uc *SalesUseCase) saveTransaction(ctx context.Context, entry *TransactionLine, productID, sellerID string) error {
	transaction := &e.Transaction{
		ID:        uuid.NewString(),
		TType:     entry.TType,
		TDate:     entry.TDate,
		ProductID: productID,
		Amount:    entry.Amount,
		SellerID:  sellerID,
	}

	if _, err := uc.TransactionRepo.Save(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (uc *SalesUseCase) updateSellerBalance(ctx context.Context, entry *TransactionLine, sellerID string) error {
	sellerBalance := &e.SellerBalance{
		ID:        uuid.NewString(),
		SellerID:  sellerID,
		Balance:   e.TransactionTypeToOperationMap[entry.TType](entry.Amount),
		UpdatedAt: time.Now(),
	}

	updatedBalance, err := uc.SellerBalanceRepo.Upsert(ctx, sellerBalance)
	if err != nil {
		return err
	}

	fmt.Println(updatedBalance)

	return nil
}
