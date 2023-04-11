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

// GetUserByEmail retrieves a user by their email address from the user repository and returns it.
// If no user is found with the specified email address, a ValidationError is returned with an EntityNotFoundError.
func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*e.User, error) {
	u, err := uc.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		if err, ok := err.(common.EntityNotFoundError); ok {
			return nil, common.UnauthorizedError{
				ErrCode: "ErrInvalidEmailOrPassword",
				Message: ErrInvalidEmailOrPassword.Error(),
				Err:     err,
			}
		}
		return nil, err
	}

	return u, nil
}

// GetUserByID retrieves a user by their ID from the UserRepo.
// Returns a User pointer and an error if there's any issue.
func (uc *UserUseCase) GetUserByID(ctx context.Context, ID string) (*e.User, error) {
	return uc.UserRepo.FindByID(ctx, ID)
}
