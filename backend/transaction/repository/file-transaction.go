package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/file-transaction_repository_mock.go -package=mocks . FileTransactionRepository

// FileTransactionRepository defines the methods for storing and retrieving file transactions.
type FileTransactionRepository interface {
	Find(ctx context.Context, ID string) (*e.FileTransaction, error)
	Save(ctx context.Context, ft *e.FileTransaction) error
}
