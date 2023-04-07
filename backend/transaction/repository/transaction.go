package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/transaction_repository_mock.go -package=mocks . TransactionRepository

// TransactionRepository manages transaction repository operations
type TransactionRepository interface {
	ListTransactionsByFileID(ctx context.Context, fileID string) ([]*e.Transaction, error)
	Save(ctx context.Context, t *e.Transaction) error
}
