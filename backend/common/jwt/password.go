package jwt

import (
	"github.com/Ralphbaer/hubla/backend/common"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed version of the provided password.
// Returns the hashed password as a string.
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword checks if the provided password matches the hashed password.
// Returns an error if the comparison fails.
func ComparePassword(hashedPassword string, signInUserPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(signInUserPassword)); err != nil {
		return common.UnauthorizedError{
			Message: ErrPasswordMismatch.Error(),
			Err:     ErrPasswordMismatch,
		}
	}

	return nil
}
