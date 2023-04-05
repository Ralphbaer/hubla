package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
	r "github.com/Ralphbaer/hubla/backend/auth/repository"
)

// UserUseCase represents a collection of use cases for Transaction operations
type UserUseCase struct {
	UserRepo r.UserRepository
}

// StoreFileContent stores a new Transaction
func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*e.User, error) {
	// get user by email on postgres
	return nil, nil
}
