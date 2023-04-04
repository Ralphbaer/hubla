package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/transaction/entity"
)

// TransactionRepository manages transaction repository operations
type TransactionRepository interface {
	List(ctx context.Context, fileID string) ([]*e.Transaction, error)
	Save(ctx context.Context, t *e.Transaction) (*string, error)
}
