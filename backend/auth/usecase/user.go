package usecase

import (
	"context"

	e "github.com/Ralphbaer/hubla/backend/auth/entity"
	r "github.com/Ralphbaer/hubla/backend/auth/repository"
	"github.com/Ralphbaer/hubla/backend/common"
)

// UserUseCase represents a collection of use cases for Transaction operations
type UserUseCase struct {
	UserRepo r.UserRepository
}

func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*e.User, error) {
	u, err := uc.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		if err, ok := err.(common.EntityNotFoundError); ok {
			return nil, common.ValidationError{
				Message: ErrInvalidEmailOrPassword.Error(),
				Err:     err,
			}
		}
		return nil, err
	}

	return u, nil
}

// StoreFileContent stores a new Transaction
func (uc *UserUseCase) GetUserByID(ctx context.Context, ID string) (*e.User, error) {
	return uc.UserRepo.FindByID(ctx, ID)
}
