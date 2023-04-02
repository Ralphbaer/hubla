package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SalesRepository manages sales repository operations
type SalesRepository interface {
	Save(ctx context.Context, t *e.Transaction) (*e.Transaction, error)
}
