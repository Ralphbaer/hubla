package usecase

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/Ralphbaer/hubla/backend/common"
	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// CreateFileTransactions creates a new file transaction for each transaction in the provided slice.
// Returns an error if there's any issue.
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

func (uc *TransactionUseCase) processFileData(ctx context.Context, binaryData []byte) ([]*e.Transaction, error) {
	var transactions []*e.Transaction
	sellerNameToID := make(map[string]string)
	product := make(map[string]*e.Product)

	lines, err := uc.readFileData(binaryData)
	if err != nil {
		return nil, err
	}

	for lineNumber, line := range lines {
		transaction, err := uc.processLine(ctx, line, lineNumber, sellerNameToID, product)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (uc *TransactionUseCase) readFileData(binaryData []byte) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(binaryData))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, common.UnprocessableOperationError{
			Err:     err,
			Message: err.Error(),
		}
	}

	return lines, nil
}

func (uc *TransactionUseCase) processLine(ctx context.Context, line string, lineNumber int, seller map[string]string, product map[string]*e.Product) (*e.Transaction, error) {
	entry, err := parseLine(line)
	if err != nil {
		return nil, common.UnprocessableOperationError{
			Message: fmt.Sprintf(ErrParsingParsingLine.Error(), err),
			Err:     err,
		}
	}

	transaction, err := uc.handleTransaction(ctx, entry, seller, product)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func parseLine(line string) (*CreateTransaction, error) {
	if len(line) < 70 {
		return nil, ErrInvalidLineFormat
	}

	ID := uuid.New().String()
	tTypeCode, err := strconv.Atoi(string(line[0]))
	if err != nil {
		return nil, ErrInvalidTransactionType
	}
	tType, ok := e.TransactionTypeMap[uint8(tTypeCode)]
	if !ok {
		return nil, ErrInvalidTransactionType
	}

	tDate, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])
	if err != nil {
		return nil, ErrInvalidDate
	}

	productName := strings.TrimSpace(line[26:50])

	amount, err := decimal.NewFromString(strings.TrimSpace(line[50:66]))
	if err != nil {
		return nil, ErrInvalidAmount
	}

	sellerName := strings.TrimSpace(line[66:])
	if len(sellerName) == 0 {
		return nil, ErrInvalidSellerName
	}

	return &CreateTransaction{
		ID:          ID,
		TType:       tType,
		TDate:       tDate,
		ProductName: productName,
		Amount:      amount,
		SellerName:  sellerName,
	}, nil
}
