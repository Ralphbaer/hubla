package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
)

// TransactionRepository manages transaction repository operations
type UserRepository interface {
	FindByEmail(ctx context.Context, string string) (*e.User, error)
	FindByID(ctx context.Context, ID string) (*e.User, error)
}
