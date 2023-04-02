package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// TransactionRepository manages sales repository operations
type TransactionRepository interface {
	Save(ctx context.Context, t *e.Transaction) (*e.Transaction, error)
}
