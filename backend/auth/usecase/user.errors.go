package usecase

import "errors"

// ErrSalesDocumentConflict is throwed when a document (transaction unique field) already exists in the repository
var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
