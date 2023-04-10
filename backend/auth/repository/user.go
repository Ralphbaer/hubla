package repository

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
)

// UserRepository is an interface for retrieving user data from a data store.
type UserRepository interface {
	FindByEmail(ctx context.Context, string string) (*e.User, error)
	FindByID(ctx context.Context, ID string) (*e.User, error)
}
