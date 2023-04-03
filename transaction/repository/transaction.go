package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/transaction/entity"
)

// TransactionRepository manages transaction repository operations
type TransactionRepository interface {
	Save(ctx context.Context, t *e.Transaction) (*e.Transaction, error)
}
