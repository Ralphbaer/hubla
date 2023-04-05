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

func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*e.User, error) {
	return uc.UserRepo.FindByEmail(ctx, email)
}

// StoreFileContent stores a new Transaction
func (uc *UserUseCase) GetUserByID(ctx context.Context, ID string) (*e.User, error) {
	return uc.UserRepo.FindByID(ctx, ID)
}
