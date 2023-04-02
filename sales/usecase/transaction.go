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
	SellerBalanceRepo r.SellerBalanceRepository
}

// StoreFileContent stores a new Sales
func (uc *SalesUseCase) StoreFileContent(ctx context.Context, binaryData []byte) (*TransactionLine, error) {
	// Process the file content
	entries, err := uc.processFileData(ctx, binaryData)
	if err != nil {
		return nil, err
	}

	fmt.Println(entries)
	// uc.SalesRepo.Save(ctx, entries)

	//	if err := uc.orchestrateAndPersist(ctx, entries); err != nil {
	//		return nil, err
	//	}
	// Get the productName

	// if _, err := uc.SalesRepo.Save(ctx, nil); err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (uc *SalesUseCase) processFileData(ctx context.Context, binaryData []byte) ([]TransactionLine, error) {
	var entries []TransactionLine
	sellers := make(map[string]string)

	scanner := bufio.NewScanner(bytes.NewReader(binaryData))
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			log.Printf("Error parsing line: %v", err)
			continue
		}
		// check if the seller already exists
		if _, ok := sellers[entry.SellerName]; !ok {
			seller := &e.Seller{
				ID:         uuid.NewString(),
				Name:       entry.SellerName,
				SellerType: e.TransactionTypeToSellerTypeMap[entry.TType],
			}
			if _, err := uc.SellerRepo.Save(ctx, seller); err != nil {
				return nil, err
			}
			sellers[entry.SellerName] = seller.ID // joga no mapa
		}
		t := &e.Transaction{
			ID:          uuid.NewString(),
			TType:       entry.TType,
			TDate:       entry.TDate,
			ProductName: entry.ProductName,
			Amount:      entry.Amount,
			SellerID:    sellers[entry.SellerName],
		}
		if _, err := uc.TransactionRepo.Save(ctx, t); err != nil {
			return nil, err
		}

		sb := &e.SellerBalance{
			ID:        uuid.NewString(),
			SellerID:  t.SellerID,
			UpdatedAt: time.Now(),
			Balance:   e.TransactionTypeToOperationMap[entry.TType](entry.Amount),
		}
		newBalance, err := uc.SellerBalanceRepo.Upsert(ctx, sb)
		if err != nil {
			return nil, err
		}
		fmt.Println(newBalance)
		// Check if the context is cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return entries, nil
}

func parseLine(line string) (TransactionLine, error) {
	var entry TransactionLine

	if len(line) < 70 {
		return entry, fmt.Errorf("invalid line format")
	}

	entry.ID = uuid.New().String()

	ttype, err := strconv.ParseUint(line[:1], 10, 8)
	if err != nil {
		return entry, fmt.Errorf("error parsing code: %v", err)
	}
	entry.TType = e.TransactionTypeMap[uint8(ttype)]

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])
	if err != nil {
		return entry, fmt.Errorf("error parsing date: %v", err)
	}
	entry.TDate = date

	entry.ProductName = strings.TrimSpace(line[26:50])

	value, err := decimal.NewFromString(strings.TrimSpace(line[50:66]))
	if err != nil {
		return entry, fmt.Errorf("error parsing amount: %v", err)
	}
	entry.Amount = value

	entry.SellerName = strings.TrimSpace(line[66:])

	return entry, nil
}
