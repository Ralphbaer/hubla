package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/sales/entity"
)

// SalesRepository manages sales repository operations
type SalesRepository interface {
	Save(ctx context.Context, s []e.Sales) (*string, error)
}
