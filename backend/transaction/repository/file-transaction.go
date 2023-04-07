package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

//go:generate mockgen -destination=../gen/mocks/file-transaction_repository_mock.go -package=mocks . FileTransactionRepository

// FileMetadataRepository manages transaction repository operations
type FileTransactionRepository interface {
	Find(ctx context.Context, ID string) (*e.FileTransaction, error)
	Save(ctx context.Context, ft *e.FileTransaction) error
}
